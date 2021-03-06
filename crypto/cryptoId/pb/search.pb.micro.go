// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: search.proto

package search

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lights-T/lib-go/openapiv2"
	proto1 "github.com/micro/go-micro/v2/api/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Search service

func NewSearchEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Search service

type SearchService interface {
	// 搜索匹配
	Match(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type searchService struct {
	c    client.Client
	name string
}

func NewSearchService(name string, c client.Client) SearchService {
	return &searchService{
		c:    c,
		name: name,
	}
}

func (c *searchService) Match(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Search.Match", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Search service

type SearchHandler interface {
	// 搜索匹配
	Match(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterSearchHandler(s server.Server, hdlr SearchHandler, opts ...server.HandlerOption) error {
	type search interface {
		Match(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Search struct {
		search
	}
	h := &searchHandler{hdlr}
	return s.Handle(s.NewHandler(&Search{h}, opts...))
}

type searchHandler struct {
	SearchHandler
}

func (h *searchHandler) Match(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.SearchHandler.Match(ctx, in, out)
}
