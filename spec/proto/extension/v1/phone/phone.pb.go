// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spec/proto/extension/v1/phone/phone.proto

package phone

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request of SendVoiceWithTemplate method
type SendVoiceWithTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If your system uses multiple IVR services at the same time,
	// you can specify which service to use with this field.
	ComponentName string `protobuf:"bytes,1,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	// Required
	Template *VoiceTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Required
	ToMobile []string `protobuf:"bytes,3,rep,name=to_mobile,json=toMobile,proto3" json:"to_mobile,omitempty"`
	// This field is required by some cloud providers.
	FromMobile string `protobuf:"bytes,4,opt,name=from_mobile,json=fromMobile,proto3" json:"from_mobile,omitempty"`
}

func (x *SendVoiceWithTemplateRequest) Reset() {
	*x = SendVoiceWithTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVoiceWithTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVoiceWithTemplateRequest) ProtoMessage() {}

func (x *SendVoiceWithTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVoiceWithTemplateRequest.ProtoReflect.Descriptor instead.
func (*SendVoiceWithTemplateRequest) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_phone_phone_proto_rawDescGZIP(), []int{0}
}

func (x *SendVoiceWithTemplateRequest) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *SendVoiceWithTemplateRequest) GetTemplate() *VoiceTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *SendVoiceWithTemplateRequest) GetToMobile() []string {
	if x != nil {
		return x.ToMobile
	}
	return nil
}

func (x *SendVoiceWithTemplateRequest) GetFromMobile() string {
	if x != nil {
		return x.FromMobile
	}
	return ""
}

// VoiceTemplate
type VoiceTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	TemplateId string `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// Required
	TemplateParams map[string]string `protobuf:"bytes,2,rep,name=template_params,json=templateParams,proto3" json:"template_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VoiceTemplate) Reset() {
	*x = VoiceTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceTemplate) ProtoMessage() {}

func (x *VoiceTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceTemplate.ProtoReflect.Descriptor instead.
func (*VoiceTemplate) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_phone_phone_proto_rawDescGZIP(), []int{1}
}

func (x *VoiceTemplate) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *VoiceTemplate) GetTemplateParams() map[string]string {
	if x != nil {
		return x.TemplateParams
	}
	return nil
}

// The response of `SendVoiceWithTemplate` method
type SendVoiceWithTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of this request.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *SendVoiceWithTemplateResponse) Reset() {
	*x = SendVoiceWithTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVoiceWithTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVoiceWithTemplateResponse) ProtoMessage() {}

func (x *SendVoiceWithTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spec_proto_extension_v1_phone_phone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVoiceWithTemplateResponse.ProtoReflect.Descriptor instead.
func (*SendVoiceWithTemplateResponse) Descriptor() ([]byte, []int) {
	return file_spec_proto_extension_v1_phone_phone_proto_rawDescGZIP(), []int{2}
}

func (x *SendVoiceWithTemplateResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_spec_proto_extension_v1_phone_phone_proto protoreflect.FileDescriptor

var file_spec_proto_extension_v1_phone_phone_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x1c, 0x53,
	0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x6f, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x6f, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x0d, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x69, 0x0a,
	0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x1d, 0x53,
	0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x32, 0xa9, 0x01, 0x0a, 0x10,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x94, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x6d, 0x6f, 0x73, 0x6e, 0x2e,
	0x69, 0x6f, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x74, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x3b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_proto_extension_v1_phone_phone_proto_rawDescOnce sync.Once
	file_spec_proto_extension_v1_phone_phone_proto_rawDescData = file_spec_proto_extension_v1_phone_phone_proto_rawDesc
)

func file_spec_proto_extension_v1_phone_phone_proto_rawDescGZIP() []byte {
	file_spec_proto_extension_v1_phone_phone_proto_rawDescOnce.Do(func() {
		file_spec_proto_extension_v1_phone_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_proto_extension_v1_phone_phone_proto_rawDescData)
	})
	return file_spec_proto_extension_v1_phone_phone_proto_rawDescData
}

var file_spec_proto_extension_v1_phone_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spec_proto_extension_v1_phone_phone_proto_goTypes = []interface{}{
	(*SendVoiceWithTemplateRequest)(nil),  // 0: spec.proto.extension.v1.phone.SendVoiceWithTemplateRequest
	(*VoiceTemplate)(nil),                 // 1: spec.proto.extension.v1.phone.VoiceTemplate
	(*SendVoiceWithTemplateResponse)(nil), // 2: spec.proto.extension.v1.phone.SendVoiceWithTemplateResponse
	nil,                                   // 3: spec.proto.extension.v1.phone.VoiceTemplate.TemplateParamsEntry
}
var file_spec_proto_extension_v1_phone_phone_proto_depIdxs = []int32{
	1, // 0: spec.proto.extension.v1.phone.SendVoiceWithTemplateRequest.template:type_name -> spec.proto.extension.v1.phone.VoiceTemplate
	3, // 1: spec.proto.extension.v1.phone.VoiceTemplate.template_params:type_name -> spec.proto.extension.v1.phone.VoiceTemplate.TemplateParamsEntry
	0, // 2: spec.proto.extension.v1.phone.PhoneCallService.SendVoiceWithTemplate:input_type -> spec.proto.extension.v1.phone.SendVoiceWithTemplateRequest
	2, // 3: spec.proto.extension.v1.phone.PhoneCallService.SendVoiceWithTemplate:output_type -> spec.proto.extension.v1.phone.SendVoiceWithTemplateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spec_proto_extension_v1_phone_phone_proto_init() }
func file_spec_proto_extension_v1_phone_phone_proto_init() {
	if File_spec_proto_extension_v1_phone_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spec_proto_extension_v1_phone_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVoiceWithTemplateRequest); i {
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
		file_spec_proto_extension_v1_phone_phone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceTemplate); i {
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
		file_spec_proto_extension_v1_phone_phone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVoiceWithTemplateResponse); i {
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
			RawDescriptor: file_spec_proto_extension_v1_phone_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_proto_extension_v1_phone_phone_proto_goTypes,
		DependencyIndexes: file_spec_proto_extension_v1_phone_phone_proto_depIdxs,
		MessageInfos:      file_spec_proto_extension_v1_phone_phone_proto_msgTypes,
	}.Build()
	File_spec_proto_extension_v1_phone_phone_proto = out.File
	file_spec_proto_extension_v1_phone_phone_proto_rawDesc = nil
	file_spec_proto_extension_v1_phone_phone_proto_goTypes = nil
	file_spec_proto_extension_v1_phone_phone_proto_depIdxs = nil
}
