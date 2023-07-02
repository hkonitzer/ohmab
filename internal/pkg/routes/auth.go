package routes

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/oauth"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/user"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/log"
	"hynie.de/ohmab/internal/pkg/privacy"
	"hynie.de/ohmab/internal/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

var logger = log.GetLoggerInstance()
var configurations = config.GetConfigurationX()
var secret = configurations.OAUTHSECRETKEY

func RegisterOAuthAPI(r *chi.Mux, srv *Server) *chi.Mux {
	verifier := &UserVerifier{
		Client: srv.Client,
	}

	s := oauth.NewBearerServer(
		configurations.OAUTHSECRETKEY, // @TODO move to config
		time.Second*300,
		verifier,
		nil)
	srv.OAuthServer = s

	r.Post("/token", srv.OAuthServer.UserCredentials)
	r.Get("/test", srv.Timetables)
	// r.Post("/auth", srv.OAuthServer.ClientCredentials)
	return r
}

func CheckUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Println(ctx.Value(oauth.CredentialContext))
		fmt.Println(ctx.Value(oauth.ClaimsContext))

		/*
			r = r.WithContext(privacy.NewContext(r.Context(), privacy.UserViewer{Role: privacy.View}))

			fmt.Println("user role", privacy.FromContext(r.Context()))
		*/
		next.ServeHTTP(w, r)
	})
}

// UserVerifier provides user credentials verifier for testing.
type UserVerifier struct {
	Client *ent.Client
}

// ValidateUser validates username and password returning an error if the user credentials are wrong
func (t *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	u, err := t.Client.User.Query().Where(user.LoginEQ(username)).Only(r.Context())
	if err != nil {
		logger.Debug().Msgf("Error getting user '%s': %v", username, err)
		return errors.New("wrong username")
	}
	if utils.DoPasswordsMatch(u.Passwordhash, password) {
		if u.Role == privacy.AdminRoleAsInt() {
			r.WithContext(privacy.NewContext(r.Context(), privacy.UserViewer{Role: privacy.Admin}))
		} else {
			r.WithContext(privacy.NewContext(r.Context(), privacy.UserViewer{Role: privacy.View}))
		}

		return nil
	}
	logger.Debug().Msgf("provided password does not match for user '%s'", username)
	return errors.New("wrong password")
}

// ValidateClient validates clientID and secret returning an error if the client credentials are wrong
func (*UserVerifier) ValidateClient(clientID, clientSecret, scope string, r *http.Request) error {
	return errors.New("not implemented")
}

// ValidateCode validates token ID
func (*UserVerifier) ValidateCode(clientID, clientSecret, code, redirectURI string, r *http.Request) (string, error) {
	fmt.Println("ValidateCode", clientID, clientSecret, code, redirectURI)
	return "", nil
}

// AddClaims provides additional claims to the token
func (t *UserVerifier) AddClaims(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	u, err := t.Client.User.Query().Where(user.Login(credential)).Only(r.Context())
	if err != nil {
		return nil, err
	}
	claims := make(map[string]string)
	claims["user_id"] = u.ID.String()
	claims["role"] = strconv.Itoa(u.Role)

	return claims, nil
}

// AddProperties provides additional information to the token response
func (*UserVerifier) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	props := make(map[string]string)
	return props, nil
}

// ValidateTokenID validates token ID
func (*UserVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {

	return nil
}

// StoreTokenID saves the token id generated for the user
func (*UserVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}
