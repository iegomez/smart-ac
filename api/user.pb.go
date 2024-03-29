// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type LoginRequest struct {
	// Username of the user.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Password of the user.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	// The JWT tag to be used to access admin capabilities.
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type ProfileResponse struct {
	// User object.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileResponse) Reset()         { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()    {}
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{2}
}
func (m *ProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileResponse.Unmarshal(m, b)
}
func (m *ProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileResponse.Marshal(b, m, deterministic)
}
func (dst *ProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileResponse.Merge(dst, src)
}
func (m *ProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileResponse.Size(m)
}
func (m *ProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileResponse proto.InternalMessageInfo

func (m *ProfileResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	// User ID.
	// Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `protobuf:"varint,3,opt,name=session_ttl,json=sessionTTL,proto3" json:"session_ttl,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin              bool     `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetSessionTtl() int32 {
	if m != nil {
		return m.SessionTtl
	}
	return 0
}

func (m *User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

type UserListItem struct {
	// User ID.
	// Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `protobuf:"varint,3,opt,name=session_ttl,json=sessionTTL,proto3" json:"session_ttl,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserListItem) Reset()         { *m = UserListItem{} }
func (m *UserListItem) String() string { return proto.CompactTextString(m) }
func (*UserListItem) ProtoMessage()    {}
func (*UserListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{4}
}
func (m *UserListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListItem.Unmarshal(m, b)
}
func (m *UserListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListItem.Marshal(b, m, deterministic)
}
func (dst *UserListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListItem.Merge(dst, src)
}
func (m *UserListItem) XXX_Size() int {
	return xxx_messageInfo_UserListItem.Size(m)
}
func (m *UserListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListItem.DiscardUnknown(m)
}

var xxx_messageInfo_UserListItem proto.InternalMessageInfo

func (m *UserListItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserListItem) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserListItem) GetSessionTtl() int32 {
	if m != nil {
		return m.SessionTtl
	}
	return 0
}

func (m *UserListItem) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *UserListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateUserRequest struct {
	// User object to create.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Password of the user.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{5}
}
func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (dst *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(dst, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{6}
}
func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (dst *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(dst, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserRequest struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{7}
}
func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(dst, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserResponse struct {
	// User object.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{8}
}
func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (dst *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(dst, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetUserResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateUserRequest struct {
	// User object to update.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{9}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	// User ID.
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{10}
}
func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(dst, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListUserRequest struct {
	// Max number of user to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// When provided, the given string will be used to search on username.
	Search               string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequest) Reset()         { *m = ListUserRequest{} }
func (m *ListUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserRequest) ProtoMessage()    {}
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{11}
}
func (m *ListUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequest.Unmarshal(m, b)
}
func (m *ListUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequest.Marshal(b, m, deterministic)
}
func (dst *ListUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequest.Merge(dst, src)
}
func (m *ListUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserRequest.Size(m)
}
func (m *ListUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequest proto.InternalMessageInfo

func (m *ListUserRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListUserRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListUserRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

type ListUserResponse struct {
	// Total number of users.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Result-set.
	Result               []*UserListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListUserResponse) Reset()         { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()    {}
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{12}
}
func (m *ListUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponse.Unmarshal(m, b)
}
func (m *ListUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponse.Marshal(b, m, deterministic)
}
func (dst *ListUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponse.Merge(dst, src)
}
func (m *ListUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserResponse.Size(m)
}
func (m *ListUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponse proto.InternalMessageInfo

func (m *ListUserResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListUserResponse) GetResult() []*UserListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateUserPasswordRequest struct {
	// User ID.
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// New pasword.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserPasswordRequest) Reset()         { *m = UpdateUserPasswordRequest{} }
func (m *UpdateUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserPasswordRequest) ProtoMessage()    {}
func (*UpdateUserPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_537a6b0b8fc35c30, []int{13}
}
func (m *UpdateUserPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserPasswordRequest.Unmarshal(m, b)
}
func (m *UpdateUserPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserPasswordRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserPasswordRequest.Merge(dst, src)
}
func (m *UpdateUserPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserPasswordRequest.Size(m)
}
func (m *UpdateUserPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserPasswordRequest proto.InternalMessageInfo

func (m *UpdateUserPasswordRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateUserPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
	proto.RegisterType((*ProfileResponse)(nil), "api.ProfileResponse")
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*UserListItem)(nil), "api.UserListItem")
	proto.RegisterType((*CreateUserRequest)(nil), "api.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "api.CreateUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "api.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "api.GetUserResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "api.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "api.DeleteUserRequest")
	proto.RegisterType((*ListUserRequest)(nil), "api.ListUserRequest")
	proto.RegisterType((*ListUserResponse)(nil), "api.ListUserResponse")
	proto.RegisterType((*UpdateUserPasswordRequest)(nil), "api.UpdateUserPasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Get user list.
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Create a new user.
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Update an existing user.
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete a user.
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// UpdatePassword updates a password.
	UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Log in a user
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Get the current user's profile
	Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdateUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/api.UserService/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Get user list.
	List(context.Context, *ListUserRequest) (*ListUserResponse, error)
	// Get data for a particular user.
	Get(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Create a new user.
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Update an existing user.
	Update(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	// Delete a user.
	Delete(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	// UpdatePassword updates a password.
	UpdatePassword(context.Context, *UpdateUserPasswordRequest) (*empty.Empty, error)
	// Log in a user
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Get the current user's profile
	Profile(context.Context, *empty.Empty) (*ProfileResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdateUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserService/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_537a6b0b8fc35c30) }

var fileDescriptor_user_537a6b0b8fc35c30 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xc4, 0x6d, 0x26, 0xa5, 0x69, 0xb6, 0x69, 0x93, 0x1a, 0xda, 0x84, 0x85, 0x43,
	0xe8, 0x21, 0x41, 0xe1, 0x04, 0xb7, 0xaa, 0x40, 0x55, 0xa9, 0xaa, 0x82, 0x69, 0x41, 0x5c, 0x88,
	0xdc, 0x78, 0x53, 0x16, 0xd9, 0x5e, 0xe3, 0xdd, 0x50, 0x21, 0xd4, 0x0b, 0xbf, 0xc0, 0x3f, 0xf0,
	0x17, 0x7c, 0x05, 0xbf, 0xc0, 0x91, 0x8f, 0x40, 0xbb, 0x5e, 0xa7, 0x8e, 0xa3, 0x86, 0xc2, 0x81,
	0x5b, 0x66, 0xf6, 0xed, 0x9b, 0x99, 0x37, 0x2f, 0x5e, 0x80, 0x09, 0x27, 0x71, 0x37, 0x8a, 0x99,
	0x60, 0xc8, 0x74, 0x23, 0x6a, 0xdf, 0x39, 0x67, 0xec, 0xdc, 0x27, 0x3d, 0x37, 0xa2, 0x3d, 0x37,
	0x0c, 0x99, 0x70, 0x05, 0x65, 0x21, 0x4f, 0x20, 0x76, 0x4b, 0x9f, 0xaa, 0xe8, 0x6c, 0x32, 0xee,
	0x09, 0x1a, 0x10, 0x2e, 0xdc, 0x20, 0xd2, 0x80, 0xdb, 0x79, 0x00, 0x09, 0x22, 0xf1, 0x29, 0x39,
	0xc4, 0xcf, 0x61, 0xe5, 0x88, 0x9d, 0xd3, 0xd0, 0x21, 0x1f, 0x26, 0x84, 0x0b, 0x64, 0xc3, 0xb2,
	0x2c, 0x1f, 0xba, 0x01, 0x69, 0x1a, 0x6d, 0xa3, 0x53, 0x76, 0xa6, 0xb1, 0x3c, 0x8b, 0x5c, 0xce,
	0x2f, 0x58, 0xec, 0x35, 0x0b, 0xc9, 0x59, 0x1a, 0xe3, 0xbb, 0x70, 0x4b, 0xf3, 0xf0, 0x88, 0x85,
	0x9c, 0xa0, 0x35, 0x30, 0xdf, 0x5f, 0x08, 0xcd, 0x21, 0x7f, 0xe2, 0x87, 0x50, 0x1d, 0xc4, 0x6c,
	0x4c, 0x7d, 0x32, 0x05, 0x6d, 0x43, 0x51, 0xb2, 0x2b, 0x54, 0xa5, 0x5f, 0xee, 0xba, 0x11, 0xed,
	0x9e, 0x72, 0x12, 0x3b, 0x2a, 0x8d, 0x43, 0x28, 0xca, 0x08, 0xad, 0x42, 0x81, 0x7a, 0x0a, 0x64,
	0x3a, 0x05, 0xea, 0xcd, 0x34, 0x59, 0xc8, 0x35, 0xd9, 0x82, 0x0a, 0x27, 0x9c, 0x53, 0x16, 0x0e,
	0x85, 0xf0, 0x9b, 0x66, 0xdb, 0xe8, 0x94, 0x1c, 0xd0, 0xa9, 0x93, 0x93, 0x23, 0xb4, 0x05, 0xcb,
	0x94, 0x0f, 0x5d, 0x2f, 0xa0, 0x61, 0xb3, 0xd8, 0x36, 0x3a, 0xcb, 0xce, 0x12, 0xe5, 0x7b, 0x32,
	0xc4, 0xbf, 0x0c, 0x58, 0x91, 0x05, 0x8f, 0x28, 0x17, 0x87, 0x82, 0x04, 0xff, 0xab, 0x30, 0x7a,
	0x0c, 0x30, 0x8a, 0x89, 0x2b, 0x88, 0x37, 0x74, 0x45, 0xb3, 0xa4, 0xd4, 0xb0, 0xbb, 0xc9, 0xde,
	0xba, 0xe9, 0xde, 0xba, 0x27, 0xe9, 0x62, 0x9d, 0xb2, 0x46, 0xef, 0x09, 0x79, 0x75, 0x12, 0x79,
	0xe9, 0x55, 0xeb, 0xcf, 0x57, 0x35, 0x7a, 0x4f, 0xe0, 0x63, 0xa8, 0xed, 0x2b, 0x1e, 0x25, 0xb9,
	0x36, 0xc0, 0xe2, 0x95, 0x2c, 0xf4, 0xc0, 0x7d, 0x40, 0x59, 0x3e, 0xbd, 0xe3, 0x9c, 0x86, 0xb8,
	0x0d, 0xab, 0x07, 0x44, 0x64, 0x4b, 0xe6, 0x11, 0xdf, 0x0c, 0xa8, 0x4e, 0x21, 0x37, 0x72, 0x4a,
	0x4e, 0xc0, 0xc2, 0xbf, 0x0b, 0x68, 0xfe, 0x8d, 0x80, 0x7d, 0xa8, 0x9d, 0xaa, 0xe0, 0xe6, 0x02,
	0xe2, 0x7b, 0x50, 0x7b, 0x4a, 0x7c, 0x32, 0x7b, 0x27, 0xaf, 0xc0, 0x6b, 0xa8, 0x4a, 0x0f, 0x66,
	0x21, 0x75, 0x28, 0xf9, 0x34, 0xa0, 0x42, 0xa3, 0x92, 0x00, 0x6d, 0x82, 0xc5, 0xc6, 0x63, 0x4e,
	0x92, 0x99, 0x4d, 0x47, 0x47, 0x32, 0xcf, 0x89, 0x1b, 0x8f, 0xde, 0xa9, 0x81, 0xca, 0x8e, 0x8e,
	0xf0, 0x5b, 0x58, 0xbb, 0x22, 0xd6, 0xd2, 0xb6, 0xa0, 0x22, 0x98, 0x70, 0xfd, 0xe1, 0x88, 0x4d,
	0xc2, 0x94, 0x1f, 0x54, 0x6a, 0x5f, 0x66, 0xd0, 0x03, 0xb0, 0x62, 0xc2, 0x27, 0xbe, 0x2c, 0x62,
	0x76, 0x2a, 0xfd, 0xda, 0x74, 0xa6, 0xf4, 0x8f, 0xe2, 0x68, 0x00, 0x1e, 0xc0, 0xd6, 0x95, 0x22,
	0x03, 0x6d, 0x8c, 0x74, 0x84, 0x06, 0x2c, 0x49, 0x09, 0x86, 0xd3, 0x51, 0x2d, 0x19, 0x1e, 0x7a,
	0x8b, 0x4c, 0xd5, 0xff, 0x5e, 0x82, 0x8a, 0x24, 0x7b, 0x49, 0xe2, 0x8f, 0x74, 0x44, 0xd0, 0x01,
	0x14, 0x65, 0x55, 0x54, 0x57, 0x4d, 0xe4, 0x54, 0xb2, 0x37, 0x72, 0xd9, 0x64, 0x44, 0x8c, 0xbe,
	0xfc, 0xf8, 0xf9, 0xb5, 0xb0, 0x82, 0x40, 0x7d, 0x43, 0x65, 0x55, 0x8e, 0x0e, 0xc1, 0x3c, 0x20,
	0x02, 0xad, 0xab, 0x1b, 0xb3, 0x8e, 0xb4, 0xeb, 0xb3, 0x49, 0xcd, 0xd2, 0x50, 0x2c, 0x35, 0x54,
	0xbd, 0x62, 0xe9, 0x7d, 0xa6, 0xde, 0x25, 0x1a, 0x80, 0x95, 0x18, 0x1f, 0x6d, 0xaa, 0x8b, 0x73,
	0xff, 0x2a, 0xbb, 0x31, 0x97, 0xd7, 0x9c, 0x1b, 0x8a, 0xb3, 0x8a, 0x33, 0x9d, 0x3d, 0x31, 0x76,
	0xd1, 0x1b, 0xb0, 0x12, 0x1d, 0x35, 0xe3, 0x9c, 0xcd, 0xec, 0xcd, 0x39, 0x8b, 0x3e, 0x93, 0x9f,
	0x75, 0xdc, 0x52, 0x84, 0x5b, 0x76, 0x3d, 0xdb, 0xa4, 0x7a, 0x51, 0xa8, 0x77, 0x29, 0xa9, 0x5f,
	0x80, 0x95, 0x18, 0x50, 0x53, 0xcf, 0xb9, 0xf1, 0x5a, 0x6a, 0x3d, 0xff, 0xee, 0xdc, 0xfc, 0x31,
	0xac, 0x26, 0x0d, 0xa6, 0x1b, 0x47, 0x3b, 0xb9, 0xae, 0x73, 0x56, 0xb8, 0xb6, 0x44, 0x47, 0x95,
	0xc0, 0xf6, 0x76, 0xbe, 0xfb, 0x21, 0xf5, 0x2e, 0x7b, 0xa9, 0x29, 0xe4, 0x18, 0xc7, 0x50, 0x52,
	0x0f, 0x0e, 0x4a, 0xdc, 0x98, 0x7d, 0xc4, 0x6c, 0x94, 0x4d, 0x69, 0xa1, 0x77, 0x14, 0x73, 0x13,
	0xaf, 0x2b, 0x66, 0x1a, 0x0a, 0xf9, 0xd9, 0xf6, 0x7b, 0xbe, 0x04, 0x49, 0xbe, 0x57, 0xb0, 0xa4,
	0x5f, 0x27, 0x74, 0x4d, 0x73, 0xda, 0x15, 0xb9, 0x37, 0x0c, 0x6f, 0x2b, 0xe2, 0x06, 0xda, 0x98,
	0x25, 0x8e, 0x12, 0xd8, 0x99, 0xa5, 0x48, 0x1e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x05,
	0x14, 0x30, 0xd6, 0x07, 0x00, 0x00,
}
