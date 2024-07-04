package service

import (
	"github.com/dawn303/go-app/internal/markets/biz"
	v1 "github.com/dawn303/go-app/pkg/api/markets/v1"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMarketsService)

type MarketsService struct {
	v1.UnimplementedMarketsServer
	biz biz.IBiz
}

func NewMarketsService(biz biz.IBiz) *MarketsService {
	return &MarketsService{biz: biz}
}
