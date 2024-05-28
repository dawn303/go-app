package bootstrap

import (
	"github.com/go-kratos/kratos/v2"
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(wire.Struct(new(AppConfig), "*"), NewLogger, NewApp)

type AppInfo struct {
	ID      string
	Name    string
	Version string
}

func NewAppInfo(id, name, version string) *AppInfo {
	return &AppInfo{ID: id, Name: name, Version: version}
}

type AppConfig struct {
	Info   *AppInfo
	Logger krtlog.Logger
}

func NewApp(c *AppConfig, servers ...transport.Server) *kratos.App {
	return kratos.New(
		kratos.ID(c.Info.ID),
		kratos.Name(c.Info.Name),
		kratos.Version(c.Info.Version),
		kratos.Logger(c.Logger),
		kratos.Server(servers...),
	)
}
