// +build wireinject

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/terashi58/wire-example/app/config"
	"github.com/terashi58/wire-example/app/service"
	"github.com/terashi58/wire-example/base/database/mysql"
	"github.com/terashi58/wire-example/base/rpc"
	"github.com/terashi58/wire-example/base/server"
)

func httpConfig(flags *cliFlags) server.Config {
	return server.Config{
		Addr:         fmt.Sprintf(":%d", flags.StatusPort),
		ServeTimeout: 5 * time.Second,
	}
}

func rpcConfig(flags *cliFlags) rpc.Config {
	return rpc.Config{
		Addr: fmt.Sprintf(":%d", flags.Port),
	}
}

func initializeApp(ctx context.Context, flags *cliFlags) (*app, func(), error) {
	wire.Build(
		appSet,
		server.Set,
		service.Set,
		mysql.Set,
		config.Set,
		httpConfig,
		rpcConfig,
	)
	return nil, nil, nil
}
