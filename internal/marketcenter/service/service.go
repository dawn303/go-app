package service

import (
	"github.com/dawn303/go-app/internal/marketcenter/biz"
	v1 "github.com/dawn303/go-app/pkg/api/marketcenter/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMarketCenterService)

type MarketCenterService struct {
	v1.UnimplementedMarketCenterServer
	biz biz.IBiz
}

func NewMarketCenterService(biz biz.IBiz) *MarketCenterService {
	return &MarketCenterService{biz: biz}
}
