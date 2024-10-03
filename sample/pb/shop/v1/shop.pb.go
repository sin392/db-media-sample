// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: shop/v1/shop.proto

package shop

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Location
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 県
	Prefecture string `protobuf:"bytes,1,opt,name=prefecture,proto3" json:"prefecture,omitempty"`
	// 市
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	// 住所
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetPrefecture() string {
	if x != nil {
		return x.Prefecture
	}
	return ""
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Menu
type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 価格
	Price int32 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	// 商品説明
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{1}
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Menu) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

// Shop
type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 固有ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 店舗名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 住所
	Location *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// 電話番号
	Tel string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
	// 画像URL
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// サイトURL
	SiteUrl string `protobuf:"bytes,6,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
	// 評価
	Rating float32 `protobuf:"fixed32,7,opt,name=rating,proto3" json:"rating,omitempty"`
	// タグのリスト
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// メニューのリスト
	Menus []*Menu `protobuf:"bytes,9,rep,name=menus,proto3" json:"menus,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{2}
}

func (x *Shop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Shop) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *Shop) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Shop) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *Shop) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Shop) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Shop) GetMenus() []*Menu {
	if x != nil {
		return x.Menus
	}
	return nil
}

// FindShopByNameRequest
type FindShopByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 店舗名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FindShopByNameRequest) Reset() {
	*x = FindShopByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindShopByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindShopByNameRequest) ProtoMessage() {}

func (x *FindShopByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindShopByNameRequest.ProtoReflect.Descriptor instead.
func (*FindShopByNameRequest) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{3}
}

func (x *FindShopByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// FindShopByNameResponse
type FindShopByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 固有ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 店舗名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 住所
	Location *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// 電話番号
	Tel string `protobuf:"bytes,4,opt,name=tel,proto3" json:"tel,omitempty"`
	// 画像URL
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// サイトURL
	SiteUrl string `protobuf:"bytes,6,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
	// 評価
	Rating float32 `protobuf:"fixed32,7,opt,name=rating,proto3" json:"rating,omitempty"`
	// タグのリスト
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// メニューのリスト
	Menus []*Menu `protobuf:"bytes,9,rep,name=menus,proto3" json:"menus,omitempty"`
}

func (x *FindShopByNameResponse) Reset() {
	*x = FindShopByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_v1_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindShopByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindShopByNameResponse) ProtoMessage() {}

func (x *FindShopByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_v1_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindShopByNameResponse.ProtoReflect.Descriptor instead.
func (*FindShopByNameResponse) Descriptor() ([]byte, []int) {
	return file_shop_v1_shop_proto_rawDescGZIP(), []int{4}
}

func (x *FindShopByNameResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FindShopByNameResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindShopByNameResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *FindShopByNameResponse) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *FindShopByNameResponse) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *FindShopByNameResponse) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *FindShopByNameResponse) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *FindShopByNameResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *FindShopByNameResponse) GetMenus() []*Menu {
	if x != nil {
		return x.Menus
	}
	return nil
}

var File_shop_v1_shop_proto protoreflect.FileDescriptor

var file_shop_v1_shop_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0xf4, 0x01, 0x0a, 0x04,
	0x53, 0x68, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x23, 0x0a,
	0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x05, 0x6d, 0x65, 0x6e,
	0x75, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x86, 0x02, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x05, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x32, 0x73, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x53,
	0x68, 0x6f, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x42, 0xd3, 0x01,
	0x92, 0x41, 0x41, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x41, 0x50,
	0x49, 0x22, 0x12, 0x1a, 0x10, 0x73, 0x69, 0x6e, 0x33, 0x39, 0x32, 0x40, 0x67, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x19, 0x0a, 0x0a, 0x4d, 0x79, 0x20, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x12, 0x0b, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2e, 0x74, 0x78, 0x74,
	0x32, 0x02, 0x76, 0x31, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x53, 0x68, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6e, 0x33, 0x39,
	0x32, 0x2f, 0x64, 0x62, 0x2d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x6f, 0x70,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x6f, 0x70, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02,
	0x07, 0x53, 0x68, 0x6f, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x53, 0x68, 0x6f, 0x70, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x13, 0x53, 0x68, 0x6f, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x68, 0x6f, 0x70, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_v1_shop_proto_rawDescOnce sync.Once
	file_shop_v1_shop_proto_rawDescData = file_shop_v1_shop_proto_rawDesc
)

func file_shop_v1_shop_proto_rawDescGZIP() []byte {
	file_shop_v1_shop_proto_rawDescOnce.Do(func() {
		file_shop_v1_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_v1_shop_proto_rawDescData)
	})
	return file_shop_v1_shop_proto_rawDescData
}

var file_shop_v1_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shop_v1_shop_proto_goTypes = []any{
	(*Location)(nil),               // 0: shop.v1.Location
	(*Menu)(nil),                   // 1: shop.v1.Menu
	(*Shop)(nil),                   // 2: shop.v1.Shop
	(*FindShopByNameRequest)(nil),  // 3: shop.v1.FindShopByNameRequest
	(*FindShopByNameResponse)(nil), // 4: shop.v1.FindShopByNameResponse
}
var file_shop_v1_shop_proto_depIdxs = []int32{
	0, // 0: shop.v1.Shop.location:type_name -> shop.v1.Location
	1, // 1: shop.v1.Shop.menus:type_name -> shop.v1.Menu
	0, // 2: shop.v1.FindShopByNameResponse.location:type_name -> shop.v1.Location
	1, // 3: shop.v1.FindShopByNameResponse.menus:type_name -> shop.v1.Menu
	3, // 4: shop.v1.ShopService.FindShopByName:input_type -> shop.v1.FindShopByNameRequest
	4, // 5: shop.v1.ShopService.FindShopByName:output_type -> shop.v1.FindShopByNameResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_shop_v1_shop_proto_init() }
func file_shop_v1_shop_proto_init() {
	if File_shop_v1_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_v1_shop_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Location); i {
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
		file_shop_v1_shop_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Menu); i {
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
		file_shop_v1_shop_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Shop); i {
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
		file_shop_v1_shop_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FindShopByNameRequest); i {
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
		file_shop_v1_shop_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FindShopByNameResponse); i {
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
			RawDescriptor: file_shop_v1_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_v1_shop_proto_goTypes,
		DependencyIndexes: file_shop_v1_shop_proto_depIdxs,
		MessageInfos:      file_shop_v1_shop_proto_msgTypes,
	}.Build()
	File_shop_v1_shop_proto = out.File
	file_shop_v1_shop_proto_rawDesc = nil
	file_shop_v1_shop_proto_goTypes = nil
	file_shop_v1_shop_proto_depIdxs = nil
}
