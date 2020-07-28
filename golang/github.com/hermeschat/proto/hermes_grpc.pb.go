// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HermesClient is the client API for Hermes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HermesClient interface {
	EventBuff(ctx context.Context, opts ...grpc.CallOption) (Hermes_EventBuffClient, error)
	Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ListChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Channels, error)
	GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error)
}

type hermesClient struct {
	cc grpc.ClientConnInterface
}

func NewHermesClient(cc grpc.ClientConnInterface) HermesClient {
	return &hermesClient{cc}
}

func (c *hermesClient) EventBuff(ctx context.Context, opts ...grpc.CallOption) (Hermes_EventBuffClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hermes_serviceDesc.Streams[0], "/Hermes/EventBuff", opts...)
	if err != nil {
		return nil, err
	}
	x := &hermesEventBuffClient{stream}
	return x, nil
}

type Hermes_EventBuffClient interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type hermesEventBuffClient struct {
	grpc.ClientStream
}

func (x *hermesEventBuffClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hermesEventBuffClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hermesClient) Echo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Hermes/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesClient) ListChannels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Channels, error) {
	out := new(Channels)
	err := c.cc.Invoke(ctx, "/Hermes/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hermesClient) GetChannel(ctx context.Context, in *GetChannelRequest, opts ...grpc.CallOption) (*Channel, error) {
	out := new(Channel)
	err := c.cc.Invoke(ctx, "/Hermes/GetChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HermesServer is the server API for Hermes service.
// All implementations must embed UnimplementedHermesServer
// for forward compatibility
type HermesServer interface {
	EventBuff(Hermes_EventBuffServer) error
	Echo(context.Context, *Empty) (*Empty, error)
	ListChannels(context.Context, *Empty) (*Channels, error)
	GetChannel(context.Context, *GetChannelRequest) (*Channel, error)
	mustEmbedUnimplementedHermesServer()
}

// UnimplementedHermesServer must be embedded to have forward compatible implementations.
type UnimplementedHermesServer struct {
}

func (*UnimplementedHermesServer) EventBuff(Hermes_EventBuffServer) error {
	return status.Errorf(codes.Unimplemented, "method EventBuff not implemented")
}
func (*UnimplementedHermesServer) Echo(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedHermesServer) ListChannels(context.Context, *Empty) (*Channels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (*UnimplementedHermesServer) GetChannel(context.Context, *GetChannelRequest) (*Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannel not implemented")
}
func (*UnimplementedHermesServer) mustEmbedUnimplementedHermesServer() {}

func RegisterHermesServer(s *grpc.Server, srv HermesServer) {
	s.RegisterService(&_Hermes_serviceDesc, srv)
}

func _Hermes_EventBuff_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HermesServer).EventBuff(&hermesEventBuffServer{stream})
}

type Hermes_EventBuffServer interface {
	Send(*Event) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type hermesEventBuffServer struct {
	grpc.ServerStream
}

func (x *hermesEventBuffServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hermesEventBuffServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hermes_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hermes/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServer).Echo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hermes_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hermes/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServer).ListChannels(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hermes_GetChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HermesServer).GetChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hermes/GetChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HermesServer).GetChannel(ctx, req.(*GetChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hermes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hermes",
	HandlerType: (*HermesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Hermes_Echo_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _Hermes_ListChannels_Handler,
		},
		{
			MethodName: "GetChannel",
			Handler:    _Hermes_GetChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventBuff",
			Handler:       _Hermes_EventBuff_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hermes.proto",
}
