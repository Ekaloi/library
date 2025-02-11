package server

import (
	"Github.com/Ekaloi/library/api"
)

type Server struct {
	client *api.Client
}


func NewServer() *Server{
	return &Server{client: api.NewClient()}
}
