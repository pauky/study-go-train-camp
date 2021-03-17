// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"study-go-train-camp/internal/biz"
	"study-go-train-camp/internal/conf"
	"study-go-train-camp/internal/data"
	"study-go-train-camp/internal/server"
	"study-go-train-camp/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
