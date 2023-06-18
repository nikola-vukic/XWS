// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/grade_service/grade_service.proto

package grade

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGradeRequest) Reset() {
	*x = GetGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGradeRequest) ProtoMessage() {}

func (x *GetGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGradeRequest.ProtoReflect.Descriptor instead.
func (*GetGradeRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGradeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade *Grade `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *GetGradeResponse) Reset() {
	*x = GetGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGradeResponse) ProtoMessage() {}

func (x *GetGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGradeResponse.ProtoReflect.Descriptor instead.
func (*GetGradeResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetGradeResponse) GetGrade() *Grade {
	if x != nil {
		return x.Grade
	}
	return nil
}

type GetAllGradesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllGradesRequest) Reset() {
	*x = GetAllGradesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGradesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGradesRequest) ProtoMessage() {}

func (x *GetAllGradesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGradesRequest.ProtoReflect.Descriptor instead.
func (*GetAllGradesRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{2}
}

type GetAllGradesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grades []*Grade `protobuf:"bytes,1,rep,name=grades,proto3" json:"grades,omitempty"`
}

func (x *GetAllGradesResponse) Reset() {
	*x = GetAllGradesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGradesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGradesResponse) ProtoMessage() {}

func (x *GetAllGradesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGradesResponse.ProtoReflect.Descriptor instead.
func (*GetAllGradesResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllGradesResponse) GetGrades() []*Grade {
	if x != nil {
		return x.Grades
	}
	return nil
}

type CreateGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade *NewGrade `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *CreateGradeRequest) Reset() {
	*x = CreateGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGradeRequest) ProtoMessage() {}

func (x *CreateGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGradeRequest.ProtoReflect.Descriptor instead.
func (*CreateGradeRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGradeRequest) GetGrade() *NewGrade {
	if x != nil {
		return x.Grade
	}
	return nil
}

type CreateGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade *Grade `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *CreateGradeResponse) Reset() {
	*x = CreateGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGradeResponse) ProtoMessage() {}

func (x *CreateGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGradeResponse.ProtoReflect.Descriptor instead.
func (*CreateGradeResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGradeResponse) GetGrade() *Grade {
	if x != nil {
		return x.Grade
	}
	return nil
}

type UpdateGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateGradeRequest) Reset() {
	*x = UpdateGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGradeRequest) ProtoMessage() {}

func (x *UpdateGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGradeRequest.ProtoReflect.Descriptor instead.
func (*UpdateGradeRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGradeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGradeRequest) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type UpdateGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grade *Grade `protobuf:"bytes,1,opt,name=grade,proto3" json:"grade,omitempty"`
}

func (x *UpdateGradeResponse) Reset() {
	*x = UpdateGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGradeResponse) ProtoMessage() {}

func (x *UpdateGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGradeResponse.ProtoReflect.Descriptor instead.
func (*UpdateGradeResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGradeResponse) GetGrade() *Grade {
	if x != nil {
		return x.Grade
	}
	return nil
}

type DeleteGradeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGradeRequest) Reset() {
	*x = DeleteGradeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGradeRequest) ProtoMessage() {}

func (x *DeleteGradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGradeRequest.ProtoReflect.Descriptor instead.
func (*DeleteGradeRequest) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGradeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGradeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGradeResponse) Reset() {
	*x = DeleteGradeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGradeResponse) ProtoMessage() {}

func (x *DeleteGradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGradeResponse.ProtoReflect.Descriptor instead.
func (*DeleteGradeResponse) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{9}
}

type Grade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GuestId     string                 `protobuf:"bytes,2,opt,name=guestId,proto3" json:"guestId,omitempty"`
	GradedName  string                 `protobuf:"bytes,3,opt,name=gradedName,proto3" json:"gradedName,omitempty"`
	Grade       float32                `protobuf:"fixed32,4,opt,name=grade,proto3" json:"grade,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	IsHostGrade bool                   `protobuf:"varint,6,opt,name=isHostGrade,proto3" json:"isHostGrade,omitempty"`
}

func (x *Grade) Reset() {
	*x = Grade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grade) ProtoMessage() {}

func (x *Grade) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grade.ProtoReflect.Descriptor instead.
func (*Grade) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{10}
}

func (x *Grade) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Grade) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *Grade) GetGradedName() string {
	if x != nil {
		return x.GradedName
	}
	return ""
}

func (x *Grade) GetGrade() float32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Grade) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Grade) GetIsHostGrade() bool {
	if x != nil {
		return x.IsHostGrade
	}
	return false
}

type NewGrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuestId     string  `protobuf:"bytes,1,opt,name=guestId,proto3" json:"guestId,omitempty"`
	GradedId    string  `protobuf:"bytes,2,opt,name=gradedId,proto3" json:"gradedId,omitempty"`
	Value       float32 `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	IsHostGrade bool    `protobuf:"varint,4,opt,name=isHostGrade,proto3" json:"isHostGrade,omitempty"`
}

func (x *NewGrade) Reset() {
	*x = NewGrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grade_service_grade_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGrade) ProtoMessage() {}

func (x *NewGrade) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grade_service_grade_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGrade.ProtoReflect.Descriptor instead.
func (*NewGrade) Descriptor() ([]byte, []int) {
	return file_proto_grade_service_grade_service_proto_rawDescGZIP(), []int{11}
}

func (x *NewGrade) GetGuestId() string {
	if x != nil {
		return x.GuestId
	}
	return ""
}

func (x *NewGrade) GetGradedId() string {
	if x != nil {
		return x.GradedId
	}
	return ""
}

func (x *NewGrade) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *NewGrade) GetIsHostGrade() bool {
	if x != nil {
		return x.IsHostGrade
	}
	return false
}

var File_proto_grade_service_grade_service_proto protoreflect.FileDescriptor

var file_proto_grade_service_grade_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x06, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x22,
	0x3b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x4e, 0x65, 0x77,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x39, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x22, 0x24,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x05,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x48, 0x6f,
	0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x22, 0x78, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x32, 0x85, 0x06, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x6d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x64, 0x42, 0x79, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x7f,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65, 0x64, 0x42, 0x79, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x72, 0x61, 0x64, 0x65, 0x64, 0x12, 0x16,
	0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x06, 0x2f, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x3a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x1a, 0x0b, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grade_service_grade_service_proto_rawDescOnce sync.Once
	file_proto_grade_service_grade_service_proto_rawDescData = file_proto_grade_service_grade_service_proto_rawDesc
)

func file_proto_grade_service_grade_service_proto_rawDescGZIP() []byte {
	file_proto_grade_service_grade_service_proto_rawDescOnce.Do(func() {
		file_proto_grade_service_grade_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grade_service_grade_service_proto_rawDescData)
	})
	return file_proto_grade_service_grade_service_proto_rawDescData
}

var file_proto_grade_service_grade_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_grade_service_grade_service_proto_goTypes = []interface{}{
	(*GetGradeRequest)(nil),       // 0: grade.GetGradeRequest
	(*GetGradeResponse)(nil),      // 1: grade.GetGradeResponse
	(*GetAllGradesRequest)(nil),   // 2: grade.GetAllGradesRequest
	(*GetAllGradesResponse)(nil),  // 3: grade.GetAllGradesResponse
	(*CreateGradeRequest)(nil),    // 4: grade.CreateGradeRequest
	(*CreateGradeResponse)(nil),   // 5: grade.CreateGradeResponse
	(*UpdateGradeRequest)(nil),    // 6: grade.UpdateGradeRequest
	(*UpdateGradeResponse)(nil),   // 7: grade.UpdateGradeResponse
	(*DeleteGradeRequest)(nil),    // 8: grade.DeleteGradeRequest
	(*DeleteGradeResponse)(nil),   // 9: grade.DeleteGradeResponse
	(*Grade)(nil),                 // 10: grade.Grade
	(*NewGrade)(nil),              // 11: grade.NewGrade
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_proto_grade_service_grade_service_proto_depIdxs = []int32{
	10, // 0: grade.GetGradeResponse.grade:type_name -> grade.Grade
	10, // 1: grade.GetAllGradesResponse.grades:type_name -> grade.Grade
	11, // 2: grade.CreateGradeRequest.grade:type_name -> grade.NewGrade
	10, // 3: grade.CreateGradeResponse.grade:type_name -> grade.Grade
	10, // 4: grade.UpdateGradeResponse.grade:type_name -> grade.Grade
	12, // 5: grade.Grade.date:type_name -> google.protobuf.Timestamp
	0,  // 6: grade.GradeService.Get:input_type -> grade.GetGradeRequest
	0,  // 7: grade.GradeService.GetHostsGradedByGuest:input_type -> grade.GetGradeRequest
	0,  // 8: grade.GradeService.GetAccommodationsGradedByGuest:input_type -> grade.GetGradeRequest
	0,  // 9: grade.GradeService.GetByGraded:input_type -> grade.GetGradeRequest
	2,  // 10: grade.GradeService.GetAll:input_type -> grade.GetAllGradesRequest
	4,  // 11: grade.GradeService.Create:input_type -> grade.CreateGradeRequest
	6,  // 12: grade.GradeService.Update:input_type -> grade.UpdateGradeRequest
	8,  // 13: grade.GradeService.Delete:input_type -> grade.DeleteGradeRequest
	1,  // 14: grade.GradeService.Get:output_type -> grade.GetGradeResponse
	3,  // 15: grade.GradeService.GetHostsGradedByGuest:output_type -> grade.GetAllGradesResponse
	3,  // 16: grade.GradeService.GetAccommodationsGradedByGuest:output_type -> grade.GetAllGradesResponse
	3,  // 17: grade.GradeService.GetByGraded:output_type -> grade.GetAllGradesResponse
	3,  // 18: grade.GradeService.GetAll:output_type -> grade.GetAllGradesResponse
	5,  // 19: grade.GradeService.Create:output_type -> grade.CreateGradeResponse
	7,  // 20: grade.GradeService.Update:output_type -> grade.UpdateGradeResponse
	9,  // 21: grade.GradeService.Delete:output_type -> grade.DeleteGradeResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_grade_service_grade_service_proto_init() }
func file_proto_grade_service_grade_service_proto_init() {
	if File_proto_grade_service_grade_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grade_service_grade_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGradeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGradeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGradesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGradesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGradeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGradeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGradeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGradeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGradeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGradeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grade); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grade_service_grade_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGrade); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_grade_service_grade_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grade_service_grade_service_proto_goTypes,
		DependencyIndexes: file_proto_grade_service_grade_service_proto_depIdxs,
		MessageInfos:      file_proto_grade_service_grade_service_proto_msgTypes,
	}.Build()
	File_proto_grade_service_grade_service_proto = out.File
	file_proto_grade_service_grade_service_proto_rawDesc = nil
	file_proto_grade_service_grade_service_proto_goTypes = nil
	file_proto_grade_service_grade_service_proto_depIdxs = nil
}
