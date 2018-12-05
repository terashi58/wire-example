package greeter

import (
	"context"

	"database/sql"
	"github.com/google/wire"
	"github.com/terashi58/wire-example/app/proto"
	"google.golang.org/grpc"
)

// Set is a Wire provider set that produces a *Server.
var Set = wire.NewSet(
	NewServer,
)

// Server implements the gRPC Greeter service.
type Server struct {
	db *sql.DB
}

// NewServer creates a new Greeter server.
func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

// Register implements rpc.ServiceImpl.
func (s *Server) Register(gs *grpc.Server) {
	proto.RegisterGreeterServer(gs, s)
}

// HelloWorld returns a "hello world" message.
func (s *Server) HelloWorld(ctx context.Context, req *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	// Use s.db some how.
	return &proto.HelloWorldResponse{Message: "hello world"}, nil
}
