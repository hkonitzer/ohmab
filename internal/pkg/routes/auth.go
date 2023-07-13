package routes

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/oauth"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/user"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"net/http"
	"strings"
	"time"
)

var logger = log.GetLoggerInstance()
var configurations = config.GetX()

func RegisterOAuthAPI(r *chi.Mux, srv *Server) *chi.Mux {
	verifier := &UserVerifier{
		Client: srv.Client,
	}
	exp := time.Second
	if configurations.ENVIRONMENT == config.DevelopmentEnvironment {
		exp *= 30000
	} else {
		exp *= 300
	}
	s := oauth.NewBearerServer(
		configurations.OAUTHSECRETKEY,
		exp,
		verifier,
		nil)
	srv.OAuthServer = s

	r.Post("/token", srv.OAuthServer.UserCredentials)
	// r.Post("/auth", srv.OAuthServer.ClientCredentials)
	return r
}

func (s *Server) CheckUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		view := privacy.FromContext(ctx)
		uv := privacy.UserViewer{}
		if view == nil { // should be nil!
			claims := ctx.Value(oauth.ClaimsContext).(map[string]string)
			uv.Claims = claims
			uuid_, err := uuid.Parse(uv.GetUserID())
			if err != nil {
				logger.Fatal().Msgf("Error parsing UUID: %v", err)
			}
			u := s.Client.User.Query().Where(user.IDEQ(uuid_)).Select(user.FieldRole).OnlyX(ctx)
			uv.Scopes = strings.Split(ctx.Value(oauth.ScopeContext).(string), ",")
			switch u.Role {
			case privacy.ViewerRoleAsString():
				uv.Role = privacy.View
			case privacy.OwnerRoleAsString():
				uv.Role = privacy.Owner
			case privacy.AdminRoleAsString():
				uv.Role = privacy.Admin
			default:
				logger.Fatal().Msgf("Unknown role: %v for user-id '%v'", u.Role, uuid_)
			}
		}
		next.ServeHTTP(w, r.WithContext(privacy.NewContext(r.Context(), &uv)))
	})
}

// UserVerifier provides user credentials verifier for testing.
type UserVerifier struct {
	Client *ent.Client
}

// ValidateUser validates username and password returning an error if the user credentials are wrong
func (t *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	u, err := t.Client.User.Query().Where(user.LoginEQ(username)).WithBusinesses().Only(r.Context())
	if err != nil {
		logger.Debug().Msgf("Error getting user '%s': %v", username, err)
		return errors.New("wrong username")
	}
	if u.Role != privacy.AdminRoleAsString() { // test given scopes against allowed scopes because user has no admin role
		if scope == "" {
			return errors.New("scope is empty")
		}
		// init UserViewer (for scope testing only)
		uv := privacy.UserViewer{}
		// get scope
		scopes := strings.Split(scope, ",")
		// set scopes
		for _, s := range scopes {
			for _, b := range u.Edges.Businesses {
				if s == b.ID.String() {
					uv.Scopes = append(uv.Scopes, b.ID.String())
					break
				}
			}
		}
		// check all given scopes against the allowed ones
		if (len(scopes) != len(uv.Scopes)) || (len(uv.Scopes) == 0) {
			logger.Debug().Msgf("requested scope(s) '%v' not allowed for user '%s'", scope, username)
			return errors.New("scope not allowed")
		}
	}
	if utils.DoPasswordsMatch(u.Passwordhash, password) {
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
	uid, err := t.Client.User.Query().Where(user.Login(credential)).OnlyID(r.Context()) // @TODO set this in ValidateUser in context and get this here from context
	if err != nil {
		return nil, err
	}
	claims := make(map[string]string)
	claims["user_"+user.FieldID] = uid.String()
	return claims, nil
}

// AddProperties provides additional information to the token response
func (*UserVerifier) AddProperties(tokenType oauth.TokenType, credential, tokenID, scope string, r *http.Request) (map[string]string, error) {
	props := make(map[string]string)
	return props, nil
}

// ValidateTokenID validates token ID
func (*UserVerifier) ValidateTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	logger.Debug().Msgf("ValidateTokenID: %s, %s, %s, %s", tokenType, credential, tokenID, refreshTokenID)
	return nil
}

// StoreTokenID saves the token id generated for the user
func (*UserVerifier) StoreTokenID(tokenType oauth.TokenType, credential, tokenID, refreshTokenID string) error {
	return nil
}
