// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: proto/rdf.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Triple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject   *Subject `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate string   `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object    *Object  `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *Triple) Reset() {
	*x = Triple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rdf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Triple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Triple) ProtoMessage() {}

func (x *Triple) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rdf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Triple.ProtoReflect.Descriptor instead.
func (*Triple) Descriptor() ([]byte, []int) {
	return file_proto_rdf_proto_rawDescGZIP(), []int{0}
}

func (x *Triple) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *Triple) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

func (x *Triple) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Subject_Iri
	//	*Subject_BlankNode
	Value isSubject_Value `protobuf_oneof:"value"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rdf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rdf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_proto_rdf_proto_rawDescGZIP(), []int{1}
}

func (m *Subject) GetValue() isSubject_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Subject) GetIri() string {
	if x, ok := x.GetValue().(*Subject_Iri); ok {
		return x.Iri
	}
	return ""
}

func (x *Subject) GetBlankNode() string {
	if x, ok := x.GetValue().(*Subject_BlankNode); ok {
		return x.BlankNode
	}
	return ""
}

type isSubject_Value interface {
	isSubject_Value()
}

type Subject_Iri struct {
	Iri string `protobuf:"bytes,1,opt,name=iri,proto3,oneof"`
}

type Subject_BlankNode struct {
	BlankNode string `protobuf:"bytes,2,opt,name=blank_node,json=blankNode,proto3,oneof"`
}

func (*Subject_Iri) isSubject_Value() {}

func (*Subject_BlankNode) isSubject_Value() {}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Object_Iri
	//	*Object_BlankNode
	//	*Object_Literal
	Value isObject_Value `protobuf_oneof:"value"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rdf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rdf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_proto_rdf_proto_rawDescGZIP(), []int{2}
}

func (m *Object) GetValue() isObject_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Object) GetIri() string {
	if x, ok := x.GetValue().(*Object_Iri); ok {
		return x.Iri
	}
	return ""
}

func (x *Object) GetBlankNode() string {
	if x, ok := x.GetValue().(*Object_BlankNode); ok {
		return x.BlankNode
	}
	return ""
}

func (x *Object) GetLiteral() *Literal {
	if x, ok := x.GetValue().(*Object_Literal); ok {
		return x.Literal
	}
	return nil
}

type isObject_Value interface {
	isObject_Value()
}

type Object_Iri struct {
	Iri string `protobuf:"bytes,4,opt,name=iri,proto3,oneof"`
}

type Object_BlankNode struct {
	BlankNode string `protobuf:"bytes,5,opt,name=blank_node,json=blankNode,proto3,oneof"`
}

type Object_Literal struct {
	Literal *Literal `protobuf:"bytes,6,opt,name=literal,proto3,oneof"`
}

func (*Object_Iri) isObject_Value() {}

func (*Object_BlankNode) isObject_Value() {}

func (*Object_Literal) isObject_Value() {}

type Literal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Literal_String_
	//	*Literal_Int
	//	*Literal_Float64
	//	*Literal_Bool
	//	*Literal_Bytes
	Value isLiteral_Value `protobuf_oneof:"value"`
}

func (x *Literal) Reset() {
	*x = Literal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rdf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Literal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Literal) ProtoMessage() {}

func (x *Literal) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rdf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Literal.ProtoReflect.Descriptor instead.
func (*Literal) Descriptor() ([]byte, []int) {
	return file_proto_rdf_proto_rawDescGZIP(), []int{3}
}

func (m *Literal) GetValue() isLiteral_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Literal) GetString_() string {
	if x, ok := x.GetValue().(*Literal_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Literal) GetInt() int64 {
	if x, ok := x.GetValue().(*Literal_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Literal) GetFloat64() float64 {
	if x, ok := x.GetValue().(*Literal_Float64); ok {
		return x.Float64
	}
	return 0
}

func (x *Literal) GetBool() bool {
	if x, ok := x.GetValue().(*Literal_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Literal) GetBytes() []byte {
	if x, ok := x.GetValue().(*Literal_Bytes); ok {
		return x.Bytes
	}
	return nil
}

type isLiteral_Value interface {
	isLiteral_Value()
}

type Literal_String_ struct {
	String_ string `protobuf:"bytes,1,opt,name=string,proto3,oneof"`
}

type Literal_Int struct {
	Int int64 `protobuf:"varint,2,opt,name=int,proto3,oneof"`
}

type Literal_Float64 struct {
	Float64 float64 `protobuf:"fixed64,3,opt,name=float64,proto3,oneof"`
}

type Literal_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=bool,proto3,oneof"`
}

type Literal_Bytes struct {
	Bytes []byte `protobuf:"bytes,5,opt,name=bytes,proto3,oneof"`
}

func (*Literal_String_) isLiteral_Value() {}

func (*Literal_Int) isLiteral_Value() {}

func (*Literal_Float64) isLiteral_Value() {}

func (*Literal_Bool) isLiteral_Value() {}

func (*Literal_Bytes) isLiteral_Value() {}

var File_proto_rdf_proto protoreflect.FileDescriptor

var file_proto_rdf_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x64, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x64, 0x66, 0x22, 0x73, 0x0a, 0x06, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x64, 0x66, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72, 0x64, 0x66, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x47, 0x0a, 0x07, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x03, 0x69, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x69, 0x72, 0x69, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6c,
	0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x70, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x03, 0x69, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x69,
	0x72, 0x69, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6c, 0x61, 0x6e, 0x6b, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x64, 0x66, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x03,
	0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x03, 0x69, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x7a, 0x35, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x72, 0x64, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rdf_proto_rawDescOnce sync.Once
	file_proto_rdf_proto_rawDescData = file_proto_rdf_proto_rawDesc
)

func file_proto_rdf_proto_rawDescGZIP() []byte {
	file_proto_rdf_proto_rawDescOnce.Do(func() {
		file_proto_rdf_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rdf_proto_rawDescData)
	})
	return file_proto_rdf_proto_rawDescData
}

var file_proto_rdf_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_rdf_proto_goTypes = []interface{}{
	(*Triple)(nil),  // 0: rdf.Triple
	(*Subject)(nil), // 1: rdf.Subject
	(*Object)(nil),  // 2: rdf.Object
	(*Literal)(nil), // 3: rdf.Literal
}
var file_proto_rdf_proto_depIdxs = []int32{
	1, // 0: rdf.Triple.subject:type_name -> rdf.Subject
	2, // 1: rdf.Triple.object:type_name -> rdf.Object
	3, // 2: rdf.Object.literal:type_name -> rdf.Literal
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_rdf_proto_init() }
func file_proto_rdf_proto_init() {
	if File_proto_rdf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rdf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Triple); i {
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
		file_proto_rdf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_proto_rdf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_proto_rdf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Literal); i {
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
	file_proto_rdf_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Subject_Iri)(nil),
		(*Subject_BlankNode)(nil),
	}
	file_proto_rdf_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Object_Iri)(nil),
		(*Object_BlankNode)(nil),
		(*Object_Literal)(nil),
	}
	file_proto_rdf_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Literal_String_)(nil),
		(*Literal_Int)(nil),
		(*Literal_Float64)(nil),
		(*Literal_Bool)(nil),
		(*Literal_Bytes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rdf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_rdf_proto_goTypes,
		DependencyIndexes: file_proto_rdf_proto_depIdxs,
		MessageInfos:      file_proto_rdf_proto_msgTypes,
	}.Build()
	File_proto_rdf_proto = out.File
	file_proto_rdf_proto_rawDesc = nil
	file_proto_rdf_proto_goTypes = nil
	file_proto_rdf_proto_depIdxs = nil
}
