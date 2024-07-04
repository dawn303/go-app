//go:build wireinject
// +build wireinject

package markets

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/dawn303/go-app/internal/markets/biz"
	"github.com/dawn303/go-app/internal/markets/server"
	"github.com/dawn303/go-app/internal/markets/service"
	"github.com/dawn303/go-app/internal/markets/store"
	"github.com/dawn303/go-app/internal/pkg/bootstrap"
	"github.com/dawn303/go-app/pkg/db"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func wireApp(*bootstrap.AppInfo, *server.Config, *db.MySQLOptions) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		db.ProviderSet,
		store.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	)

	return nil, nil, nil
}
