package routes

import (
	"hynie.de/ohmab/ent"
)

type Server struct {
	Client *ent.Client
}

func NewServer(client *ent.Client) *Server {
	return &Server{Client: client}
}
