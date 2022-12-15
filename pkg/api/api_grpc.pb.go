// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CipactonalClient is the client API for Cipactonal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CipactonalClient interface {
	// Aggregated category scores over a period of time
	CategoryScores(ctx context.Context, in *CategoryScoresRequest, opts ...grpc.CallOption) (*CategoryScoresResponse, error)
	// Aggregate scores for categories within defined period by ticket.
	TicketScores(ctx context.Context, in *TicketScoresRequest, opts ...grpc.CallOption) (*TicketScoresResponse, error)
	// Overal quality score
	QualityScore(ctx context.Context, in *QualityScoreRequest, opts ...grpc.CallOption) (*QualityScoreResponse, error)
	// Period over Period score change
	ScoreChange(ctx context.Context, in *ScoreChangeRequest, opts ...grpc.CallOption) (*ScoreChangeResponse, error)
}

type cipactonalClient struct {
	cc grpc.ClientConnInterface
}

func NewCipactonalClient(cc grpc.ClientConnInterface) CipactonalClient {
	return &cipactonalClient{cc}
}

func (c *cipactonalClient) CategoryScores(ctx context.Context, in *CategoryScoresRequest, opts ...grpc.CallOption) (*CategoryScoresResponse, error) {
	out := new(CategoryScoresResponse)
	err := c.cc.Invoke(ctx, "/Cipactonal/CategoryScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipactonalClient) TicketScores(ctx context.Context, in *TicketScoresRequest, opts ...grpc.CallOption) (*TicketScoresResponse, error) {
	out := new(TicketScoresResponse)
	err := c.cc.Invoke(ctx, "/Cipactonal/TicketScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipactonalClient) QualityScore(ctx context.Context, in *QualityScoreRequest, opts ...grpc.CallOption) (*QualityScoreResponse, error) {
	out := new(QualityScoreResponse)
	err := c.cc.Invoke(ctx, "/Cipactonal/QualityScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipactonalClient) ScoreChange(ctx context.Context, in *ScoreChangeRequest, opts ...grpc.CallOption) (*ScoreChangeResponse, error) {
	out := new(ScoreChangeResponse)
	err := c.cc.Invoke(ctx, "/Cipactonal/ScoreChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CipactonalServer is the server API for Cipactonal service.
// All implementations should embed UnimplementedCipactonalServer
// for forward compatibility
type CipactonalServer interface {
	// Aggregated category scores over a period of time
	CategoryScores(context.Context, *CategoryScoresRequest) (*CategoryScoresResponse, error)
	// Aggregate scores for categories within defined period by ticket.
	TicketScores(context.Context, *TicketScoresRequest) (*TicketScoresResponse, error)
	// Overal quality score
	QualityScore(context.Context, *QualityScoreRequest) (*QualityScoreResponse, error)
	// Period over Period score change
	ScoreChange(context.Context, *ScoreChangeRequest) (*ScoreChangeResponse, error)
}

// UnimplementedCipactonalServer should be embedded to have forward compatible implementations.
type UnimplementedCipactonalServer struct {
}

func (UnimplementedCipactonalServer) CategoryScores(context.Context, *CategoryScoresRequest) (*CategoryScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryScores not implemented")
}
func (UnimplementedCipactonalServer) TicketScores(context.Context, *TicketScoresRequest) (*TicketScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TicketScores not implemented")
}
func (UnimplementedCipactonalServer) QualityScore(context.Context, *QualityScoreRequest) (*QualityScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QualityScore not implemented")
}
func (UnimplementedCipactonalServer) ScoreChange(context.Context, *ScoreChangeRequest) (*ScoreChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScoreChange not implemented")
}

// UnsafeCipactonalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CipactonalServer will
// result in compilation errors.
type UnsafeCipactonalServer interface {
	mustEmbedUnimplementedCipactonalServer()
}

func RegisterCipactonalServer(s grpc.ServiceRegistrar, srv CipactonalServer) {
	s.RegisterService(&Cipactonal_ServiceDesc, srv)
}

func _Cipactonal_CategoryScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipactonalServer).CategoryScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cipactonal/CategoryScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipactonalServer).CategoryScores(ctx, req.(*CategoryScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cipactonal_TicketScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipactonalServer).TicketScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cipactonal/TicketScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipactonalServer).TicketScores(ctx, req.(*TicketScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cipactonal_QualityScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QualityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipactonalServer).QualityScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cipactonal/QualityScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipactonalServer).QualityScore(ctx, req.(*QualityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cipactonal_ScoreChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipactonalServer).ScoreChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cipactonal/ScoreChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipactonalServer).ScoreChange(ctx, req.(*ScoreChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cipactonal_ServiceDesc is the grpc.ServiceDesc for Cipactonal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cipactonal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Cipactonal",
	HandlerType: (*CipactonalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategoryScores",
			Handler:    _Cipactonal_CategoryScores_Handler,
		},
		{
			MethodName: "TicketScores",
			Handler:    _Cipactonal_TicketScores_Handler,
		},
		{
			MethodName: "QualityScore",
			Handler:    _Cipactonal_QualityScore_Handler,
		},
		{
			MethodName: "ScoreChange",
			Handler:    _Cipactonal_ScoreChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}