// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpcServer.proto

package grpsServer

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

type CreateTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAddress []byte `protobuf:"bytes,1,opt,name=senderAddress,proto3" json:"senderAddress,omitempty"`
	Data          []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	PubKey        []byte `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Signature     []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce         int64  `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Cid           []byte `protobuf:"bytes,6,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *CreateTx) Reset() {
	*x = CreateTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTx) ProtoMessage() {}

func (x *CreateTx) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTx.ProtoReflect.Descriptor instead.
func (*CreateTx) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTx) GetSenderAddress() []byte {
	if x != nil {
		return x.SenderAddress
	}
	return nil
}

func (x *CreateTx) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateTx) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *CreateTx) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CreateTx) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *CreateTx) GetCid() []byte {
	if x != nil {
		return x.Cid
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAddress []byte `protobuf:"bytes,1,opt,name=senderAddress,proto3" json:"senderAddress,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetSenderAddress() []byte {
	if x != nil {
		return x.SenderAddress
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash []byte `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentHash []byte `protobuf:"bytes,1,opt,name=documentHash,proto3" json:"documentHash,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{4}
}

func (x *Document) GetDocumentHash() []byte {
	if x != nil {
		return x.DocumentHash
	}
	return nil
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	HashPrevBlock  []byte `protobuf:"bytes,2,opt,name=HashPrevBlock,proto3" json:"HashPrevBlock,omitempty"`
	HashMerkleRoot []byte `protobuf:"bytes,3,opt,name=HashMerkleRoot,proto3" json:"HashMerkleRoot,omitempty"`
	Time           int64  `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{5}
}

func (x *BlockHeader) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BlockHeader) GetHashPrevBlock() []byte {
	if x != nil {
		return x.HashPrevBlock
	}
	return nil
}

func (x *BlockHeader) GetHashMerkleRoot() []byte {
	if x != nil {
		return x.HashMerkleRoot
	}
	return nil
}

func (x *BlockHeader) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocksize    uint64       `protobuf:"varint,1,opt,name=blocksize,proto3" json:"blocksize,omitempty"`
	Blockheader  *BlockHeader `protobuf:"bytes,2,opt,name=blockheader,proto3" json:"blockheader,omitempty"`
	TxCounter    uint64       `protobuf:"varint,3,opt,name=TxCounter,proto3" json:"TxCounter,omitempty"`
	Transactions []*CreateTx  `protobuf:"bytes,4,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcServer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_grpcServer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_grpcServer_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetBlocksize() uint64 {
	if x != nil {
		return x.Blocksize
	}
	return 0
}

func (x *Block) GetBlockheader() *BlockHeader {
	if x != nil {
		return x.Blockheader
	}
	return nil
}

func (x *Block) GetTxCounter() uint64 {
	if x != nil {
		return x.TxCounter
	}
	return 0
}

func (x *Block) GetTransactions() []*CreateTx {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_grpcServer_proto protoreflect.FileDescriptor

var file_grpcServer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa2, 0x01, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x63, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x25, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2e, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72,
	0x65, 0x76, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x48, 0x61, 0x73, 0x68, 0x4d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x0a,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x54, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xba, 0x02, 0x0a, 0x08, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x78, 0x54, 0x6f,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x09, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x78, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x78, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x05, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x78, 0x42, 0x79,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x09, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcServer_proto_rawDescOnce sync.Once
	file_grpcServer_proto_rawDescData = file_grpcServer_proto_rawDesc
)

func file_grpcServer_proto_rawDescGZIP() []byte {
	file_grpcServer_proto_rawDescOnce.Do(func() {
		file_grpcServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcServer_proto_rawDescData)
	})
	return file_grpcServer_proto_rawDescData
}

var file_grpcServer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpcServer_proto_goTypes = []interface{}{
	(*CreateTx)(nil),       // 0: CreateTx
	(*CreateResponse)(nil), // 1: CreateResponse
	(*User)(nil),           // 2: User
	(*Transaction)(nil),    // 3: Transaction
	(*Document)(nil),       // 4: Document
	(*BlockHeader)(nil),    // 5: BlockHeader
	(*Block)(nil),          // 6: Block
	(*emptypb.Empty)(nil),  // 7: google.protobuf.Empty
}
var file_grpcServer_proto_depIdxs = []int32{
	5, // 0: Block.blockheader:type_name -> BlockHeader
	0, // 1: Block.transactions:type_name -> CreateTx
	0, // 2: Invoicer.AddTxToBlockchain:input_type -> CreateTx
	7, // 3: Invoicer.GetLastBlock:input_type -> google.protobuf.Empty
	7, // 4: Invoicer.GetBlockchain:input_type -> google.protobuf.Empty
	2, // 5: Invoicer.GetUserTxHistory:input_type -> User
	3, // 6: Invoicer.GetTxByHash:input_type -> Transaction
	4, // 7: Invoicer.FindDocumentByHash:input_type -> Document
	1, // 8: Invoicer.AddTxToBlockchain:output_type -> CreateResponse
	1, // 9: Invoicer.GetLastBlock:output_type -> CreateResponse
	1, // 10: Invoicer.GetBlockchain:output_type -> CreateResponse
	1, // 11: Invoicer.GetUserTxHistory:output_type -> CreateResponse
	1, // 12: Invoicer.GetTxByHash:output_type -> CreateResponse
	1, // 13: Invoicer.FindDocumentByHash:output_type -> CreateResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpcServer_proto_init() }
func file_grpcServer_proto_init() {
	if File_grpcServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTx); i {
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
		file_grpcServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_grpcServer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_grpcServer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_grpcServer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_grpcServer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_grpcServer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_grpcServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcServer_proto_goTypes,
		DependencyIndexes: file_grpcServer_proto_depIdxs,
		MessageInfos:      file_grpcServer_proto_msgTypes,
	}.Build()
	File_grpcServer_proto = out.File
	file_grpcServer_proto_rawDesc = nil
	file_grpcServer_proto_goTypes = nil
	file_grpcServer_proto_depIdxs = nil
}