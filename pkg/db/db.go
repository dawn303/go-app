package db

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewMySQL)
