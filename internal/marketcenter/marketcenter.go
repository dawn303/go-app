package marketcenter

import (
	"os"

	"github.com/dawn303/go-app/internal/marketcenter/server"
	"github.com/dawn303/go-app/internal/pkg/bootstrap"
	"github.com/dawn303/go-app/pkg/db"
	"github.com/dawn303/go-app/pkg/options"
	"github.com/go-kratos/kratos/v2"
)

var (
	ID, _ = os.Hostname()

	Name = "marketcenter"

	Version = "1.0.0"
)

type Config struct {
	HTTPOptions  *options.HTTPOptions
	MysqlOptions *db.MysqlOptions
}

func (c *Config) New() (*Server, error) {

	cfg := &server.Config{
		Http: *c.HTTPOptions,
	}

	appInfo := bootstrap.NewAppInfo(ID, Name, Version)

	app, cleanup, err := wireApp(appInfo, cfg, c.MysqlOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}

type Server struct {
	app *kratos.App
}

func (s *Server) Run() error {

	return nil
}
