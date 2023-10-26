// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: Gateway/proto/veterinary.proto

package pb

import (
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

type MakeAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetId    string                 `protobuf:"bytes,1,opt,name=petId,proto3" json:"petId,omitempty"`
	DateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
}

func (x *MakeAppointmentRequest) Reset() {
	*x = MakeAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Gateway_proto_veterinary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeAppointmentRequest) ProtoMessage() {}

func (x *MakeAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Gateway_proto_veterinary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeAppointmentRequest.ProtoReflect.Descriptor instead.
func (*MakeAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_Gateway_proto_veterinary_proto_rawDescGZIP(), []int{0}
}

func (x *MakeAppointmentRequest) GetPetId() string {
	if x != nil {
		return x.PetId
	}
	return ""
}

func (x *MakeAppointmentRequest) GetDateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DateTime
	}
	return nil
}

type MakeAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MakeAppointmentResponse) Reset() {
	*x = MakeAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Gateway_proto_veterinary_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeAppointmentResponse) ProtoMessage() {}

func (x *MakeAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Gateway_proto_veterinary_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeAppointmentResponse.ProtoReflect.Descriptor instead.
func (*MakeAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_Gateway_proto_veterinary_proto_rawDescGZIP(), []int{1}
}

func (x *MakeAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EndAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppointmentId string `protobuf:"bytes,1,opt,name=appointmentId,proto3" json:"appointmentId,omitempty"`
	Details       string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *EndAppointmentRequest) Reset() {
	*x = EndAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Gateway_proto_veterinary_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndAppointmentRequest) ProtoMessage() {}

func (x *EndAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Gateway_proto_veterinary_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndAppointmentRequest.ProtoReflect.Descriptor instead.
func (*EndAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_Gateway_proto_veterinary_proto_rawDescGZIP(), []int{2}
}

func (x *EndAppointmentRequest) GetAppointmentId() string {
	if x != nil {
		return x.AppointmentId
	}
	return ""
}

func (x *EndAppointmentRequest) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type EndAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EndAppointmentResponse) Reset() {
	*x = EndAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Gateway_proto_veterinary_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndAppointmentResponse) ProtoMessage() {}

func (x *EndAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Gateway_proto_veterinary_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndAppointmentResponse.ProtoReflect.Descriptor instead.
func (*EndAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_Gateway_proto_veterinary_proto_rawDescGZIP(), []int{3}
}

func (x *EndAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Gateway_proto_veterinary_proto protoreflect.FileDescriptor

var file_Gateway_proto_veterinary_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x66, 0x0a, 0x16, 0x4d, 0x61, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x17, 0x4d, 0x61, 0x6b,
	0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57,
	0x0a, 0x15, 0x45, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x32, 0x0a, 0x16, 0x45, 0x6e, 0x64, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x0a,
	0x56, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x44, 0x0a, 0x0f, 0x4d, 0x61,
	0x6b, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e,
	0x4d, 0x61, 0x6b, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x41, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x45, 0x6e, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x45, 0x6e, 0x64,
	0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x76, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Gateway_proto_veterinary_proto_rawDescOnce sync.Once
	file_Gateway_proto_veterinary_proto_rawDescData = file_Gateway_proto_veterinary_proto_rawDesc
)

func file_Gateway_proto_veterinary_proto_rawDescGZIP() []byte {
	file_Gateway_proto_veterinary_proto_rawDescOnce.Do(func() {
		file_Gateway_proto_veterinary_proto_rawDescData = protoimpl.X.CompressGZIP(file_Gateway_proto_veterinary_proto_rawDescData)
	})
	return file_Gateway_proto_veterinary_proto_rawDescData
}

var file_Gateway_proto_veterinary_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Gateway_proto_veterinary_proto_goTypes = []interface{}{
	(*MakeAppointmentRequest)(nil),  // 0: MakeAppointmentRequest
	(*MakeAppointmentResponse)(nil), // 1: MakeAppointmentResponse
	(*EndAppointmentRequest)(nil),   // 2: EndAppointmentRequest
	(*EndAppointmentResponse)(nil),  // 3: EndAppointmentResponse
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_Gateway_proto_veterinary_proto_depIdxs = []int32{
	4, // 0: MakeAppointmentRequest.dateTime:type_name -> google.protobuf.Timestamp
	0, // 1: Veterinary.MakeAppointment:input_type -> MakeAppointmentRequest
	2, // 2: Veterinary.EndAppointment:input_type -> EndAppointmentRequest
	1, // 3: Veterinary.MakeAppointment:output_type -> MakeAppointmentResponse
	3, // 4: Veterinary.EndAppointment:output_type -> EndAppointmentResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Gateway_proto_veterinary_proto_init() }
func file_Gateway_proto_veterinary_proto_init() {
	if File_Gateway_proto_veterinary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Gateway_proto_veterinary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeAppointmentRequest); i {
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
		file_Gateway_proto_veterinary_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeAppointmentResponse); i {
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
		file_Gateway_proto_veterinary_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndAppointmentRequest); i {
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
		file_Gateway_proto_veterinary_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndAppointmentResponse); i {
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
			RawDescriptor: file_Gateway_proto_veterinary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Gateway_proto_veterinary_proto_goTypes,
		DependencyIndexes: file_Gateway_proto_veterinary_proto_depIdxs,
		MessageInfos:      file_Gateway_proto_veterinary_proto_msgTypes,
	}.Build()
	File_Gateway_proto_veterinary_proto = out.File
	file_Gateway_proto_veterinary_proto_rawDesc = nil
	file_Gateway_proto_veterinary_proto_goTypes = nil
	file_Gateway_proto_veterinary_proto_depIdxs = nil
}
