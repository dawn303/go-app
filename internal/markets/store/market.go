package store

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/markets/v1"
	"gorm.io/gorm"
)

type MarketStore interface {
	TokenList(ctx context.Context, req *v1.TokenListRequest) (*v1.TokenListResponse, error)
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

func (s *marketStore) TokenList(ctx context.Context, req *v1.TokenListRequest) (*v1.TokenListResponse, error) {
	// ...
	var models *v1.TokenListResponse
	err := s.db(ctx).Where("kind = ?", req.Kind).Find(models).Error
	return models, err

	// return &v1.TokenListResponse{
	// 	TotalCount: 1,
	// 	Tokens: []*v1.TokenReply{
	// 		{Token: "0x003144B41d9743D402c5bdF3f72Ca0f327aA0Bca", Price: "1.90803284952", QuoteChange: "1.89%"},
	// 	}}, nil
}
