// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: pb/dataformat.proto

package geocube

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

// *
// Type of data supported by the Geocube & GDAL
type DataFormat_Dtype int32

const (
	DataFormat_UNDEFINED DataFormat_Dtype = 0
	DataFormat_UInt8     DataFormat_Dtype = 1
	DataFormat_UInt16    DataFormat_Dtype = 2
	DataFormat_UInt32    DataFormat_Dtype = 3
	DataFormat_Int8      DataFormat_Dtype = 4
	DataFormat_Int16     DataFormat_Dtype = 5
	DataFormat_Int32     DataFormat_Dtype = 6
	DataFormat_Float32   DataFormat_Dtype = 7
	DataFormat_Float64   DataFormat_Dtype = 8
	DataFormat_Complex64 DataFormat_Dtype = 9 // Pair of float32
)

// Enum value maps for DataFormat_Dtype.
var (
	DataFormat_Dtype_name = map[int32]string{
		0: "UNDEFINED",
		1: "UInt8",
		2: "UInt16",
		3: "UInt32",
		4: "Int8",
		5: "Int16",
		6: "Int32",
		7: "Float32",
		8: "Float64",
		9: "Complex64",
	}
	DataFormat_Dtype_value = map[string]int32{
		"UNDEFINED": 0,
		"UInt8":     1,
		"UInt16":    2,
		"UInt32":    3,
		"Int8":      4,
		"Int16":     5,
		"Int32":     6,
		"Float32":   7,
		"Float64":   8,
		"Complex64": 9,
	}
)

func (x DataFormat_Dtype) Enum() *DataFormat_Dtype {
	p := new(DataFormat_Dtype)
	*p = x
	return p
}

func (x DataFormat_Dtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataFormat_Dtype) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_dataformat_proto_enumTypes[0].Descriptor()
}

func (DataFormat_Dtype) Type() protoreflect.EnumType {
	return &file_pb_dataformat_proto_enumTypes[0]
}

func (x DataFormat_Dtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataFormat_Dtype.Descriptor instead.
func (DataFormat_Dtype) EnumDescriptor() ([]byte, []int) {
	return file_pb_dataformat_proto_rawDescGZIP(), []int{0, 0}
}

// *
// Format of the data of a dataset.
// Format is defined by the type of the data, its no-data value and the range of values (its interpretation depends on the use)
type DataFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtype    DataFormat_Dtype `protobuf:"varint,1,opt,name=dtype,proto3,enum=geocube.DataFormat_Dtype" json:"dtype,omitempty"` // Type of the data
	NoData   float64          `protobuf:"fixed64,2,opt,name=no_data,json=noData,proto3" json:"no_data,omitempty"`              // No-data value (supports any float values, including NaN)
	MinValue float64          `protobuf:"fixed64,3,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`        // Min value (usually used to map from one min value to another)
	MaxValue float64          `protobuf:"fixed64,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`        // Max value (usually used to map from one min value to another)
}

func (x *DataFormat) Reset() {
	*x = DataFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dataformat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFormat) ProtoMessage() {}

func (x *DataFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dataformat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFormat.ProtoReflect.Descriptor instead.
func (*DataFormat) Descriptor() ([]byte, []int) {
	return file_pb_dataformat_proto_rawDescGZIP(), []int{0}
}

func (x *DataFormat) GetDtype() DataFormat_Dtype {
	if x != nil {
		return x.Dtype
	}
	return DataFormat_UNDEFINED
}

func (x *DataFormat) GetNoData() float64 {
	if x != nil {
		return x.NoData
	}
	return 0
}

func (x *DataFormat) GetMinValue() float64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *DataFormat) GetMaxValue() float64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

var File_pb_dataformat_proto protoreflect.FileDescriptor

var file_pb_dataformat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x22, 0x95,
	0x02, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2f, 0x0a,
	0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67,
	0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x2e, 0x44, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x6e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x82, 0x01, 0x0a, 0x05, 0x44, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49,
	0x6e, 0x74, 0x38, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x6e, 0x74, 0x31, 0x36, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x49, 0x6e, 0x74, 0x38, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x31, 0x36,
	0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x10, 0x06, 0x12, 0x0b, 0x0a,
	0x07, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x36, 0x34, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x36, 0x34, 0x10, 0x09, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x67,
	0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_dataformat_proto_rawDescOnce sync.Once
	file_pb_dataformat_proto_rawDescData = file_pb_dataformat_proto_rawDesc
)

func file_pb_dataformat_proto_rawDescGZIP() []byte {
	file_pb_dataformat_proto_rawDescOnce.Do(func() {
		file_pb_dataformat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_dataformat_proto_rawDescData)
	})
	return file_pb_dataformat_proto_rawDescData
}

var file_pb_dataformat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_dataformat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_dataformat_proto_goTypes = []interface{}{
	(DataFormat_Dtype)(0), // 0: geocube.DataFormat.Dtype
	(*DataFormat)(nil),    // 1: geocube.DataFormat
}
var file_pb_dataformat_proto_depIdxs = []int32{
	0, // 0: geocube.DataFormat.dtype:type_name -> geocube.DataFormat.Dtype
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_dataformat_proto_init() }
func file_pb_dataformat_proto_init() {
	if File_pb_dataformat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_dataformat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFormat); i {
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
			RawDescriptor: file_pb_dataformat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_dataformat_proto_goTypes,
		DependencyIndexes: file_pb_dataformat_proto_depIdxs,
		EnumInfos:         file_pb_dataformat_proto_enumTypes,
		MessageInfos:      file_pb_dataformat_proto_msgTypes,
	}.Build()
	File_pb_dataformat_proto = out.File
	file_pb_dataformat_proto_rawDesc = nil
	file_pb_dataformat_proto_goTypes = nil
	file_pb_dataformat_proto_depIdxs = nil
}
