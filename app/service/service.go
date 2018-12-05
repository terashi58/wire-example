package service

import (
	"github.com/google/wire"

	"github.com/terashi58/wire-example/app/service/greeter"
	"github.com/terashi58/wire-example/app/service/pinger"
	"github.com/terashi58/wire-example/base/rpc"
	"github.com/terashi58/wire-example/base/rpc/interceptor"
)

// Set is a Wire provider set that produces a *rpc.Server.
var Set = wire.NewSet(
	rpc.Set,             // Provider Set for grpc.Server
	interceptor.Default, // Provider of gRPC interceptors
	greeter.Set,
	pinger.Set,
	NewServices,
)

// NewServices provides a list of gRPC services.
func NewServices(gs *greeter.Server, ps *pinger.Server) []rpc.ServiceImpl {
	return []rpc.ServiceImpl{gs, ps}
}
