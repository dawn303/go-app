package biz

import (
	"github.com/dawn303/go-app/internal/marketcenter/biz/market"
	"github.com/dawn303/go-app/internal/marketcenter/store"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
	Markets() market.MarketBiz
}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

func (b *biz) Markets() market.MarketBiz {
	return market.New(b.ds)
}
