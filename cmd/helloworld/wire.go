//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/dawn303/go-app/internal/helloworld/biz"
	"github.com/dawn303/go-app/internal/helloworld/conf"
	"github.com/dawn303/go-app/internal/helloworld/data"
	"github.com/dawn303/go-app/internal/helloworld/server"
	"github.com/dawn303/go-app/internal/helloworld/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
