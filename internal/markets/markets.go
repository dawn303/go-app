package markets

import (
	"os"

	"github.com/dawn303/go-app/internal/markets/server"
	"github.com/dawn303/go-app/internal/pkg/bootstrap"
	"github.com/dawn303/go-app/pkg/db"
	"github.com/dawn303/go-app/pkg/options"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
)

var (
	ID, _ = os.Hostname()

	Name = "ark-markets"

	Version = "1.0.0"
)

type Config struct {
	HTTPOptions  *options.HTTPOptions
	MySQLOptions *options.MySQLOptions
}

func (c *Config) Complete() *completeConfig {
	return &completeConfig{c}
}

type completeConfig struct {
	*Config
}

func (c *completeConfig) New() (*Server, error) {

	cfg := &server.Config{
		Http: *c.HTTPOptions,
	}

	appInfo := bootstrap.NewAppInfo(ID, Name, Version)

	var mysqlOptions db.MySQLOptions
	_ = copier.Copy(&mysqlOptions, c.MySQLOptions)

	app, cleanup, err := wireApp(appInfo, cfg, &mysqlOptions)
	if err != nil {
		return nil, err
	}
	defer cleanup()

	return &Server{app: app}, nil
}

type Server struct {
	app *kratos.App
}

func (s *Server) Run(stopCh <-chan struct{}) error {
	go func() {
		if err := s.app.Run(); err != nil {
			panic("Failed to start server: " + err.Error())
		}
	}()

	<-stopCh

	log.Infof("Gracefully shutting down server ...")

	if err := s.app.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
		return err
	}

	return nil
}
