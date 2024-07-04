package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewServers, NewHTTPServer)

func NewServers(hs *http.Server) []transport.Server {
	return []transport.Server{hs}
}
