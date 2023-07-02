package routes

import (
	"github.com/go-chi/oauth"
	"hynie.de/ohmab/ent"
)

type Server struct {
	Client      *ent.Client
	OAuthServer *oauth.BearerServer
}

func NewServer(client *ent.Client) *Server {
	return &Server{Client: client}
}
