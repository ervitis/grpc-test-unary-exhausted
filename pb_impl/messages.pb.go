// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/messages.proto

package pb_impl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string title = 1;
	// int32 total_size = 2;
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DataRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendNotificationBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Personalizations []*Personalization `protobuf:"bytes,1,rep,name=personalizations,proto3" json:"personalizations,omitempty"`
	Text             string             `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SendNotificationBatchRequest) Reset() {
	*x = SendNotificationBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotificationBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotificationBatchRequest) ProtoMessage() {}

func (x *SendNotificationBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotificationBatchRequest.ProtoReflect.Descriptor instead.
func (*SendNotificationBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{1}
}

func (x *SendNotificationBatchRequest) GetPersonalizations() []*Personalization {
	if x != nil {
		return x.Personalizations
	}
	return nil
}

func (x *SendNotificationBatchRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Personalization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email         string            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Substitutions map[string]string `protobuf:"bytes,2,rep,name=substitutions,proto3" json:"substitutions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Personalization) Reset() {
	*x = Personalization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Personalization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Personalization) ProtoMessage() {}

func (x *Personalization) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Personalization.ProtoReflect.Descriptor instead.
func (*Personalization) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Personalization) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Personalization) GetSubstitutions() map[string]string {
	if x != nil {
		return x.Substitutions
	}
	return nil
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_proto_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_messages_proto protoreflect.FileDescriptor

var file_proto_messages_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x78, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x5f, 0x69,
	0x6d, 0x70, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xbc, 0x01, 0x0a, 0x0f, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x51, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x62, 0x5f,
	0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xd1, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x69, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x37, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70,
	0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x32, 0x4b, 0x0a, 0x0d, 0x4e, 0x6f, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x32, 0x4b, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x65,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_messages_proto_rawDescOnce sync.Once
	file_proto_messages_proto_rawDescData = file_proto_messages_proto_rawDesc
)

func file_proto_messages_proto_rawDescGZIP() []byte {
	file_proto_messages_proto_rawDescOnce.Do(func() {
		file_proto_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_messages_proto_rawDescData)
	})
	return file_proto_messages_proto_rawDescData
}

var file_proto_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_messages_proto_goTypes = []interface{}{
	(*DataRequest)(nil),                  // 0: pb_impl.DataRequest
	(*SendNotificationBatchRequest)(nil), // 1: pb_impl.SendNotificationBatchRequest
	(*Personalization)(nil),              // 2: pb_impl.Personalization
	(*DataResponse)(nil),                 // 3: pb_impl.DataResponse
	nil,                                  // 4: pb_impl.Personalization.SubstitutionsEntry
	(*emptypb.Empty)(nil),                // 5: google.protobuf.Empty
}
var file_proto_messages_proto_depIdxs = []int32{
	2, // 0: pb_impl.SendNotificationBatchRequest.personalizations:type_name -> pb_impl.Personalization
	4, // 1: pb_impl.Personalization.substitutions:type_name -> pb_impl.Personalization.SubstitutionsEntry
	1, // 2: pb_impl.DataService.SendBigData:input_type -> pb_impl.SendNotificationBatchRequest
	0, // 3: pb_impl.DataService.SendData:input_type -> pb_impl.DataRequest
	0, // 4: pb_impl.DataService.SendStream:input_type -> pb_impl.DataRequest
	5, // 5: pb_impl.NoAuthService.SayHello:input_type -> google.protobuf.Empty
	5, // 6: pb_impl.AuthService.SayGoodBye:input_type -> google.protobuf.Empty
	5, // 7: pb_impl.DataService.SendBigData:output_type -> google.protobuf.Empty
	3, // 8: pb_impl.DataService.SendData:output_type -> pb_impl.DataResponse
	3, // 9: pb_impl.DataService.SendStream:output_type -> pb_impl.DataResponse
	5, // 10: pb_impl.NoAuthService.SayHello:output_type -> google.protobuf.Empty
	5, // 11: pb_impl.AuthService.SayGoodBye:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_messages_proto_init() }
func file_proto_messages_proto_init() {
	if File_proto_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_proto_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotificationBatchRequest); i {
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
		file_proto_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Personalization); i {
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
		file_proto_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
			RawDescriptor: file_proto_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proto_messages_proto_goTypes,
		DependencyIndexes: file_proto_messages_proto_depIdxs,
		MessageInfos:      file_proto_messages_proto_msgTypes,
	}.Build()
	File_proto_messages_proto = out.File
	file_proto_messages_proto_rawDesc = nil
	file_proto_messages_proto_goTypes = nil
	file_proto_messages_proto_depIdxs = nil
}
