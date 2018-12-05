package main

import (
	"context"
	"flag"
	"log"

	"github.com/google/wire"
	"github.com/terashi58/wire-example/base/rpc"
	"github.com/terashi58/wire-example/base/server"
)

var appSet = wire.NewSet(
	newApp,
	appParams{},
)

type cliFlags struct {
	Port       int
	StatusPort int
}

func parseFlags() *cliFlags {
	fg := &cliFlags{}
	flag.IntVar(&fg.Port, "port", 50051, "port no for gRPC")
	flag.IntVar(&fg.StatusPort, "status_port", 8080, "port no for HTTP")
	flag.Parse()
	return fg
}

type appParams struct {
	statusServer *server.Server
	rpcServer    *rpc.Server
}

type app struct {
	params appParams
}

func newApp(params appParams) *app {
	return &app{params}
}

func (a *app) Run() error {
	// TODO: graceful shutdown
	go a.params.statusServer.ListenAndServe()
	return a.params.rpcServer.ListenAndServe()
}

func main() {
	flags := parseFlags()

	app, cleanup, err := initializeApp(context.Background(), flags)
	if err != nil {
		log.Fatalf("fail: %v", err)
	}
	defer cleanup()
	if err := app.Run(); err != nil {
		log.Fatalf("fail: %v", err)
	}
}
