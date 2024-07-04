package service

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/markets/v1"
)

func (s *MarketsService) TokenList(ctx context.Context, req *v1.TokenListRequest) (*v1.TokenListResponse, error) {	
	return s.biz.Markets().TokenList(ctx, req)
}
