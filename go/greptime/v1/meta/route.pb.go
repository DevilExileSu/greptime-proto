// Copyright 2023 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/route.proto

package meta

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TableName  *TableName     `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Partitions []*Partition   `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	TableInfo  []byte         `protobuf:"bytes,4,opt,name=table_info,json=tableInfo,proto3" json:"table_info,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CreateRequest) GetTableName() *TableName {
	if x != nil {
		return x.TableName
	}
	return nil
}

func (x *CreateRequest) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

func (x *CreateRequest) GetTableInfo() []byte {
	if x != nil {
		return x.TableInfo
	}
	return nil
}

type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TableNames []*TableName   `protobuf:"bytes,2,rep,name=table_names,json=tableNames,proto3" json:"table_names,omitempty"`
	TableIds   []*TableId     `protobuf:"bytes,3,rep,name=table_ids,json=tableIds,proto3" json:"table_ids,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{1}
}

func (x *RouteRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RouteRequest) GetTableNames() []*TableName {
	if x != nil {
		return x.TableNames
	}
	return nil
}

func (x *RouteRequest) GetTableIds() []*TableId {
	if x != nil {
		return x.TableIds
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TableName *TableName     `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	TableId   *TableId       `protobuf:"bytes,3,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DeleteRequest) GetTableName() *TableName {
	if x != nil {
		return x.TableName
	}
	return nil
}

func (x *DeleteRequest) GetTableId() *TableId {
	if x != nil {
		return x.TableId
	}
	return nil
}

type RouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header      *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Peers       []*Peer         `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	TableRoutes []*TableRoute   `protobuf:"bytes,3,rep,name=table_routes,json=tableRoutes,proto3" json:"table_routes,omitempty"`
}

func (x *RouteResponse) Reset() {
	*x = RouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteResponse) ProtoMessage() {}

func (x *RouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteResponse.ProtoReflect.Descriptor instead.
func (*RouteResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{3}
}

func (x *RouteResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RouteResponse) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *RouteResponse) GetTableRoutes() []*TableRoute {
	if x != nil {
		return x.TableRoutes
	}
	return nil
}

type TableRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table        *Table         `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	RegionRoutes []*RegionRoute `protobuf:"bytes,2,rep,name=region_routes,json=regionRoutes,proto3" json:"region_routes,omitempty"`
}

func (x *TableRoute) Reset() {
	*x = TableRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableRoute) ProtoMessage() {}

func (x *TableRoute) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableRoute.ProtoReflect.Descriptor instead.
func (*TableRoute) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{4}
}

func (x *TableRoute) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *TableRoute) GetRegionRoutes() []*RegionRoute {
	if x != nil {
		return x.RegionRoutes
	}
	return nil
}

type RegionRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region *Region `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// single leader node for write task
	LeaderPeerIndex uint64 `protobuf:"varint,2,opt,name=leader_peer_index,json=leaderPeerIndex,proto3" json:"leader_peer_index,omitempty"`
	// multiple follower nodes for read task
	FollowerPeerIndexes []uint64 `protobuf:"varint,3,rep,packed,name=follower_peer_indexes,json=followerPeerIndexes,proto3" json:"follower_peer_indexes,omitempty"`
}

func (x *RegionRoute) Reset() {
	*x = RegionRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionRoute) ProtoMessage() {}

func (x *RegionRoute) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionRoute.ProtoReflect.Descriptor instead.
func (*RegionRoute) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{5}
}

func (x *RegionRoute) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *RegionRoute) GetLeaderPeerIndex() uint64 {
	if x != nil {
		return x.LeaderPeerIndex
	}
	return 0
}

func (x *RegionRoute) GetFollowerPeerIndexes() []uint64 {
	if x != nil {
		return x.FollowerPeerIndexes
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TableName   *TableName `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	TableSchema []byte     `protobuf:"bytes,3,opt,name=table_schema,json=tableSchema,proto3" json:"table_schema,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{6}
}

func (x *Table) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Table) GetTableName() *TableName {
	if x != nil {
		return x.TableName
	}
	return nil
}

func (x *Table) GetTableSchema() []byte {
	if x != nil {
		return x.TableSchema
	}
	return nil
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO(LFC): Maybe use message RegionNumber?
	Id        uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Partition *Partition        `protobuf:"bytes,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Attrs     map[string]string `protobuf:"bytes,100,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{7}
}

func (x *Region) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Region) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Region) GetPartition() *Partition {
	if x != nil {
		return x.Partition
	}
	return nil
}

func (x *Region) GetAttrs() map[string]string {
	if x != nil {
		return x.Attrs
	}
	return nil
}

// PARTITION `region_name` VALUES LESS THAN (value_list)
type Partition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnList [][]byte `protobuf:"bytes,1,rep,name=column_list,json=columnList,proto3" json:"column_list,omitempty"`
	ValueList  [][]byte `protobuf:"bytes,2,rep,name=value_list,json=valueList,proto3" json:"value_list,omitempty"`
}

func (x *Partition) Reset() {
	*x = Partition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partition) ProtoMessage() {}

func (x *Partition) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partition.ProtoReflect.Descriptor instead.
func (*Partition) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{8}
}

func (x *Partition) GetColumnList() [][]byte {
	if x != nil {
		return x.ColumnList
	}
	return nil
}

func (x *Partition) GetValueList() [][]byte {
	if x != nil {
		return x.ValueList
	}
	return nil
}

// This message is only for saving into store.
type TableRouteValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers      []*Peer     `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	TableRoute *TableRoute `protobuf:"bytes,2,opt,name=table_route,json=tableRoute,proto3" json:"table_route,omitempty"`
}

func (x *TableRouteValue) Reset() {
	*x = TableRouteValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableRouteValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableRouteValue) ProtoMessage() {}

func (x *TableRouteValue) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableRouteValue.ProtoReflect.Descriptor instead.
func (*TableRouteValue) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_route_proto_rawDescGZIP(), []int{9}
}

func (x *TableRouteValue) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

func (x *TableRouteValue) GetTableRoute() *TableRoute {
	if x != nil {
		return x.TableRoute
	}
	return nil
}

var File_greptime_v1_meta_route_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_route_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x1a, 0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0a,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x09, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x64, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0xb8, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0b, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x7f, 0x0a, 0x0a, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x11, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x15, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x13, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0x76, 0x0a,
	0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0xdc, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x74,
	0x74, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x7e, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0xf0, 0x01, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x05, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_route_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_route_proto_rawDescData = file_greptime_v1_meta_route_proto_rawDesc
)

func file_greptime_v1_meta_route_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_route_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_route_proto_rawDescData)
	})
	return file_greptime_v1_meta_route_proto_rawDescData
}

var file_greptime_v1_meta_route_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_greptime_v1_meta_route_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: greptime.v1.meta.CreateRequest
	(*RouteRequest)(nil),    // 1: greptime.v1.meta.RouteRequest
	(*DeleteRequest)(nil),   // 2: greptime.v1.meta.DeleteRequest
	(*RouteResponse)(nil),   // 3: greptime.v1.meta.RouteResponse
	(*TableRoute)(nil),      // 4: greptime.v1.meta.TableRoute
	(*RegionRoute)(nil),     // 5: greptime.v1.meta.RegionRoute
	(*Table)(nil),           // 6: greptime.v1.meta.Table
	(*Region)(nil),          // 7: greptime.v1.meta.Region
	(*Partition)(nil),       // 8: greptime.v1.meta.Partition
	(*TableRouteValue)(nil), // 9: greptime.v1.meta.TableRouteValue
	nil,                     // 10: greptime.v1.meta.Region.AttrsEntry
	(*RequestHeader)(nil),   // 11: greptime.v1.meta.RequestHeader
	(*TableName)(nil),       // 12: greptime.v1.meta.TableName
	(*TableId)(nil),         // 13: greptime.v1.meta.TableId
	(*ResponseHeader)(nil),  // 14: greptime.v1.meta.ResponseHeader
	(*Peer)(nil),            // 15: greptime.v1.meta.Peer
}
var file_greptime_v1_meta_route_proto_depIdxs = []int32{
	11, // 0: greptime.v1.meta.CreateRequest.header:type_name -> greptime.v1.meta.RequestHeader
	12, // 1: greptime.v1.meta.CreateRequest.table_name:type_name -> greptime.v1.meta.TableName
	8,  // 2: greptime.v1.meta.CreateRequest.partitions:type_name -> greptime.v1.meta.Partition
	11, // 3: greptime.v1.meta.RouteRequest.header:type_name -> greptime.v1.meta.RequestHeader
	12, // 4: greptime.v1.meta.RouteRequest.table_names:type_name -> greptime.v1.meta.TableName
	13, // 5: greptime.v1.meta.RouteRequest.table_ids:type_name -> greptime.v1.meta.TableId
	11, // 6: greptime.v1.meta.DeleteRequest.header:type_name -> greptime.v1.meta.RequestHeader
	12, // 7: greptime.v1.meta.DeleteRequest.table_name:type_name -> greptime.v1.meta.TableName
	13, // 8: greptime.v1.meta.DeleteRequest.table_id:type_name -> greptime.v1.meta.TableId
	14, // 9: greptime.v1.meta.RouteResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	15, // 10: greptime.v1.meta.RouteResponse.peers:type_name -> greptime.v1.meta.Peer
	4,  // 11: greptime.v1.meta.RouteResponse.table_routes:type_name -> greptime.v1.meta.TableRoute
	6,  // 12: greptime.v1.meta.TableRoute.table:type_name -> greptime.v1.meta.Table
	5,  // 13: greptime.v1.meta.TableRoute.region_routes:type_name -> greptime.v1.meta.RegionRoute
	7,  // 14: greptime.v1.meta.RegionRoute.region:type_name -> greptime.v1.meta.Region
	12, // 15: greptime.v1.meta.Table.table_name:type_name -> greptime.v1.meta.TableName
	8,  // 16: greptime.v1.meta.Region.partition:type_name -> greptime.v1.meta.Partition
	10, // 17: greptime.v1.meta.Region.attrs:type_name -> greptime.v1.meta.Region.AttrsEntry
	15, // 18: greptime.v1.meta.TableRouteValue.peers:type_name -> greptime.v1.meta.Peer
	4,  // 19: greptime.v1.meta.TableRouteValue.table_route:type_name -> greptime.v1.meta.TableRoute
	0,  // 20: greptime.v1.meta.Router.Create:input_type -> greptime.v1.meta.CreateRequest
	1,  // 21: greptime.v1.meta.Router.Route:input_type -> greptime.v1.meta.RouteRequest
	2,  // 22: greptime.v1.meta.Router.Delete:input_type -> greptime.v1.meta.DeleteRequest
	3,  // 23: greptime.v1.meta.Router.Create:output_type -> greptime.v1.meta.RouteResponse
	3,  // 24: greptime.v1.meta.Router.Route:output_type -> greptime.v1.meta.RouteResponse
	3,  // 25: greptime.v1.meta.Router.Delete:output_type -> greptime.v1.meta.RouteResponse
	23, // [23:26] is the sub-list for method output_type
	20, // [20:23] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_route_proto_init() }
func file_greptime_v1_meta_route_proto_init() {
	if File_greptime_v1_meta_route_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRequest); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteResponse); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableRoute); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionRoute); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partition); i {
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
		file_greptime_v1_meta_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableRouteValue); i {
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
			RawDescriptor: file_greptime_v1_meta_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_route_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_route_proto_depIdxs,
		MessageInfos:      file_greptime_v1_meta_route_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_route_proto = out.File
	file_greptime_v1_meta_route_proto_rawDesc = nil
	file_greptime_v1_meta_route_proto_goTypes = nil
	file_greptime_v1_meta_route_proto_depIdxs = nil
}
