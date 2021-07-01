// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: pb/catalog.proto

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

type ByteOrder int32

const (
	ByteOrder_LittleEndian ByteOrder = 0
	ByteOrder_BigEndian    ByteOrder = 1
)

// Enum value maps for ByteOrder.
var (
	ByteOrder_name = map[int32]string{
		0: "LittleEndian",
		1: "BigEndian",
	}
	ByteOrder_value = map[string]int32{
		"LittleEndian": 0,
		"BigEndian":    1,
	}
)

func (x ByteOrder) Enum() *ByteOrder {
	p := new(ByteOrder)
	*p = x
	return p
}

func (x ByteOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ByteOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_catalog_proto_enumTypes[0].Descriptor()
}

func (ByteOrder) Type() protoreflect.EnumType {
	return &file_pb_catalog_proto_enumTypes[0]
}

func (x ByteOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ByteOrder.Descriptor instead.
func (ByteOrder) EnumDescriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{0}
}

type FileFormat int32

const (
	FileFormat_Raw   FileFormat = 0 // raw bitmap
	FileFormat_GTiff FileFormat = 1
)

// Enum value maps for FileFormat.
var (
	FileFormat_name = map[int32]string{
		0: "Raw",
		1: "GTiff",
	}
	FileFormat_value = map[string]int32{
		"Raw":   0,
		"GTiff": 1,
	}
)

func (x FileFormat) Enum() *FileFormat {
	p := new(FileFormat)
	*p = x
	return p
}

func (x FileFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_catalog_proto_enumTypes[1].Descriptor()
}

func (FileFormat) Type() protoreflect.EnumType {
	return &file_pb_catalog_proto_enumTypes[1]
}

func (x FileFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileFormat.Descriptor instead.
func (FileFormat) EnumDescriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{1}
}

type Shape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dim1 int32 `protobuf:"varint,1,opt,name=dim1,proto3" json:"dim1,omitempty"`
	Dim2 int32 `protobuf:"varint,2,opt,name=dim2,proto3" json:"dim2,omitempty"`
	Dim3 int32 `protobuf:"varint,3,opt,name=dim3,proto3" json:"dim3,omitempty"`
}

func (x *Shape) Reset() {
	*x = Shape{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shape) ProtoMessage() {}

func (x *Shape) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shape.ProtoReflect.Descriptor instead.
func (*Shape) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Shape) GetDim1() int32 {
	if x != nil {
		return x.Dim1
	}
	return 0
}

func (x *Shape) GetDim2() int32 {
	if x != nil {
		return x.Dim2
	}
	return 0
}

func (x *Shape) GetDim3() int32 {
	if x != nil {
		return x.Dim3
	}
	return 0
}

type ImageHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shape       *Shape           `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype       DataFormat_Dtype `protobuf:"varint,2,opt,name=dtype,proto3,enum=geocube.DataFormat_Dtype" json:"dtype,omitempty"`
	Data        []byte           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	NbParts     int32            `protobuf:"varint,4,opt,name=nb_parts,json=nbParts,proto3" json:"nb_parts,omitempty"`
	Size        int64            `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Order       ByteOrder        `protobuf:"varint,6,opt,name=order,proto3,enum=geocube.ByteOrder" json:"order,omitempty"`
	Compression bool             `protobuf:"varint,7,opt,name=compression,proto3" json:"compression,omitempty"` // Deflate compressed data format, described in RFC 1951
	Records     []*Record        `protobuf:"bytes,8,rep,name=records,proto3" json:"records,omitempty"`
	Error       string           `protobuf:"bytes,9,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ImageHeader) Reset() {
	*x = ImageHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageHeader) ProtoMessage() {}

func (x *ImageHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageHeader.ProtoReflect.Descriptor instead.
func (*ImageHeader) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *ImageHeader) GetShape() *Shape {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *ImageHeader) GetDtype() DataFormat_Dtype {
	if x != nil {
		return x.Dtype
	}
	return DataFormat_UNDEFINED
}

func (x *ImageHeader) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImageHeader) GetNbParts() int32 {
	if x != nil {
		return x.NbParts
	}
	return 0
}

func (x *ImageHeader) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ImageHeader) GetOrder() ByteOrder {
	if x != nil {
		return x.Order
	}
	return ByteOrder_LittleEndian
}

func (x *ImageHeader) GetCompression() bool {
	if x != nil {
		return x.Compression
	}
	return false
}

func (x *ImageHeader) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *ImageHeader) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ImageChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Part int32  `protobuf:"varint,1,opt,name=part,proto3" json:"part,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ImageChunk) Reset() {
	*x = ImageChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageChunk) ProtoMessage() {}

func (x *ImageChunk) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageChunk.ProtoReflect.Descriptor instead.
func (*ImageChunk) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *ImageChunk) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *ImageChunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ImageFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ImageFile) Reset() {
	*x = ImageFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageFile) ProtoMessage() {}

func (x *ImageFile) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageFile.ProtoReflect.Descriptor instead.
func (*ImageFile) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *ImageFile) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCubeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RecordsLister:
	//	*GetCubeRequest_Records
	//	*GetCubeRequest_Filters
	//	*GetCubeRequest_Grecords
	RecordsLister    isGetCubeRequest_RecordsLister `protobuf_oneof:"records_lister"`
	InstancesId      []string                       `protobuf:"bytes,3,rep,name=instances_id,json=instancesId,proto3" json:"instances_id,omitempty"` // At least one, and all must be instance of the same variable. Only one is actually supported
	Crs              string                         `protobuf:"bytes,4,opt,name=crs,proto3" json:"crs,omitempty"`
	PixToCrs         *GeoTransform                  `protobuf:"bytes,5,opt,name=pix_to_crs,json=pixToCrs,proto3" json:"pix_to_crs,omitempty"`
	Size             *Size                          `protobuf:"bytes,6,opt,name=size,proto3" json:"size,omitempty"`
	CompressionLevel int32                          `protobuf:"varint,7,opt,name=compression_level,json=compressionLevel,proto3" json:"compression_level,omitempty"` // -1 to 10 (-1:default, 0: no compression, 1->9: level of compression from the fastest to the best compression)
	HeadersOnly      bool                           `protobuf:"varint,8,opt,name=headers_only,json=headersOnly,proto3" json:"headers_only,omitempty"`                // Only returns headers
	Format           FileFormat                     `protobuf:"varint,9,opt,name=format,proto3,enum=geocube.FileFormat" json:"format,omitempty"`
}

func (x *GetCubeRequest) Reset() {
	*x = GetCubeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCubeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCubeRequest) ProtoMessage() {}

func (x *GetCubeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCubeRequest.ProtoReflect.Descriptor instead.
func (*GetCubeRequest) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{4}
}

func (m *GetCubeRequest) GetRecordsLister() isGetCubeRequest_RecordsLister {
	if m != nil {
		return m.RecordsLister
	}
	return nil
}

func (x *GetCubeRequest) GetRecords() *RecordList {
	if x, ok := x.GetRecordsLister().(*GetCubeRequest_Records); ok {
		return x.Records
	}
	return nil
}

func (x *GetCubeRequest) GetFilters() *RecordFilters {
	if x, ok := x.GetRecordsLister().(*GetCubeRequest_Filters); ok {
		return x.Filters
	}
	return nil
}

func (x *GetCubeRequest) GetGrecords() *GroupedRecordsList {
	if x, ok := x.GetRecordsLister().(*GetCubeRequest_Grecords); ok {
		return x.Grecords
	}
	return nil
}

func (x *GetCubeRequest) GetInstancesId() []string {
	if x != nil {
		return x.InstancesId
	}
	return nil
}

func (x *GetCubeRequest) GetCrs() string {
	if x != nil {
		return x.Crs
	}
	return ""
}

func (x *GetCubeRequest) GetPixToCrs() *GeoTransform {
	if x != nil {
		return x.PixToCrs
	}
	return nil
}

func (x *GetCubeRequest) GetSize() *Size {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *GetCubeRequest) GetCompressionLevel() int32 {
	if x != nil {
		return x.CompressionLevel
	}
	return 0
}

func (x *GetCubeRequest) GetHeadersOnly() bool {
	if x != nil {
		return x.HeadersOnly
	}
	return false
}

func (x *GetCubeRequest) GetFormat() FileFormat {
	if x != nil {
		return x.Format
	}
	return FileFormat_Raw
}

type isGetCubeRequest_RecordsLister interface {
	isGetCubeRequest_RecordsLister()
}

type GetCubeRequest_Records struct {
	Records *RecordList `protobuf:"bytes,1,opt,name=records,proto3,oneof"` // At least one
}

type GetCubeRequest_Filters struct {
	Filters *RecordFilters `protobuf:"bytes,2,opt,name=filters,proto3,oneof"`
}

type GetCubeRequest_Grecords struct {
	Grecords *GroupedRecordsList `protobuf:"bytes,10,opt,name=grecords,proto3,oneof"`
}

func (*GetCubeRequest_Records) isGetCubeRequest_RecordsLister() {}

func (*GetCubeRequest_Filters) isGetCubeRequest_RecordsLister() {}

func (*GetCubeRequest_Grecords) isGetCubeRequest_RecordsLister() {}

type GetCubeResponseHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	NbDatasets int64 `protobuf:"varint,2,opt,name=nb_datasets,json=nbDatasets,proto3" json:"nb_datasets,omitempty"`
}

func (x *GetCubeResponseHeader) Reset() {
	*x = GetCubeResponseHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCubeResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCubeResponseHeader) ProtoMessage() {}

func (x *GetCubeResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCubeResponseHeader.ProtoReflect.Descriptor instead.
func (*GetCubeResponseHeader) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *GetCubeResponseHeader) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetCubeResponseHeader) GetNbDatasets() int64 {
	if x != nil {
		return x.NbDatasets
	}
	return 0
}

type GetCubeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*GetCubeResponse_GlobalHeader
	//	*GetCubeResponse_Header
	//	*GetCubeResponse_Chunk
	Response isGetCubeResponse_Response `protobuf_oneof:"response"`
}

func (x *GetCubeResponse) Reset() {
	*x = GetCubeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCubeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCubeResponse) ProtoMessage() {}

func (x *GetCubeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCubeResponse.ProtoReflect.Descriptor instead.
func (*GetCubeResponse) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{6}
}

func (m *GetCubeResponse) GetResponse() isGetCubeResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *GetCubeResponse) GetGlobalHeader() *GetCubeResponseHeader {
	if x, ok := x.GetResponse().(*GetCubeResponse_GlobalHeader); ok {
		return x.GlobalHeader
	}
	return nil
}

func (x *GetCubeResponse) GetHeader() *ImageHeader {
	if x, ok := x.GetResponse().(*GetCubeResponse_Header); ok {
		return x.Header
	}
	return nil
}

func (x *GetCubeResponse) GetChunk() *ImageChunk {
	if x, ok := x.GetResponse().(*GetCubeResponse_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isGetCubeResponse_Response interface {
	isGetCubeResponse_Response()
}

type GetCubeResponse_GlobalHeader struct {
	GlobalHeader *GetCubeResponseHeader `protobuf:"bytes,3,opt,name=global_header,json=globalHeader,proto3,oneof"`
}

type GetCubeResponse_Header struct {
	Header *ImageHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type GetCubeResponse_Chunk struct {
	Chunk *ImageChunk `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*GetCubeResponse_GlobalHeader) isGetCubeResponse_Response() {}

func (*GetCubeResponse_Header) isGetCubeResponse_Response() {}

func (*GetCubeResponse_Chunk) isGetCubeResponse_Response() {}

type GetTileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	X          int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y          int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Z          int32  `protobuf:"varint,4,opt,name=z,proto3" json:"z,omitempty"`
	// Types that are assignable to RecordsLister:
	//	*GetTileRequest_Records
	RecordsLister isGetTileRequest_RecordsLister `protobuf_oneof:"records_lister"`
}

func (x *GetTileRequest) Reset() {
	*x = GetTileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTileRequest) ProtoMessage() {}

func (x *GetTileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTileRequest.ProtoReflect.Descriptor instead.
func (*GetTileRequest) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{7}
}

func (x *GetTileRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *GetTileRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *GetTileRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *GetTileRequest) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (m *GetTileRequest) GetRecordsLister() isGetTileRequest_RecordsLister {
	if m != nil {
		return m.RecordsLister
	}
	return nil
}

func (x *GetTileRequest) GetRecords() *RecordList {
	if x, ok := x.GetRecordsLister().(*GetTileRequest_Records); ok {
		return x.Records
	}
	return nil
}

type isGetTileRequest_RecordsLister interface {
	isGetTileRequest_RecordsLister()
}

type GetTileRequest_Records struct {
	Records *RecordList `protobuf:"bytes,5,opt,name=records,proto3,oneof"` // At least one
}

func (*GetTileRequest_Records) isGetTileRequest_RecordsLister() {}

type GetTileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *ImageFile `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetTileResponse) Reset() {
	*x = GetTileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_catalog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTileResponse) ProtoMessage() {}

func (x *GetTileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_catalog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTileResponse.ProtoReflect.Descriptor instead.
func (*GetTileResponse) Descriptor() ([]byte, []int) {
	return file_pb_catalog_proto_rawDescGZIP(), []int{8}
}

func (x *GetTileResponse) GetImage() *ImageFile {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_pb_catalog_proto protoreflect.FileDescriptor

var file_pb_catalog_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x1a, 0x13, 0x70, 0x62, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x70, 0x62, 0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x69, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x69, 0x6d,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x64, 0x69, 0x6d, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x69, 0x6d, 0x33, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x69, 0x6d, 0x33, 0x22, 0xb4, 0x02, 0x0a, 0x0b, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x53, 0x68, 0x61, 0x70, 0x65, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12,
	0x2f, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x2e, 0x44, 0x74, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x62, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x62, 0x50, 0x61, 0x72, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x29, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x34, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcc, 0x03, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65,
	0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x39, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x08, 0x67, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x72, 0x73, 0x12,
	0x33, 0x0a, 0x0a, 0x70, 0x69, 0x78, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x47, 0x65,
	0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x69, 0x78, 0x54,
	0x6f, 0x43, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x62,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x62, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x0a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6f, 0x63,
	0x75, 0x62, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x54, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2a, 0x2c, 0x0a, 0x09, 0x42, 0x79, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x45,
	0x6e, 0x64, 0x69, 0x61, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x69, 0x67, 0x45, 0x6e,
	0x64, 0x69, 0x61, 0x6e, 0x10, 0x01, 0x2a, 0x20, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x47, 0x54, 0x69, 0x66, 0x66, 0x10, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x62,
	0x3b, 0x67, 0x65, 0x6f, 0x63, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_catalog_proto_rawDescOnce sync.Once
	file_pb_catalog_proto_rawDescData = file_pb_catalog_proto_rawDesc
)

func file_pb_catalog_proto_rawDescGZIP() []byte {
	file_pb_catalog_proto_rawDescOnce.Do(func() {
		file_pb_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_catalog_proto_rawDescData)
	})
	return file_pb_catalog_proto_rawDescData
}

var file_pb_catalog_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_catalog_proto_goTypes = []interface{}{
	(ByteOrder)(0),                // 0: geocube.ByteOrder
	(FileFormat)(0),               // 1: geocube.FileFormat
	(*Shape)(nil),                 // 2: geocube.Shape
	(*ImageHeader)(nil),           // 3: geocube.ImageHeader
	(*ImageChunk)(nil),            // 4: geocube.ImageChunk
	(*ImageFile)(nil),             // 5: geocube.ImageFile
	(*GetCubeRequest)(nil),        // 6: geocube.GetCubeRequest
	(*GetCubeResponseHeader)(nil), // 7: geocube.GetCubeResponseHeader
	(*GetCubeResponse)(nil),       // 8: geocube.GetCubeResponse
	(*GetTileRequest)(nil),        // 9: geocube.GetTileRequest
	(*GetTileResponse)(nil),       // 10: geocube.GetTileResponse
	(DataFormat_Dtype)(0),         // 11: geocube.DataFormat.Dtype
	(*Record)(nil),                // 12: geocube.Record
	(*RecordList)(nil),            // 13: geocube.RecordList
	(*RecordFilters)(nil),         // 14: geocube.RecordFilters
	(*GroupedRecordsList)(nil),    // 15: geocube.GroupedRecordsList
	(*GeoTransform)(nil),          // 16: geocube.GeoTransform
	(*Size)(nil),                  // 17: geocube.Size
}
var file_pb_catalog_proto_depIdxs = []int32{
	2,  // 0: geocube.ImageHeader.shape:type_name -> geocube.Shape
	11, // 1: geocube.ImageHeader.dtype:type_name -> geocube.DataFormat.Dtype
	0,  // 2: geocube.ImageHeader.order:type_name -> geocube.ByteOrder
	12, // 3: geocube.ImageHeader.records:type_name -> geocube.Record
	13, // 4: geocube.GetCubeRequest.records:type_name -> geocube.RecordList
	14, // 5: geocube.GetCubeRequest.filters:type_name -> geocube.RecordFilters
	15, // 6: geocube.GetCubeRequest.grecords:type_name -> geocube.GroupedRecordsList
	16, // 7: geocube.GetCubeRequest.pix_to_crs:type_name -> geocube.GeoTransform
	17, // 8: geocube.GetCubeRequest.size:type_name -> geocube.Size
	1,  // 9: geocube.GetCubeRequest.format:type_name -> geocube.FileFormat
	7,  // 10: geocube.GetCubeResponse.global_header:type_name -> geocube.GetCubeResponseHeader
	3,  // 11: geocube.GetCubeResponse.header:type_name -> geocube.ImageHeader
	4,  // 12: geocube.GetCubeResponse.chunk:type_name -> geocube.ImageChunk
	13, // 13: geocube.GetTileRequest.records:type_name -> geocube.RecordList
	5,  // 14: geocube.GetTileResponse.image:type_name -> geocube.ImageFile
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_pb_catalog_proto_init() }
func file_pb_catalog_proto_init() {
	if File_pb_catalog_proto != nil {
		return
	}
	file_pb_dataformat_proto_init()
	file_pb_records_proto_init()
	file_pb_layouts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shape); i {
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
		file_pb_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageHeader); i {
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
		file_pb_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageChunk); i {
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
		file_pb_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageFile); i {
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
		file_pb_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCubeRequest); i {
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
		file_pb_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCubeResponseHeader); i {
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
		file_pb_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCubeResponse); i {
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
		file_pb_catalog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTileRequest); i {
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
		file_pb_catalog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTileResponse); i {
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
	file_pb_catalog_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetCubeRequest_Records)(nil),
		(*GetCubeRequest_Filters)(nil),
		(*GetCubeRequest_Grecords)(nil),
	}
	file_pb_catalog_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*GetCubeResponse_GlobalHeader)(nil),
		(*GetCubeResponse_Header)(nil),
		(*GetCubeResponse_Chunk)(nil),
	}
	file_pb_catalog_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*GetTileRequest_Records)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_catalog_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_catalog_proto_goTypes,
		DependencyIndexes: file_pb_catalog_proto_depIdxs,
		EnumInfos:         file_pb_catalog_proto_enumTypes,
		MessageInfos:      file_pb_catalog_proto_msgTypes,
	}.Build()
	File_pb_catalog_proto = out.File
	file_pb_catalog_proto_rawDesc = nil
	file_pb_catalog_proto_goTypes = nil
	file_pb_catalog_proto_depIdxs = nil
}
