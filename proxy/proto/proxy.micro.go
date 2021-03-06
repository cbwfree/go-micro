// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/go-proxy/proto/proxy.proto

/*
Package proxy is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-proxy/proto/proxy.proto

It has these top-level messages:
	Request
	Response
	Message
	Empty
*/
package proxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Service service

type Service interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, opts ...client.CallOption) (Service_StreamService, error)
	Publish(ctx context.Context, in *Message, opts ...client.CallOption) (*Empty, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proxy"
	}
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Stream(ctx context.Context, opts ...client.CallOption) (Service_StreamService, error) {
	req := c.c.NewRequest(c.name, "Service.Stream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &serviceStream{stream}, nil
}

type Service_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type serviceStream struct {
	stream client.Stream
}

func (x *serviceStream) Close() error {
	return x.stream.Close()
}

func (x *serviceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *serviceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *serviceStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *serviceStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *service) Publish(ctx context.Context, in *Message, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Service.Publish", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, Service_StreamStream) error
	Publish(context.Context, *Message, *Empty) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) error {
	type service interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		Publish(ctx context.Context, in *Message, out *Empty) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.Call(ctx, in, out)
}

func (h *serviceHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.ServiceHandler.Stream(ctx, &serviceStreamStream{stream})
}

type Service_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type serviceStreamStream struct {
	stream server.Stream
}

func (x *serviceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *serviceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *serviceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *serviceStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *serviceStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *serviceHandler) Publish(ctx context.Context, in *Message, out *Empty) error {
	return h.ServiceHandler.Publish(ctx, in, out)
}
