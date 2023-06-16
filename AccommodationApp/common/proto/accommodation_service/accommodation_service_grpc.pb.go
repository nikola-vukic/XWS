// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accommodation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Get(ctx context.Context, in *GetAccommodationRequest, opts ...grpc.CallOption) (*GetAccommodationResponse, error)
	GetByHost(ctx context.Context, in *GetAccommodationRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error)
	GetAll(ctx context.Context, in *GetAllAccommodationsRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error)
	GetAllFiltered(ctx context.Context, in *GetAllFilterRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error)
	Create(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
	Update(ctx context.Context, in *UpdateAccommodationRequest, opts ...grpc.CallOption) (*UpdateAccommodationResponse, error)
	UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error)
	GetAccommodationAvailableDatesForTimePeriod(ctx context.Context, in *AccommodationTimePeriodRequest, opts ...grpc.CallOption) (*AccommodationAvailableDatesForTimePeriodResponse, error)
	Delete(ctx context.Context, in *DeleteAccommodationRequest, opts ...grpc.CallOption) (*DeleteAccommodationResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Get(ctx context.Context, in *GetAccommodationRequest, opts ...grpc.CallOption) (*GetAccommodationResponse, error) {
	out := new(GetAccommodationResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetByHost(ctx context.Context, in *GetAccommodationRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error) {
	out := new(GetAllAccommodationsResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/GetByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllAccommodationsRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error) {
	out := new(GetAllAccommodationsResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAllFiltered(ctx context.Context, in *GetAllFilterRequest, opts ...grpc.CallOption) (*GetAllAccommodationsResponse, error) {
	out := new(GetAllAccommodationsResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/GetAllFiltered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Create(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Update(ctx context.Context, in *UpdateAccommodationRequest, opts ...grpc.CallOption) (*UpdateAccommodationResponse, error) {
	out := new(UpdateAccommodationResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAvailability(ctx context.Context, in *UpdateAvailabilityRequest, opts ...grpc.CallOption) (*UpdateAvailabilityResponse, error) {
	out := new(UpdateAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/UpdateAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationAvailableDatesForTimePeriod(ctx context.Context, in *AccommodationTimePeriodRequest, opts ...grpc.CallOption) (*AccommodationAvailableDatesForTimePeriodResponse, error) {
	out := new(AccommodationAvailableDatesForTimePeriodResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/GetAccommodationAvailableDatesForTimePeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) Delete(ctx context.Context, in *DeleteAccommodationRequest, opts ...grpc.CallOption) (*DeleteAccommodationResponse, error) {
	out := new(DeleteAccommodationResponse)
	err := c.cc.Invoke(ctx, "/profile.AccommodationService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Get(context.Context, *GetAccommodationRequest) (*GetAccommodationResponse, error)
	GetByHost(context.Context, *GetAccommodationRequest) (*GetAllAccommodationsResponse, error)
	GetAll(context.Context, *GetAllAccommodationsRequest) (*GetAllAccommodationsResponse, error)
	GetAllFiltered(context.Context, *GetAllFilterRequest) (*GetAllAccommodationsResponse, error)
	Create(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error)
	Update(context.Context, *UpdateAccommodationRequest) (*UpdateAccommodationResponse, error)
	UpdateAvailability(context.Context, *UpdateAvailabilityRequest) (*UpdateAvailabilityResponse, error)
	GetAccommodationAvailableDatesForTimePeriod(context.Context, *AccommodationTimePeriodRequest) (*AccommodationAvailableDatesForTimePeriodResponse, error)
	Delete(context.Context, *DeleteAccommodationRequest) (*DeleteAccommodationResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (*UnimplementedAccommodationServiceServer) Get(context.Context, *GetAccommodationRequest) (*GetAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccommodationServiceServer) GetByHost(context.Context, *GetAccommodationRequest) (*GetAllAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByHost not implemented")
}
func (*UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllAccommodationsRequest) (*GetAllAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedAccommodationServiceServer) GetAllFiltered(context.Context, *GetAllFilterRequest) (*GetAllAccommodationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFiltered not implemented")
}
func (*UnimplementedAccommodationServiceServer) Create(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccommodationServiceServer) Update(context.Context, *UpdateAccommodationRequest) (*UpdateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAccommodationServiceServer) UpdateAvailability(context.Context, *UpdateAvailabilityRequest) (*UpdateAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailability not implemented")
}
func (*UnimplementedAccommodationServiceServer) GetAccommodationAvailableDatesForTimePeriod(context.Context, *AccommodationTimePeriodRequest) (*AccommodationAvailableDatesForTimePeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationAvailableDatesForTimePeriod not implemented")
}
func (*UnimplementedAccommodationServiceServer) Delete(context.Context, *DeleteAccommodationRequest) (*DeleteAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

func RegisterAccommodationServiceServer(s *grpc.Server, srv AccommodationServiceServer) {
	s.RegisterService(&_AccommodationService_serviceDesc, srv)
}

func _AccommodationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Get(ctx, req.(*GetAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/GetByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetByHost(ctx, req.(*GetAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAccommodationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllAccommodationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAllFiltered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAllFiltered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/GetAllFiltered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllFiltered(ctx, req.(*GetAllFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Create(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Update(ctx, req.(*UpdateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/UpdateAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAvailability(ctx, req.(*UpdateAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationAvailableDatesForTimePeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccommodationTimePeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationAvailableDatesForTimePeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/GetAccommodationAvailableDatesForTimePeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationAvailableDatesForTimePeriod(ctx, req.(*AccommodationTimePeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.AccommodationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Delete(ctx, req.(*DeleteAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccommodationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccommodationService_Get_Handler,
		},
		{
			MethodName: "GetByHost",
			Handler:    _AccommodationService_GetByHost_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "GetAllFiltered",
			Handler:    _AccommodationService_GetAllFiltered_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccommodationService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AccommodationService_Update_Handler,
		},
		{
			MethodName: "UpdateAvailability",
			Handler:    _AccommodationService_UpdateAvailability_Handler,
		},
		{
			MethodName: "GetAccommodationAvailableDatesForTimePeriod",
			Handler:    _AccommodationService_GetAccommodationAvailableDatesForTimePeriod_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccommodationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/accommodation_service/accommodation_service.proto",
}
