package market

import (
	"context"

	v1 "github.com/dawn303/go-app/pkg/api/markets/v1"
)

func (b *marketBiz) TokenList(ctx context.Context, req *v1.TokenListRequest) (*v1.TokenListResponse, error) {
	list, err := b.ds.Markets().TokenList(ctx, req)
	if err != nil {
		return nil, v1.ErrorTokenNotFound("kind: %s", req.Kind)
	}
	return list, nil
}
