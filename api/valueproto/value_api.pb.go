// Code generated by protoc-gen-go.
// source: value_api.proto
// DO NOT EDIT!

/*
Package valueproto is a generated protocol buffer package.

It is generated from these files:
	value_api.proto

It has these top-level messages:
	EmptyMsg
	KeyValue
	Key
	WriteResponse
	LookupResponse
	GetResponse
	DelResponse
*/
package valueproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmptyMsg struct {
}

func (m *EmptyMsg) Reset()         { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string { return proto.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()    {}

type KeyValue struct {
	A     uint64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B     uint64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value,omitempty"`
	Tsm   int64  `protobuf:"varint,4,opt,name=tsm" json:"tsm,omitempty"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

type Key struct {
	A   uint64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B   uint64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
	Tsm int64  `protobuf:"varint,3,opt,name=tsm" json:"tsm,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}

type WriteResponse struct {
	Tsm int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

type LookupResponse struct {
	Tsm    int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Length uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Err    string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}

type GetResponse struct {
	Tsm   int64  `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Err   string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}

type DelResponse struct {
	Tsm int64 `protobuf:"varint,1,opt,name=tsm" json:"tsm,omitempty"`
	Err bool  `protobuf:"varint,2,opt,name=err" json:"err,omitempty"`
}

func (m *DelResponse) Reset()         { *m = DelResponse{} }
func (m *DelResponse) String() string { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*EmptyMsg)(nil), "valueproto.EmptyMsg")
	proto.RegisterType((*KeyValue)(nil), "valueproto.KeyValue")
	proto.RegisterType((*Key)(nil), "valueproto.Key")
	proto.RegisterType((*WriteResponse)(nil), "valueproto.WriteResponse")
	proto.RegisterType((*LookupResponse)(nil), "valueproto.LookupResponse")
	proto.RegisterType((*GetResponse)(nil), "valueproto.GetResponse")
	proto.RegisterType((*DelResponse)(nil), "valueproto.DelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for ValueStore service

type ValueStoreClient interface {
	Write(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*WriteResponse, error)
	Lookup(ctx context.Context, in *Key, opts ...grpc.CallOption) (*LookupResponse, error)
	Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DelResponse, error)
}

type valueStoreClient struct {
	cc *grpc.ClientConn
}

func NewValueStoreClient(cc *grpc.ClientConn) ValueStoreClient {
	return &valueStoreClient{cc}
}

func (c *valueStoreClient) Write(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/valueproto.ValueStore/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueStoreClient) Lookup(ctx context.Context, in *Key, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/valueproto.ValueStore/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueStoreClient) Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/valueproto.ValueStore/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *valueStoreClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := grpc.Invoke(ctx, "/valueproto.ValueStore/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ValueStore service

type ValueStoreServer interface {
	Write(context.Context, *KeyValue) (*WriteResponse, error)
	Lookup(context.Context, *Key) (*LookupResponse, error)
	Read(context.Context, *Key) (*GetResponse, error)
	Delete(context.Context, *Key) (*DelResponse, error)
}

func RegisterValueStoreServer(s *grpc.Server, srv ValueStoreServer) {
	s.RegisterService(&_ValueStore_serviceDesc, srv)
}

func _ValueStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ValueStoreServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ValueStore_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ValueStoreServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ValueStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ValueStoreServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ValueStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ValueStoreServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ValueStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "valueproto.ValueStore",
	HandlerType: (*ValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _ValueStore_Write_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _ValueStore_Lookup_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ValueStore_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ValueStore_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
