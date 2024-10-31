package api

import gen "go-time/generated"

var _ gen.StrictServerInterface = (*Server)(nil)

type Server struct {
}

func NewServer() gen.ServerInterface {
	return gen.NewStrictHandler(Server{}, nil)
}
