// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: services/balances.proto

package services

import (
	proto "github.com/ubtools/ubt/go/api/proto"
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

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId  *proto.NetworkId `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Address    string           `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	CurrencyId string           `protobuf:"bytes,3,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetNetworkId() *proto.NetworkId {
	if x != nil {
		return x.NetworkId
	}
	return nil
}

func (x *GetBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetBalanceRequest) GetCurrencyId() string {
	if x != nil {
		return x.CurrencyId
	}
	return ""
}

type BalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount     *proto.CurrencyAmount `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	LastUpdate uint64                `protobuf:"varint,3,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"` // last update block number
}

func (x *BalanceResponse) Reset() {
	*x = BalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceResponse) ProtoMessage() {}

func (x *BalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceResponse.ProtoReflect.Descriptor instead.
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BalanceResponse) GetAmount() *proto.CurrencyAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *BalanceResponse) GetLastUpdate() uint64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

type ListAccountBalancesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId   *proto.NetworkId `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Address     string           `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	CurrencyIds []string         `protobuf:"bytes,3,rep,name=currency_ids,json=currencyIds,proto3" json:"currency_ids,omitempty"`
}

func (x *ListAccountBalancesRequest) Reset() {
	*x = ListAccountBalancesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountBalancesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountBalancesRequest) ProtoMessage() {}

func (x *ListAccountBalancesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountBalancesRequest.ProtoReflect.Descriptor instead.
func (*ListAccountBalancesRequest) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccountBalancesRequest) GetNetworkId() *proto.NetworkId {
	if x != nil {
		return x.NetworkId
	}
	return nil
}

func (x *ListAccountBalancesRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ListAccountBalancesRequest) GetCurrencyIds() []string {
	if x != nil {
		return x.CurrencyIds
	}
	return nil
}

type ListAccountBalancesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amounts    []*proto.CurrencyAmount `protobuf:"bytes,1,rep,name=amounts,proto3" json:"amounts,omitempty"`
	Currencies []*proto.Currency       `protobuf:"bytes,2,rep,name=currencies,proto3" json:"currencies,omitempty"`
}

func (x *ListAccountBalancesResponse) Reset() {
	*x = ListAccountBalancesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountBalancesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountBalancesResponse) ProtoMessage() {}

func (x *ListAccountBalancesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountBalancesResponse.ProtoReflect.Descriptor instead.
func (*ListAccountBalancesResponse) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{3}
}

func (x *ListAccountBalancesResponse) GetAmounts() []*proto.CurrencyAmount {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *ListAccountBalancesResponse) GetCurrencies() []*proto.Currency {
	if x != nil {
		return x.Currencies
	}
	return nil
}

type ListCurrencyHoldersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId  *proto.NetworkId `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	CurrencyId string           `protobuf:"bytes,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
}

func (x *ListCurrencyHoldersRequest) Reset() {
	*x = ListCurrencyHoldersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCurrencyHoldersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrencyHoldersRequest) ProtoMessage() {}

func (x *ListCurrencyHoldersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrencyHoldersRequest.ProtoReflect.Descriptor instead.
func (*ListCurrencyHoldersRequest) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{4}
}

func (x *ListCurrencyHoldersRequest) GetNetworkId() *proto.NetworkId {
	if x != nil {
		return x.NetworkId
	}
	return nil
}

func (x *ListCurrencyHoldersRequest) GetCurrencyId() string {
	if x != nil {
		return x.CurrencyId
	}
	return ""
}

type ListCurrencyHoldersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balances map[string]*proto.Uint256 `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListCurrencyHoldersResponse) Reset() {
	*x = ListCurrencyHoldersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_balances_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCurrencyHoldersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrencyHoldersResponse) ProtoMessage() {}

func (x *ListCurrencyHoldersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_balances_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrencyHoldersResponse.ProtoReflect.Descriptor instead.
func (*ListCurrencyHoldersResponse) Descriptor() ([]byte, []int) {
	return file_services_balances_proto_rawDescGZIP(), []int{5}
}

func (x *ListCurrencyHoldersResponse) GetBalances() map[string]*proto.Uint256 {
	if x != nil {
		return x.Balances
	}
	return nil
}

var File_services_balances_proto protoreflect.FileDescriptor

var file_services_balances_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x62, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x75, 0x62, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x52, 0x09, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x2b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x88,
	0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x64, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x73, 0x22, 0x7b, 0x0a, 0x1b, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x62, 0x74, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x62,
	0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x49, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x49, 0x0a, 0x0d, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x62,
	0x74, 0x2e, 0x75, 0x69, 0x6e, 0x74, 0x32, 0x35, 0x36, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xb9, 0x02, 0x0a, 0x11, 0x55, 0x62, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x62, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x6c, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x28, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x75, 0x62, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x13, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x75, 0x62,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xa0, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x62, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x0d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x75, 0x62, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x75, 0x62, 0x74, 0x2f, 0x67,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x55, 0x53, 0x58, 0xaa, 0x02, 0x0c, 0x55, 0x62, 0x74,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x0c, 0x55, 0x62, 0x74, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x18, 0x55, 0x62, 0x74, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x55, 0x62, 0x74, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_balances_proto_rawDescOnce sync.Once
	file_services_balances_proto_rawDescData = file_services_balances_proto_rawDesc
)

func file_services_balances_proto_rawDescGZIP() []byte {
	file_services_balances_proto_rawDescOnce.Do(func() {
		file_services_balances_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_balances_proto_rawDescData)
	})
	return file_services_balances_proto_rawDescData
}

var file_services_balances_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_balances_proto_goTypes = []interface{}{
	(*GetBalanceRequest)(nil),           // 0: ubt.services.GetBalanceRequest
	(*BalanceResponse)(nil),             // 1: ubt.services.BalanceResponse
	(*ListAccountBalancesRequest)(nil),  // 2: ubt.services.ListAccountBalancesRequest
	(*ListAccountBalancesResponse)(nil), // 3: ubt.services.ListAccountBalancesResponse
	(*ListCurrencyHoldersRequest)(nil),  // 4: ubt.services.ListCurrencyHoldersRequest
	(*ListCurrencyHoldersResponse)(nil), // 5: ubt.services.ListCurrencyHoldersResponse
	nil,                                 // 6: ubt.services.ListCurrencyHoldersResponse.BalancesEntry
	(*proto.NetworkId)(nil),             // 7: ubt.NetworkId
	(*proto.CurrencyAmount)(nil),        // 8: ubt.CurrencyAmount
	(*proto.Currency)(nil),              // 9: ubt.Currency
	(*proto.Uint256)(nil),               // 10: ubt.uint256
}
var file_services_balances_proto_depIdxs = []int32{
	7,  // 0: ubt.services.GetBalanceRequest.network_id:type_name -> ubt.NetworkId
	8,  // 1: ubt.services.BalanceResponse.amount:type_name -> ubt.CurrencyAmount
	7,  // 2: ubt.services.ListAccountBalancesRequest.network_id:type_name -> ubt.NetworkId
	8,  // 3: ubt.services.ListAccountBalancesResponse.amounts:type_name -> ubt.CurrencyAmount
	9,  // 4: ubt.services.ListAccountBalancesResponse.currencies:type_name -> ubt.Currency
	7,  // 5: ubt.services.ListCurrencyHoldersRequest.network_id:type_name -> ubt.NetworkId
	6,  // 6: ubt.services.ListCurrencyHoldersResponse.balances:type_name -> ubt.services.ListCurrencyHoldersResponse.BalancesEntry
	10, // 7: ubt.services.ListCurrencyHoldersResponse.BalancesEntry.value:type_name -> ubt.uint256
	0,  // 8: ubt.services.UbtBalanceService.getBalance:input_type -> ubt.services.GetBalanceRequest
	2,  // 9: ubt.services.UbtBalanceService.listAccountBalances:input_type -> ubt.services.ListAccountBalancesRequest
	4,  // 10: ubt.services.UbtBalanceService.listCurrencyHolders:input_type -> ubt.services.ListCurrencyHoldersRequest
	1,  // 11: ubt.services.UbtBalanceService.getBalance:output_type -> ubt.services.BalanceResponse
	3,  // 12: ubt.services.UbtBalanceService.listAccountBalances:output_type -> ubt.services.ListAccountBalancesResponse
	5,  // 13: ubt.services.UbtBalanceService.listCurrencyHolders:output_type -> ubt.services.ListCurrencyHoldersResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_services_balances_proto_init() }
func file_services_balances_proto_init() {
	if File_services_balances_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_balances_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_services_balances_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceResponse); i {
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
		file_services_balances_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountBalancesRequest); i {
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
		file_services_balances_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountBalancesResponse); i {
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
		file_services_balances_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCurrencyHoldersRequest); i {
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
		file_services_balances_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCurrencyHoldersResponse); i {
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
			RawDescriptor: file_services_balances_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_balances_proto_goTypes,
		DependencyIndexes: file_services_balances_proto_depIdxs,
		MessageInfos:      file_services_balances_proto_msgTypes,
	}.Build()
	File_services_balances_proto = out.File
	file_services_balances_proto_rawDesc = nil
	file_services_balances_proto_goTypes = nil
	file_services_balances_proto_depIdxs = nil
}
