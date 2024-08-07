// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.25.3
// source: markets/v1/markets.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMarketsTokenList = "/markets.v1.Markets/TokenList"

type MarketsHTTPServer interface {
	TokenList(context.Context, *TokenListRequest) (*TokenListResponse, error)
}

func RegisterMarketsHTTPServer(s *http.Server, srv MarketsHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/token/list", _Markets_TokenList0_HTTP_Handler(srv))
}

func _Markets_TokenList0_HTTP_Handler(srv MarketsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TokenListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMarketsTokenList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TokenList(ctx, req.(*TokenListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenListResponse)
		return ctx.Result(200, reply)
	}
}

type MarketsHTTPClient interface {
	TokenList(ctx context.Context, req *TokenListRequest, opts ...http.CallOption) (rsp *TokenListResponse, err error)
}

type MarketsHTTPClientImpl struct {
	cc *http.Client
}

func NewMarketsHTTPClient(client *http.Client) MarketsHTTPClient {
	return &MarketsHTTPClientImpl{client}
}

func (c *MarketsHTTPClientImpl) TokenList(ctx context.Context, in *TokenListRequest, opts ...http.CallOption) (*TokenListResponse, error) {
	var out TokenListResponse
	pattern := "/v1/token/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMarketsTokenList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
