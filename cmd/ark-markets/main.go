package main

import (
	"github.com/dawn303/go-app/internal/markets"
	"github.com/dawn303/go-app/pkg/options"

	genericapiserver "k8s.io/apiserver/pkg/server"
)

func main() {
	cfg := markets.Config{
		HTTPOptions:  options.NewHTTPOptions(),
		MySQLOptions: options.NewMySQLOptions(),
	}
	server, err := cfg.Complete().New()
	if err != nil {
		panic(err)
	}
	server.Run(genericapiserver.SetupSignalHandler())
}
