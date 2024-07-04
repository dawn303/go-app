package market

import (
	"context"

	"github.com/dawn303/go-app/internal/markets/store"
	v1 "github.com/dawn303/go-app/pkg/api/markets/v1"
)

type MarketBiz interface {
	TokenList(ctx context.Context, req *v1.TokenListRequest) (*v1.TokenListResponse, error)
}

type marketBiz struct {
	ds store.IStore
}

var _ MarketBiz = (*marketBiz)(nil)

func New(ds store.IStore) *marketBiz {
	return &marketBiz{ds}
}
