// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/examples/template/api/proto/example/example.proto

/*
Package go_micro_api_template is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/template/api/proto/example/example.proto

It has these top-level messages:
*/
package go_micro_api_template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleService interface {
	Call(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type exampleService struct {
	c           client.Client
	serviceName string
}

func ExampleServiceClient(serviceName string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.api.template"
	}
	return &exampleService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleService) Call(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Example{hdlr}, opts...))
}

type Example struct {
	ExampleHandler
}

func (h *Example) Call(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ExampleHandler.Call(ctx, in, out)
}
