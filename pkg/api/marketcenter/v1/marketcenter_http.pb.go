// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.25.3
// source: marketcenter/v1/marketcenter.proto

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

const OperationMarketCenterCurrencyList = "/marketcenter.v1.MarketCenter/CurrencyList"

type MarketCenterHTTPServer interface {
	CurrencyList(context.Context, *CurrencyListRequest) (*CurrencyListResponse, error)
}

func RegisterMarketCenterHTTPServer(s *http.Server, srv MarketCenterHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/currency/list", _MarketCenter_CurrencyList0_HTTP_Handler(srv))
}

func _MarketCenter_CurrencyList0_HTTP_Handler(srv MarketCenterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrencyListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMarketCenterCurrencyList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CurrencyList(ctx, req.(*CurrencyListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrencyListResponse)
		return ctx.Result(200, reply)
	}
}

type MarketCenterHTTPClient interface {
	CurrencyList(ctx context.Context, req *CurrencyListRequest, opts ...http.CallOption) (rsp *CurrencyListResponse, err error)
}

type MarketCenterHTTPClientImpl struct {
	cc *http.Client
}

func NewMarketCenterHTTPClient(client *http.Client) MarketCenterHTTPClient {
	return &MarketCenterHTTPClientImpl{client}
}

func (c *MarketCenterHTTPClientImpl) CurrencyList(ctx context.Context, in *CurrencyListRequest, opts ...http.CallOption) (*CurrencyListResponse, error) {
	var out CurrencyListResponse
	pattern := "/v1/currency/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMarketCenterCurrencyList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
