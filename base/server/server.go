package server

import (
	"net/http"
	"time"

	"github.com/google/wire"
)

// Set is a Wire provider set that produces a *Server.
var Set = wire.NewSet(
	New,
	Params{},
)

// Server is a http server with basic functionalities.
type Server struct {
	httpServer   *http.Server
	serveTimeout time.Duration
}

// Config is configurations for the HTTP server.
type Config struct {
	Addr         string
	ServeTimeout time.Duration
}

// Params is a set of parameters for New.
type Params struct {
	Config Config
}

// New creates a new server.
func New(params *Params) *Server {
	httpServer := &http.Server{Addr: params.Config.Addr}

	return &Server{
		httpServer:   httpServer,
		serveTimeout: params.Config.ServeTimeout,
	}
}

// ListenAndServe starts the server.
func (s *Server) ListenAndServe() error {
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	return s.httpServer.ListenAndServe()
}
