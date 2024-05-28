package store

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/marketcenter/v1"
	"gorm.io/gorm"
)

type MarketStore interface {
	CurrencyList(ctx context.Context, req *v1.CurrencyListRequest) (*v1.CurrencyListResponse, error)
}

type marketStore struct {
	ds *datastore
}

func newMarketStore(ds *datastore) *marketStore {
	return &marketStore{ds}
}

func (s *marketStore) db(ctx context.Context) *gorm.DB {
	return s.ds.Core(ctx)
}

func (s *marketStore) CurrencyList(ctx context.Context, req *v1.CurrencyListRequest) (*v1.CurrencyListResponse, error) {
	// ...
	var models *v1.CurrencyListResponse
	err := s.db(ctx).Where("token = ?", req.Token).Find(models).Error
	return models, err
}
