package market

import (
	"context"

	"github.com/dawn303/go-app/internal/marketcenter/store"
	v1 "github.com/dawn303/go-app/pkg/api/marketcenter/v1"
)

type MarketBiz interface {
	CurrencyList(ctx context.Context, req *v1.CurrencyListRequest) (*v1.CurrencyListResponse, error)
}

type marketBiz struct {
	ds store.IStore
}

var _ MarketBiz = (*marketBiz)(nil)

func New(ds store.IStore) *marketBiz {
	return &marketBiz{ds}
}
