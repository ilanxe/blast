// Code generated by protoc-gen-go.
// source: data.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	data.proto

It has these top-level messages:
	Metadata
	JoinRequest
	JoinResponse
	LeaveRequest
	LeaveResponse
	Peer
	PeersResponse
	SnapshotResponse
	GetRequest
	GetResponse
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	Document
	UpdateRequest
	BulkRequest
	BulkResponse
	SearchRequest
	SearchResponse
	Proposal
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateRequest_Type int32

const (
	UpdateRequest_NOOP   UpdateRequest_Type = 0
	UpdateRequest_PUT    UpdateRequest_Type = 1
	UpdateRequest_DELETE UpdateRequest_Type = 2
)

var UpdateRequest_Type_name = map[int32]string{
	0: "NOOP",
	1: "PUT",
	2: "DELETE",
}
var UpdateRequest_Type_value = map[string]int32{
	"NOOP":   0,
	"PUT":    1,
	"DELETE": 2,
}

func (x UpdateRequest_Type) String() string {
	return proto.EnumName(UpdateRequest_Type_name, int32(x))
}
func (UpdateRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{15, 0} }

type Proposal_Type int32

const (
	Proposal_NOOP   Proposal_Type = 0
	Proposal_PUT    Proposal_Type = 1
	Proposal_DELETE Proposal_Type = 2
	Proposal_BULK   Proposal_Type = 3
	Proposal_JOIN   Proposal_Type = 4
	Proposal_LEAVE  Proposal_Type = 5
)

var Proposal_Type_name = map[int32]string{
	0: "NOOP",
	1: "PUT",
	2: "DELETE",
	3: "BULK",
	4: "JOIN",
	5: "LEAVE",
}
var Proposal_Type_value = map[string]int32{
	"NOOP":   0,
	"PUT":    1,
	"DELETE": 2,
	"BULK":   3,
	"JOIN":   4,
	"LEAVE":  5,
}

func (x Proposal_Type) String() string {
	return proto.EnumName(Proposal_Type_name, int32(x))
}
func (Proposal_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{20, 0} }

type Metadata struct {
	GrpcAddress string `protobuf:"bytes,2,opt,name=grpc_address,json=grpcAddress" json:"grpc_address,omitempty"`
	HttpAddress string `protobuf:"bytes,3,opt,name=http_address,json=httpAddress" json:"http_address,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Metadata) GetGrpcAddress() string {
	if m != nil {
		return m.GrpcAddress
	}
	return ""
}

func (m *Metadata) GetHttpAddress() string {
	if m != nil {
		return m.HttpAddress
	}
	return ""
}

type JoinRequest struct {
	NodeId   string    `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Address  string    `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Metadata *Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JoinRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *JoinRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *JoinRequest) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type JoinResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *JoinResponse) Reset()                    { *m = JoinResponse{} }
func (m *JoinResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()               {}
func (*JoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *JoinResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *JoinResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LeaveRequest struct {
	NodeId  string `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *LeaveRequest) Reset()                    { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string            { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()               {}
func (*LeaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LeaveRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *LeaveRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type LeaveResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *LeaveResponse) Reset()                    { *m = LeaveResponse{} }
func (m *LeaveResponse) String() string            { return proto.CompactTextString(m) }
func (*LeaveResponse) ProtoMessage()               {}
func (*LeaveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LeaveResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LeaveResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Peer struct {
	NodeId   string    `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Address  string    `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Leader   bool      `protobuf:"varint,3,opt,name=leader" json:"leader,omitempty"`
	Metadata *Metadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Peer) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Peer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Peer) GetLeader() bool {
	if m != nil {
		return m.Leader
	}
	return false
}

func (m *Peer) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PeersResponse struct {
	Peers   []*Peer `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
	Success bool    `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *PeersResponse) Reset()                    { *m = PeersResponse{} }
func (m *PeersResponse) String() string            { return proto.CompactTextString(m) }
func (*PeersResponse) ProtoMessage()               {}
func (*PeersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PeersResponse) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *PeersResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PeersResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SnapshotResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SnapshotResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SnapshotResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetResponse struct {
	Id      string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Fields  *google_protobuf.Any `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
	Success bool                 `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	Message string               `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetResponse) GetFields() *google_protobuf.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *GetResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PutRequest struct {
	Id     string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Fields *google_protobuf.Any `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PutRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutRequest) GetFields() *google_protobuf.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

type PutResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PutResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PutResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Document struct {
	Id     string               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Fields *google_protobuf.Any `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetFields() *google_protobuf.Any {
	if m != nil {
		return m.Fields
	}
	return nil
}

type UpdateRequest struct {
	Type     UpdateRequest_Type `protobuf:"varint,1,opt,name=type,enum=protobuf.UpdateRequest_Type" json:"type,omitempty"`
	Document *Document          `protobuf:"bytes,2,opt,name=document" json:"document,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UpdateRequest) GetType() UpdateRequest_Type {
	if m != nil {
		return m.Type
	}
	return UpdateRequest_NOOP
}

func (m *UpdateRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type BulkRequest struct {
	BatchSize      int32            `protobuf:"varint,1,opt,name=batch_size,json=batchSize" json:"batch_size,omitempty"`
	UpdateRequests []*UpdateRequest `protobuf:"bytes,2,rep,name=update_requests,json=updateRequests" json:"update_requests,omitempty"`
}

func (m *BulkRequest) Reset()                    { *m = BulkRequest{} }
func (m *BulkRequest) String() string            { return proto.CompactTextString(m) }
func (*BulkRequest) ProtoMessage()               {}
func (*BulkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *BulkRequest) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *BulkRequest) GetUpdateRequests() []*UpdateRequest {
	if m != nil {
		return m.UpdateRequests
	}
	return nil
}

type BulkResponse struct {
	PutCount    int32  `protobuf:"varint,1,opt,name=put_count,json=putCount" json:"put_count,omitempty"`
	DeleteCount int32  `protobuf:"varint,2,opt,name=delete_count,json=deleteCount" json:"delete_count,omitempty"`
	Success     bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	Message     string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *BulkResponse) Reset()                    { *m = BulkResponse{} }
func (m *BulkResponse) String() string            { return proto.CompactTextString(m) }
func (*BulkResponse) ProtoMessage()               {}
func (*BulkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *BulkResponse) GetPutCount() int32 {
	if m != nil {
		return m.PutCount
	}
	return 0
}

func (m *BulkResponse) GetDeleteCount() int32 {
	if m != nil {
		return m.DeleteCount
	}
	return 0
}

func (m *BulkResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *BulkResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SearchRequest struct {
	SearchRequest *google_protobuf.Any `protobuf:"bytes,1,opt,name=search_request,json=searchRequest" json:"search_request,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *SearchRequest) GetSearchRequest() *google_protobuf.Any {
	if m != nil {
		return m.SearchRequest
	}
	return nil
}

type SearchResponse struct {
	SearchResult *google_protobuf.Any `protobuf:"bytes,1,opt,name=search_result,json=searchResult" json:"search_result,omitempty"`
	Success      bool                 `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	Message      string               `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *SearchResponse) GetSearchResult() *google_protobuf.Any {
	if m != nil {
		return m.SearchResult
	}
	return nil
}

func (m *SearchResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SearchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Proposal struct {
	Type Proposal_Type        `protobuf:"varint,1,opt,name=type,enum=protobuf.Proposal_Type" json:"type,omitempty"`
	Data *google_protobuf.Any `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Proposal) Reset()                    { *m = Proposal{} }
func (m *Proposal) String() string            { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()               {}
func (*Proposal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Proposal) GetType() Proposal_Type {
	if m != nil {
		return m.Type
	}
	return Proposal_NOOP
}

func (m *Proposal) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "protobuf.Metadata")
	proto.RegisterType((*JoinRequest)(nil), "protobuf.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "protobuf.JoinResponse")
	proto.RegisterType((*LeaveRequest)(nil), "protobuf.LeaveRequest")
	proto.RegisterType((*LeaveResponse)(nil), "protobuf.LeaveResponse")
	proto.RegisterType((*Peer)(nil), "protobuf.Peer")
	proto.RegisterType((*PeersResponse)(nil), "protobuf.PeersResponse")
	proto.RegisterType((*SnapshotResponse)(nil), "protobuf.SnapshotResponse")
	proto.RegisterType((*GetRequest)(nil), "protobuf.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "protobuf.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "protobuf.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "protobuf.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "protobuf.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "protobuf.DeleteResponse")
	proto.RegisterType((*Document)(nil), "protobuf.Document")
	proto.RegisterType((*UpdateRequest)(nil), "protobuf.UpdateRequest")
	proto.RegisterType((*BulkRequest)(nil), "protobuf.BulkRequest")
	proto.RegisterType((*BulkResponse)(nil), "protobuf.BulkResponse")
	proto.RegisterType((*SearchRequest)(nil), "protobuf.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "protobuf.SearchResponse")
	proto.RegisterType((*Proposal)(nil), "protobuf.Proposal")
	proto.RegisterEnum("protobuf.UpdateRequest_Type", UpdateRequest_Type_name, UpdateRequest_Type_value)
	proto.RegisterEnum("protobuf.Proposal_Type", Proposal_Type_name, Proposal_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Data service

type DataClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
	Peers(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*PeersResponse, error)
	Snapshot(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SnapshotResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Bulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type dataClient struct {
	cc *grpc.ClientConn
}

func NewDataClient(cc *grpc.ClientConn) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Leave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Peers(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*PeersResponse, error) {
	out := new(PeersResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Peers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Snapshot(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Snapshot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Bulk(ctx context.Context, in *BulkRequest, opts ...grpc.CallOption) (*BulkResponse, error) {
	out := new(BulkResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Bulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/protobuf.Data/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Data service

type DataServer interface {
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
	Peers(context.Context, *google_protobuf1.Empty) (*PeersResponse, error)
	Snapshot(context.Context, *google_protobuf1.Empty) (*SnapshotResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Bulk(context.Context, *BulkRequest) (*BulkResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterDataServer(s *grpc.Server, srv DataServer) {
	s.RegisterService(&_Data_serviceDesc, srv)
}

func _Data_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Peers(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Snapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Snapshot(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Bulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Bulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Bulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Bulk(ctx, req.(*BulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Data_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Data/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Data_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Data_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Data_Leave_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _Data_Peers_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _Data_Snapshot_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Data_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Data_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Data_Delete_Handler,
		},
		{
			MethodName: "Bulk",
			Handler:    _Data_Bulk_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Data_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x51, 0x6f, 0xe2, 0x46,
	0x10, 0xc6, 0x60, 0x88, 0x19, 0x03, 0x45, 0xab, 0xbb, 0x84, 0x72, 0x57, 0xf5, 0xba, 0x6a, 0xa5,
	0x48, 0xad, 0xb8, 0x8a, 0x56, 0x3a, 0xb5, 0x55, 0xa5, 0x83, 0x40, 0xd3, 0xa4, 0x34, 0x41, 0x4e,
	0xd2, 0x87, 0xbe, 0x20, 0x07, 0x6f, 0x00, 0x15, 0x6c, 0xd7, 0xbb, 0xae, 0x44, 0x5e, 0xa2, 0xf6,
	0x57, 0xf4, 0xa1, 0xef, 0xfd, 0x9b, 0xd5, 0xee, 0x7a, 0xed, 0x35, 0x11, 0x51, 0x0a, 0x4f, 0xb0,
	0x33, 0xdf, 0xcc, 0x7c, 0x33, 0xbb, 0x33, 0x63, 0x00, 0xcf, 0x65, 0x6e, 0x27, 0x8c, 0x02, 0x16,
	0x20, 0x4b, 0xfc, 0xdc, 0xc6, 0x77, 0xed, 0x0f, 0x67, 0x41, 0x30, 0x5b, 0x92, 0xb7, 0x4a, 0xf0,
	0xd6, 0xf5, 0xd7, 0x12, 0xd4, 0x7e, 0xb5, 0xa9, 0x22, 0xab, 0x90, 0x25, 0x4a, 0x3c, 0x06, 0xeb,
	0x67, 0xc2, 0x5c, 0xee, 0x13, 0x7d, 0x02, 0xb5, 0x59, 0x14, 0x4e, 0x27, 0xae, 0xe7, 0x45, 0x84,
	0xd2, 0x56, 0xf1, 0x8d, 0x71, 0x5c, 0x75, 0x6c, 0x2e, 0xeb, 0x49, 0x11, 0x87, 0xcc, 0x19, 0x0b,
	0x53, 0x48, 0x49, 0x42, 0xb8, 0x2c, 0x81, 0xe0, 0x10, 0xec, 0xf3, 0x60, 0xe1, 0x3b, 0xe4, 0xf7,
	0x98, 0x50, 0x86, 0x8e, 0xe0, 0xc0, 0x0f, 0x3c, 0x32, 0x59, 0x78, 0x2d, 0x43, 0x80, 0x2b, 0xfc,
	0x78, 0xe6, 0xa1, 0x16, 0x1c, 0xe4, 0x03, 0xa9, 0x23, 0xea, 0x80, 0xb5, 0x4a, 0x38, 0x89, 0x00,
	0x76, 0x17, 0x75, 0x14, 0xf9, 0x8e, 0x62, 0xeb, 0xa4, 0x18, 0xdc, 0x87, 0x9a, 0x8c, 0x48, 0xc3,
	0xc0, 0xa7, 0x84, 0x7b, 0xa6, 0xf1, 0x74, 0xca, 0x3d, 0xf3, 0x90, 0x96, 0xa3, 0x8e, 0x5c, 0xb3,
	0x22, 0x94, 0xba, 0x33, 0xa2, 0x62, 0x26, 0x47, 0xdc, 0x83, 0xda, 0x88, 0xb8, 0x7f, 0x90, 0xdd,
	0x69, 0xe3, 0x13, 0xa8, 0x27, 0x2e, 0xf6, 0xe0, 0xf1, 0xa7, 0x01, 0xe6, 0x98, 0x90, 0x68, 0x97,
	0xba, 0x1d, 0x42, 0x65, 0x49, 0x5c, 0x8f, 0x44, 0xa2, 0x6a, 0x96, 0x93, 0x9c, 0x72, 0xf5, 0x34,
	0x9f, 0x51, 0xcf, 0x05, 0xd4, 0x39, 0x05, 0x9a, 0x26, 0xf2, 0x29, 0x94, 0x43, 0x2e, 0x68, 0x19,
	0x6f, 0x4a, 0xc7, 0x76, 0xb7, 0x91, 0x59, 0x73, 0x9c, 0x23, 0x95, 0x7a, 0xba, 0xc5, 0xad, 0xe9,
	0x96, 0xf2, 0xe9, 0xfe, 0x00, 0xcd, 0x2b, 0xdf, 0x0d, 0xe9, 0x3c, 0x60, 0x7b, 0x95, 0xed, 0x35,
	0xc0, 0x29, 0x61, 0xea, 0xf2, 0x1a, 0x50, 0x4c, 0xcb, 0x56, 0x5c, 0x78, 0xf8, 0x01, 0x6c, 0xa1,
	0x4d, 0x02, 0x6c, 0xa8, 0xd1, 0x17, 0x50, 0xb9, 0x5b, 0x90, 0xa5, 0x27, 0x79, 0xdb, 0xdd, 0x17,
	0x1d, 0xd9, 0x31, 0x59, 0x9a, 0x3d, 0x7f, 0xed, 0x24, 0x18, 0x9d, 0x5e, 0x69, 0x2b, 0x3d, 0x33,
	0x4f, 0xef, 0x1c, 0x60, 0x1c, 0x6f, 0xa3, 0xf7, 0xff, 0xe2, 0xe3, 0x1e, 0xd8, 0xc2, 0xd7, 0x1e,
	0xd5, 0xfa, 0x18, 0xea, 0x03, 0xb2, 0x24, 0x8c, 0x6c, 0x2b, 0xd8, 0x00, 0x1a, 0x0a, 0xb0, 0x47,
	0x98, 0x1f, 0xc1, 0x1a, 0x04, 0xd3, 0x78, 0x45, 0xfc, 0x7d, 0x73, 0xfe, 0xdb, 0x80, 0xfa, 0x4d,
	0xe8, 0xb9, 0x19, 0xe3, 0x2f, 0xc1, 0x64, 0xeb, 0x90, 0x08, 0x8f, 0x8d, 0xee, 0xeb, 0xcc, 0x2c,
	0x07, 0xeb, 0x5c, 0xaf, 0x43, 0xe2, 0x08, 0x24, 0xef, 0x02, 0x2f, 0x61, 0x93, 0xc4, 0xd4, 0xba,
	0x40, 0xf1, 0x74, 0x52, 0x0c, 0xfe, 0x0c, 0x4c, 0x6e, 0x8d, 0x2c, 0x30, 0x2f, 0x2e, 0x2f, 0xc7,
	0xcd, 0x02, 0x3a, 0x80, 0xd2, 0xf8, 0xe6, 0xba, 0x69, 0x20, 0x80, 0xca, 0x60, 0x38, 0x1a, 0x5e,
	0x0f, 0x9b, 0x45, 0xec, 0x83, 0xdd, 0x8f, 0x97, 0xbf, 0x29, 0x5e, 0x1f, 0x01, 0xdc, 0xba, 0x6c,
	0x3a, 0x9f, 0xd0, 0xc5, 0xbd, 0x64, 0x57, 0x76, 0xaa, 0x42, 0x72, 0xb5, 0xb8, 0x27, 0xe8, 0x3d,
	0x7c, 0x10, 0x0b, 0x82, 0x93, 0x48, 0x1a, 0xf0, 0xfc, 0x79, 0x4f, 0x1d, 0x6d, 0xc9, 0xc0, 0x69,
	0xc4, 0xfa, 0x91, 0xe2, 0xbf, 0x0c, 0xa8, 0xc9, 0x80, 0xc9, 0xcd, 0xbc, 0x82, 0x6a, 0x18, 0xb3,
	0xc9, 0x34, 0x88, 0x7d, 0x96, 0x04, 0xb4, 0xc2, 0x98, 0x9d, 0xf0, 0x33, 0x9f, 0xd7, 0x9e, 0xb8,
	0xc8, 0x44, 0x5f, 0x14, 0x7a, 0x5b, 0xca, 0x24, 0x64, 0x97, 0xf7, 0x3c, 0x82, 0xfa, 0x15, 0x71,
	0xa3, 0xe9, 0x5c, 0xa5, 0xfd, 0x1d, 0x34, 0xa8, 0x10, 0xa8, 0xbc, 0x04, 0x93, 0x6d, 0xd7, 0x5a,
	0xa7, 0xba, 0x31, 0x7e, 0x80, 0x86, 0xf2, 0x96, 0xe4, 0xf4, 0x0d, 0xd4, 0x53, 0x77, 0x34, 0x5e,
	0x3e, 0xed, 0xad, 0xa6, 0xbc, 0x71, 0xe4, 0x4e, 0x53, 0xe8, 0x5f, 0x03, 0xac, 0x71, 0x14, 0x84,
	0x01, 0x75, 0x97, 0xe8, 0xf3, 0xdc, 0xcb, 0xd2, 0xee, 0x45, 0x21, 0xf4, 0x47, 0x75, 0x0c, 0xa6,
	0x18, 0xab, 0x4f, 0x3d, 0x62, 0x81, 0xc0, 0x83, 0x67, 0x3d, 0x27, 0xae, 0xee, 0xdf, 0x8c, 0x7e,
	0x6a, 0x96, 0xf8, 0xbf, 0xf3, 0xcb, 0xb3, 0x8b, 0xa6, 0x89, 0xaa, 0x50, 0x1e, 0x0d, 0x7b, 0xbf,
	0x0c, 0x9b, 0xe5, 0xee, 0x3f, 0x26, 0x98, 0x03, 0xbe, 0xab, 0xdf, 0x81, 0xc9, 0x77, 0x1e, 0x7a,
	0x99, 0xc5, 0xd2, 0xb6, 0x6e, 0xfb, 0x70, 0x53, 0x2c, 0x0b, 0x8b, 0x0b, 0xe8, 0x5b, 0x28, 0x8b,
	0x2d, 0x85, 0x34, 0x88, 0xbe, 0xf9, 0xda, 0x47, 0x8f, 0xe4, 0xba, 0xad, 0x58, 0x0c, 0xe8, 0xf0,
	0x51, 0xa2, 0x43, 0xfe, 0x4d, 0xa1, 0xdb, 0xe6, 0x36, 0x08, 0x2e, 0xa0, 0xf7, 0x60, 0xa9, 0x49,
	0xbf, 0xd5, 0xbc, 0x9d, 0x09, 0x36, 0xb7, 0x02, 0x2e, 0xa0, 0xaf, 0xa1, 0x74, 0x4a, 0x18, 0x7a,
	0x91, 0x81, 0xb2, 0x91, 0xdf, 0x7e, 0xb9, 0x21, 0xd5, 0xad, 0xc6, 0x71, 0xce, 0x2a, 0x9b, 0xc4,
	0xba, 0x95, 0x36, 0x53, 0x71, 0x01, 0x7d, 0x0f, 0x15, 0x39, 0x00, 0x91, 0x96, 0x52, 0x6e, 0x66,
	0xb6, 0x5b, 0x8f, 0x15, 0xa9, 0xf9, 0x3b, 0x30, 0x79, 0x8f, 0xea, 0xb7, 0xa3, 0x0d, 0x09, 0xfd,
	0x76, 0xf4, 0x56, 0x96, 0x71, 0x65, 0x2b, 0xe8, 0x71, 0x73, 0xad, 0xa6, 0xc7, 0xcd, 0x77, 0x0d,
	0x2e, 0xf4, 0xe1, 0xd7, 0xf4, 0x8b, 0xf0, 0xb6, 0x22, 0xfe, 0x7d, 0xf5, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xd8, 0x83, 0xf7, 0x30, 0x0a, 0x00, 0x00,
}