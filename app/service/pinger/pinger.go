package pinger

import (
	"context"

	"github.com/google/wire"
	"github.com/terashi58/wire-example/app/proto"
	"google.golang.org/grpc"
)

// Set is a Wire provider set that produces a *Server.
var Set = wire.NewSet(
	NewServer,
)

// Server implements the gRPC Pinger service.
type Server struct {
}

// NewServer creates a new Pinger server.
func NewServer() *Server {
	return &Server{}
}

// Register implements rpc.ServiceImpl.
func (s *Server) Register(gs *grpc.Server) {
	proto.RegisterPingerServer(gs, s)
}

// Ping returns a message.
func (s *Server) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{Pong: req.Ping}, nil
}
