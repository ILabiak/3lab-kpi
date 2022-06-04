package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ILabiak/3lab-kpi/pkg/forums"
)

type HttpPortNumber int
type ForumsApiServer struct {
	Port HttpPortNumber

	ForumsHandler forums.HttpHandlerFunc

	server *http.Server
}

func (s *ForumsApiServer) Start() error {
	if s.ForumsHandler == nil {
		return fmt.Errorf("Forums HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/forums", s.ForumsHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *ForumsApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
