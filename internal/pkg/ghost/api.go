package ghost

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var logger = log.GetLoggerInstance()
var configurations = config.GetX()

type GhostSite struct {
	Site struct {
		Title       string `json:"title"`
		Description string `json:"description,omitempty"`
		Logo        string `json:"logo,omitempty"`
		Icon        string `json:"icon,omitempty"`
		AccentColor string `json:"accent_color,omitempty"`
		Locale      string `json:"locale,omitempty"`
		URL         string `json:"url,omitempty"`
		Version     string `json:"version,omitempty"`
	} `json:"site"`
}

type Entity interface {
	Request()
	PageEntity
}

type PostEntity struct {
	Posts []SingleRessource `json:"posts"`
}

type PageEntity struct {
	Pages []SingleRessource `json:"pages"`
}

type SingleRessource struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug,omitempty"`
	UpdatedAt   string `json:"updated_at"`
	PublishedAt string `json:"published_at"`
	//MobileDoc   string `json:"mobiledoc,omitempty"`
	HTML     string `json:"html,omitempty"`
	Status   string `json:"status,omitempty"`
	Featured bool   `json:"featured,omitempty"`
	Excerpt  string `json:"custom_excerpt,omitempty"`
	Tags     []Tag  `json:"tags,omitempty"`
}

type Tag struct {
	Name       string `json:"name,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type Errors struct {
	GhostError []RequestError `json:"errors"`
}

type RequestError struct {
	Message        string                `json:"message"`
	Context        string                `json:"context,omitempty"`
	Type           string                `json:"type"`
	Details        []RequestErrorDetails `json:"details,omitempty"`
	Property       string                `json:"property,omitempty"`
	Help           string                `json:"help,omitempty"`
	Code           string                `json:"code,omitempty"`
	Id             string                `json:"id"`
	GhostErrorCode string                `json:"ghostErrorCode,omitempty"`
}

type RequestErrorDetails struct {
	Keyword    string `json:"keyword,omitempty"`
	DataPath   string `json:"dataPath,omitempty"`
	SchemaPath string `json:"schemaPath,omitempty"`
	Params     string `json:"params,omitempty"`
	Message    string `json:"message,omitempty"`
}

func CreateJWT() (string, error) {
	// setup payload
	payload := fmt.Sprintf(
		`{"iat": %s, "exp": %s, "aud": "/admin/"}`,
		strconv.FormatInt(time.Now().Unix(), 10),
		strconv.FormatInt(time.Now().Add(time.Minute*3).Unix(), 10),
	)
	// get ghost admin key
	ghostAdminAPIKey := configurations.Ghost.Key
	if ghostAdminAPIKey == "" {
		logger.Fatal().Msgf("missing Ghost Admin API key")
	}
	if !strings.Contains(ghostAdminAPIKey, ":") {
		logger.Fatal().Msgf("error reading Ghost Admin API key: missing ':' separator")
	}
	ghostKey := strings.Split(ghostAdminAPIKey, ":")[1]
	ghostKeyId := strings.Split(ghostAdminAPIKey, ":")[0]
	key, _ := hex.DecodeString(ghostKey)

	token, err := jose.Sign(payload, jose.HS256, key,
		jose.Header("alg", "HS256"),
		jose.Header("kid", ghostKeyId),
		jose.Header("typ", "JWT"))
	if err != nil {
		return "", err
		//logger.Fatal().Msgf("Error creating token: %v", err)
	}
	return token, nil
}

func CreateJWTX() string {
	t, err := CreateJWT()
	if err != nil {
		logger.Fatal().Msgf("Error creating token: %v", err)
	}
	return t
}

func SetRequestHeadersForGhost(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Ghost %s", CreateJWTX()))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Version", "v4.0")
}

func (gp PostEntity) Request() error {
	// should be only one page
	if len(gp.Posts) > 1 {
		return fmt.Errorf("%d posts found, expected only one", len(gp.Posts))
	} else if len(gp.Posts) == 0 {
		return fmt.Errorf("no post found, need one with ID set")
	}
	// build the URL
	requestURL := fmt.Sprintf(
		"ghost/api/admin/posts/%s",
		gp.Posts[0].ID,
	)
	resBody, err := request(requestURL, "GET", nil)
	if err != nil {
		return err
	}
	// convert (unmarshal) the response body to a PageEntity struct
	err = json.Unmarshal(resBody, &gp)
	return err
}

func (gp PostEntity) UpdatePost() error {
	// should be only one page
	if len(gp.Posts) > 1 {
		return fmt.Errorf("more than one post found, expected only one")
	} else if len(gp.Posts) == 0 {
		return fmt.Errorf("no post entity found, need one with ID set")
	}
	// convert (marshal) the SingleRessource struct to json
	gjson, err := json.Marshal(gp)
	if err != nil {
		return err
	}
	// build the URL
	requestURL := fmt.Sprintf(
		"ghost/api/admin/posts/%s?source=html",
		gp.Posts[0].ID,
	)
	// create the request
	resBody, err := request(requestURL, http.MethodPut, bytes.NewBuffer(gjson))
	if err != nil {
		return err
	}
	// convert (unmarshal) the response body to a PageEntity struct
	err = json.Unmarshal(resBody, &gp)
	return err
}

func (gp PageEntity) Request() error {
	// should be only one page
	if len(gp.Pages) > 1 {
		return fmt.Errorf("more than one page found, expected only one")
	} else if len(gp.Pages) == 0 {
		return fmt.Errorf("no page found, need one with ID set")
	}
	// build the URL
	requestURL := fmt.Sprintf(
		"ghost/api/admin/pages/%s/?source=html",
		gp.Pages[0].ID,
	)
	resBody, err := request(requestURL, "GET", nil)
	if err != nil {
		return err
	}
	// convert (unmarshal) the response body to a PageEntity struct
	err = json.Unmarshal(resBody, &gp)
	return err
}

func (gp PageEntity) UpdatePage() error {
	// should be only one page
	if len(gp.Pages) > 1 {
		return fmt.Errorf("more than one page found, expected only one")
	} else if len(gp.Pages) == 0 {
		return fmt.Errorf("no page found, need one with ID set")
	}
	// convert (marshal) the SingleRessource struct to json
	gjson, err := json.Marshal(gp)
	if err != nil {
		return err
	}
	// build the URL
	requestURL := fmt.Sprintf(
		"ghost/api/admin/pages/%s?source=html",
		gp.Pages[0].ID,
	)
	// create the request
	logger.Debug().Msgf("UpdatePage() create request with content:\r\n %s", gjson)
	resBody, err := request(requestURL, http.MethodPut, bytes.NewBuffer(gjson))
	if err != nil {
		return err
	}
	// convert (unmarshal) the response body to a PageEntity struct
	err = json.Unmarshal(resBody, &gp)
	return err
}

func request(url string, method string, payload io.Reader) ([]byte, error) {
	// build the URL
	requestURL := fmt.Sprintf(
		"%s/%s",
		configurations.Ghost.BaseURL,
		url,
	)
	logger.Debug().Msgf("request ghost with method %s and URL: %s", method, requestURL)
	// create the request
	req, err := http.NewRequest(method, requestURL, payload)
	if err != nil {
		return nil, err
	}
	// set needed headers
	SetRequestHeadersForGhost(req)
	// fire up
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	logger.Debug().Msgf("request done, status: %d", res.StatusCode)
	if res.StatusCode != 200 {
		ghostErrors := Errors{}
		_ = json.Unmarshal(resBody, &ghostErrors)
		if len(ghostErrors.GhostError) > 0 {
			logger.Debug().Msgf("error requesting url %s with status %s; ghost error msg: %s type: %s Context: %s",
				url, res.Status, ghostErrors.GhostError[0].Message, ghostErrors.GhostError[0].Type, ghostErrors.GhostError[0].Context)
		}
		return nil, fmt.Errorf("error requesting url: %s, got status %s", url, res.Status)
	}
	return resBody, err
}

func RequestSite() (GhostSite, error) {
	site := GhostSite{}
	// build the URL
	requestURL := fmt.Sprintf(
		"%s/ghost/api/admin/site",
		configurations.Ghost.BaseURL,
	)
	logger.Debug().Msgf("request ghost site data with URL: %s", requestURL)
	// create the request
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		logger.Fatal().Msgf("Error creating request: %v", err)
	}
	SetRequestHeadersForGhost(req)
	// send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return site, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return site, err
	}
	if res.StatusCode != 200 {
		ghostErrors := Errors{}
		_ = json.Unmarshal(resBody, &ghostErrors)
		if len(ghostErrors.GhostError) > 0 {
			logger.Debug().Msgf("error requesting ghost site data with status %s; ghost error msg: %s type: %s Context: %s",
				res.Status, ghostErrors.GhostError[0].Message, ghostErrors.GhostError[0].Type, ghostErrors.GhostError[0].Context)
		}
		return site, fmt.Errorf("error requesting ghost site data, got status %s", res.Status)
	}
	// convert (unmarshal) the response body to a PageEntity struct
	err = json.Unmarshal(resBody, &site)
	return site, nil
}

var mbdregx = regexp.MustCompile(`\["html",{"html":"(.*)"}\]`)

func ReplaceMobileDoc(mobiledoc string, newMobiledoc string) string {
	// tidy html
	hc := strings.ReplaceAll(newMobiledoc, "\n", "")
	hc = strings.ReplaceAll(hc, "\"", "\\\"")
	// replace mobiledoc
	if mobiledoc == "" {
		mobiledoc = newMobiledoc
	}
	validDoc := mbdregx.ReplaceAllString(mobiledoc, fmt.Sprintf(`["html",{"html":"%s"}]`, hc))
	return validDoc
}
