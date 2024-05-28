package service

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/marketcenter/v1"
)

func (s *MarketCenterService) CurrencyList(ctx context.Context, req *v1.CurrencyListRequest) (*v1.CurrencyListResponse, error) {
	return s.biz.Markets().CurrencyList(ctx, req)
}
