package routes

import (
	"context"
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
		exp *= 1800
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

var readerContext = privacy.NewViewerContext(context.Background(), privacy.View, nil)

func (s *Server) CheckUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uv, _ := privacy.FromContext(ctx)
		if uv.IsEmpty() { // should be empty!
			claims := ctx.Value(oauth.ClaimsContext).(map[string]string)
			uv.Claims = claims // sets the user id
			uuid_, err := uuid.Parse(uv.GetUserID())
			if err != nil {
				logger.Fatal().Msgf("Error parsing UUID: %w", err)
			}
			u := s.Client.User.Query().Where(user.IDEQ(uuid_)).Select(user.FieldRole).OnlyX(readerContext)
			uv.Scopes = strings.Split(ctx.Value(oauth.ScopeContext).(string), ",")
			err = uv.SetRoleAsString(u.Role)
			if err != nil {
				logger.Fatal().Msgf("Unknown role: %v for user-id '%v'", u.Role, uuid_)
			}
			next.ServeHTTP(w, r.WithContext(uv.ToContext(r.Context())))
		}
		next.ServeHTTP(w, r)
	})
}

// UserVerifier provides user credentials verifier for testing.
type UserVerifier struct {
	Client     *ent.Client
	userViewer privacy.UserViewer
}

// ValidateUser validates username and password returning an error if the user credentials are wrong
func (t *UserVerifier) ValidateUser(username, password, scope string, r *http.Request) error {
	u, err := t.Client.User.Query().Where(user.LoginEQ(username)).WithBusinesses().Only(readerContext)
	if err != nil {
		logger.Debug().Msgf("Error getting user '%s': %v", username, err)
		return errors.New("wrong username")
	}
	// init UserViewer (for scope testing and request context later)
	ex := false
	t.userViewer, ex = privacy.FromContext(r.Context())
	if !ex {
		t.userViewer = privacy.UserViewer{}
	}
	t.userViewer.SetRoleAsString(u.Role)
	t.userViewer.SetUserID(u.ID.String())
	for _, b := range u.Edges.Businesses {
		t.userViewer.Scopes = append(t.userViewer.Scopes, b.ID.String())
	}
	// Set scopes from database on viewer, but see below @TODO: refactoring scopes for login?

	/* SKIPPING SCOPES FOR NOW
	if u.Role != privacy.AdminRoleAsString() { // test given scopes against allowed scopes because user has no admin role
		if scope == "" {
			return errors.New("scope is empty")
		}
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
	*/
	if utils.DoPasswordsMatch(u.Passwordhash, password) {
		r.WithContext(t.userViewer.ToContext(r.Context()))
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
	if len(t.userViewer.Claims) > 0 { // claims should be set through ValidateUser with SetUserID
		return t.userViewer.Claims, nil
	}
	if !t.userViewer.IsEmpty() && t.userViewer.GetUserID() != "" {
		claims := make(map[string]string)
		claims["user_"+user.FieldID] = t.userViewer.GetUserID()
		return claims, nil
	}
	uid, err := t.Client.User.Query().Where(user.Login(credential)).OnlyID(readerContext)
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
