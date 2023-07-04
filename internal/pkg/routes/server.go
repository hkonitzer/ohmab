package routes

import (
	"github.com/go-chi/oauth"
	"hynie.de/ohmab/ent"
	"strings"
)

const (
	PublicAPIRoute = "/public"
)

type Server struct {
	Client      *ent.Client
	OAuthServer *oauth.BearerServer
}

func NewServer(client *ent.Client) *Server {
	return &Server{Client: client}
}

func HasPublicRoute(routes []string) bool {
	for _, r := range routes {
		if strings.HasPrefix(r, PublicAPIRoute) {
			return true
		}
	}
	return false
}
