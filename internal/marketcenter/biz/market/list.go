package market

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/marketcenter/v1"
)

func (b *marketBiz) CurrencyList(ctx context.Context, req *v1.CurrencyListRequest) (*v1.CurrencyListResponse, error) {
	return b.ds.Markets().CurrencyList(ctx, req)
}
