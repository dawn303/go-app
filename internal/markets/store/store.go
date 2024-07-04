package store

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	s    *datastore
)

type IStore interface {
	Markets() MarketStore
}

type datastore struct {
	core *gorm.DB
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		s = &datastore{core: db}
	})

	return s
}

func (ds *datastore) Core(ctx context.Context) *gorm.DB {
	// todo
	// add transactional DB instance by ctx.transactionKey
	// ...

	return ds.core
}

func (ds *datastore) Markets() MarketStore {
	return newMarketStore(ds)
}
