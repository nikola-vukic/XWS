// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package profile

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Get(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error)
	GetAll(ctx context.Context, in *GetAllReservationsRequest, opts ...grpc.CallOption) (*GetAllReservationsResponse, error)
	GetUsersReservations(ctx context.Context, in *GetUsersReservationsRequest, opts ...grpc.CallOption) (*GetUsersReservationsResponse, error)
	GetByHost(ctx context.Context, in *GetByHostRequest, opts ...grpc.CallOption) (*GetByHostResponse, error)
	GetMyReservations(ctx context.Context, in *GetMyReservationsRequest, opts ...grpc.CallOption) (*GetMyReservationsResponse, error)
	GetBetweenDates(ctx context.Context, in *GetBetweenDatesRequest, opts ...grpc.CallOption) (*GetBetweenDatesResponse, error)
	Create(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	Update(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error)
	Approve(ctx context.Context, in *ApproveReservationRequest, opts ...grpc.CallOption) (*ApproveReservationResponse, error)
	Reject(ctx context.Context, in *RejectReservationRequest, opts ...grpc.CallOption) (*RejectReservationResponse, error)
	Cancel(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error)
	Delete(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Get(ctx context.Context, in *GetReservationRequest, opts ...grpc.CallOption) (*GetReservationResponse, error) {
	out := new(GetReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllReservationsRequest, opts ...grpc.CallOption) (*GetAllReservationsResponse, error) {
	out := new(GetAllReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetUsersReservations(ctx context.Context, in *GetUsersReservationsRequest, opts ...grpc.CallOption) (*GetUsersReservationsResponse, error) {
	out := new(GetUsersReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetUsersReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByHost(ctx context.Context, in *GetByHostRequest, opts ...grpc.CallOption) (*GetByHostResponse, error) {
	out := new(GetByHostResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetMyReservations(ctx context.Context, in *GetMyReservationsRequest, opts ...grpc.CallOption) (*GetMyReservationsResponse, error) {
	out := new(GetMyReservationsResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetMyReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetBetweenDates(ctx context.Context, in *GetBetweenDatesRequest, opts ...grpc.CallOption) (*GetBetweenDatesResponse, error) {
	out := new(GetBetweenDatesResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/GetBetweenDates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Create(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Update(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*UpdateReservationResponse, error) {
	out := new(UpdateReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Approve(ctx context.Context, in *ApproveReservationRequest, opts ...grpc.CallOption) (*ApproveReservationResponse, error) {
	out := new(ApproveReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Approve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Reject(ctx context.Context, in *RejectReservationRequest, opts ...grpc.CallOption) (*RejectReservationResponse, error) {
	out := new(RejectReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Reject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Cancel(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error) {
	out := new(CancelReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) Delete(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, "/reservation.ReservationService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Get(context.Context, *GetReservationRequest) (*GetReservationResponse, error)
	GetAll(context.Context, *GetAllReservationsRequest) (*GetAllReservationsResponse, error)
	GetUsersReservations(context.Context, *GetUsersReservationsRequest) (*GetUsersReservationsResponse, error)
	GetByHost(context.Context, *GetByHostRequest) (*GetByHostResponse, error)
	GetMyReservations(context.Context, *GetMyReservationsRequest) (*GetMyReservationsResponse, error)
	GetBetweenDates(context.Context, *GetBetweenDatesRequest) (*GetBetweenDatesResponse, error)
	Create(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	Update(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error)
	Approve(context.Context, *ApproveReservationRequest) (*ApproveReservationResponse, error)
	Reject(context.Context, *RejectReservationRequest) (*RejectReservationResponse, error)
	Cancel(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error)
	Delete(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (*UnimplementedReservationServiceServer) Get(context.Context, *GetReservationRequest) (*GetReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllReservationsRequest) (*GetAllReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedReservationServiceServer) GetUsersReservations(context.Context, *GetUsersReservationsRequest) (*GetUsersReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersReservations not implemented")
}
func (*UnimplementedReservationServiceServer) GetByHost(context.Context, *GetByHostRequest) (*GetByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHost not implemented")
}
func (*UnimplementedReservationServiceServer) GetMyReservations(context.Context, *GetMyReservationsRequest) (*GetMyReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyReservations not implemented")
}
func (*UnimplementedReservationServiceServer) GetBetweenDates(context.Context, *GetBetweenDatesRequest) (*GetBetweenDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBetweenDates not implemented")
}
func (*UnimplementedReservationServiceServer) Create(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReservationServiceServer) Update(context.Context, *UpdateReservationRequest) (*UpdateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedReservationServiceServer) Approve(context.Context, *ApproveReservationRequest) (*ApproveReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (*UnimplementedReservationServiceServer) Reject(context.Context, *RejectReservationRequest) (*RejectReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reject not implemented")
}
func (*UnimplementedReservationServiceServer) Cancel(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedReservationServiceServer) Delete(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

func RegisterReservationServiceServer(s *grpc.Server, srv ReservationServiceServer) {
	s.RegisterService(&_ReservationService_serviceDesc, srv)
}

func _ReservationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Get(ctx, req.(*GetReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetUsersReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetUsersReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetUsersReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetUsersReservations(ctx, req.(*GetUsersReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByHost(ctx, req.(*GetByHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetMyReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetMyReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetMyReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetMyReservations(ctx, req.(*GetMyReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetBetweenDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBetweenDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetBetweenDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/GetBetweenDates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetBetweenDates(ctx, req.(*GetBetweenDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Create(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Update(ctx, req.(*UpdateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Approve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Approve(ctx, req.(*ApproveReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Reject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Reject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Reject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Reject(ctx, req.(*RejectReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Cancel(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.ReservationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Delete(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReservationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReservationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "GetUsersReservations",
			Handler:    _ReservationService_GetUsersReservations_Handler,
		},
		{
			MethodName: "GetByHost",
			Handler:    _ReservationService_GetByHost_Handler,
		},
		{
			MethodName: "GetMyReservations",
			Handler:    _ReservationService_GetMyReservations_Handler,
		},
		{
			MethodName: "GetBetweenDates",
			Handler:    _ReservationService_GetBetweenDates_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ReservationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReservationService_Update_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _ReservationService_Approve_Handler,
		},
		{
			MethodName: "Reject",
			Handler:    _ReservationService_Reject_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _ReservationService_Cancel_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReservationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reservation_service/reservation_service.proto",
}
