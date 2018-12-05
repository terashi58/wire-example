package rpc

import (
	"net"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

// Set is a Wire provider set that produces a *Server.
var Set = wire.NewSet(
	New,
	Params{},
)

// Server is a grpc server with basic functionalities.
type Server struct {
	addr       string
	grpcServer *grpc.Server
}

// ServiceImpl is a server implementing a gRPC service.
type ServiceImpl interface {
	Register(*grpc.Server)
}

// Config is configurations for the gRPC server.
type Config struct {
	Addr string
}

// Params is a set of parameters for New.
type Params struct {
	Config       Config
	Services     []ServiceImpl
	Interceptors []grpc.UnaryServerInterceptor
}

// New creates a new gRPC server.
func New(params *Params) *Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(params.Interceptors...)),
	)
	for _, svc := range params.Services {
		svc.Register(server)
	}
	return &Server{
		addr:       params.Config.Addr,
		grpcServer: server,
	}
}

// ListenAndServe starts the server.
func (s *Server) ListenAndServe() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	return s.grpcServer.Serve(lis)
}
