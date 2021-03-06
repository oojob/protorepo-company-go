// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/company/service.proto

package company

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protobuf "github.com/oojob/protobuf"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
// Comapny is used for storing companies data
type Company struct {
	Identity             *protobuf.Identifier `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Admin                string               `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty"`
	Url                  string               `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Logo                 string               `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	Employees            *protobuf.Range      `protobuf:"bytes,5,opt,name=employees,proto3" json:"employees,omitempty"`
	Place                *protobuf.Place      `protobuf:"bytes,6,opt,name=place,proto3" json:"place,omitempty"`
	FoundedYear          string               `protobuf:"bytes,7,opt,name=founded_year,json=foundedYear,proto3" json:"founded_year,omitempty"`
	HiringStatus         bool                 `protobuf:"varint,8,opt,name=hiring_status,json=hiringStatus,proto3" json:"hiring_status,omitempty"`
	Skills               []string             `protobuf:"bytes,9,rep,name=skills,proto3" json:"skills,omitempty"`
	Metadata             *protobuf.Metadata   `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ccc623271977312, []int{0}
}

func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetIdentity() *protobuf.Identifier {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Company) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *Company) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Company) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *Company) GetEmployees() *protobuf.Range {
	if m != nil {
		return m.Employees
	}
	return nil
}

func (m *Company) GetPlace() *protobuf.Place {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *Company) GetFoundedYear() string {
	if m != nil {
		return m.FoundedYear
	}
	return ""
}

func (m *Company) GetHiringStatus() bool {
	if m != nil {
		return m.HiringStatus
	}
	return false
}

func (m *Company) GetSkills() []string {
	if m != nil {
		return m.Skills
	}
	return nil
}

func (m *Company) GetMetadata() *protobuf.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

//
// CompanyAllResponse returns an array of companies
type CompanyAllResponse struct {
	Companies            []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CompanyAllResponse) Reset()         { *m = CompanyAllResponse{} }
func (m *CompanyAllResponse) String() string { return proto.CompactTextString(m) }
func (*CompanyAllResponse) ProtoMessage()    {}
func (*CompanyAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ccc623271977312, []int{1}
}

func (m *CompanyAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyAllResponse.Unmarshal(m, b)
}
func (m *CompanyAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyAllResponse.Marshal(b, m, deterministic)
}
func (m *CompanyAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyAllResponse.Merge(m, src)
}
func (m *CompanyAllResponse) XXX_Size() int {
	return xxx_messageInfo_CompanyAllResponse.Size(m)
}
func (m *CompanyAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyAllResponse proto.InternalMessageInfo

func (m *CompanyAllResponse) GetCompanies() []*Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func init() {
	proto.RegisterType((*Company)(nil), "company.Company")
	proto.RegisterType((*CompanyAllResponse)(nil), "company.CompanyAllResponse")
}

func init() { proto.RegisterFile("services/company/service.proto", fileDescriptor_7ccc623271977312) }

var fileDescriptor_7ccc623271977312 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0xf1, 0x52, 0xb7, 0xc9, 0x49, 0x33, 0x82, 0x18, 0xc3, 0x78, 0x6c, 0x78, 0x2d, 0x1d,
	0xbe, 0x99, 0x53, 0x32, 0xd8, 0x65, 0xd7, 0x90, 0x16, 0xb6, 0x8b, 0x42, 0x51, 0xd9, 0xc6, 0xae,
	0x8a, 0x62, 0x9f, 0xc4, 0x5a, 0x65, 0xcb, 0xb3, 0xe4, 0x81, 0xdf, 0x63, 0x2f, 0xb0, 0x37, 0x1d,
	0x91, 0x95, 0x7a, 0x2c, 0xd4, 0x19, 0x8c, 0xdd, 0x49, 0xbf, 0xbf, 0xff, 0x3f, 0xd2, 0xd1, 0x31,
	0xbc, 0x50, 0x58, 0x7e, 0xe7, 0x31, 0xaa, 0x49, 0x2c, 0xb3, 0x82, 0xe5, 0xf5, 0xc4, 0x0a, 0x51,
	0x51, 0x4a, 0x2d, 0xc9, 0x81, 0x95, 0xfd, 0x57, 0x2b, 0xae, 0xd3, 0x6a, 0x11, 0xc5, 0x32, 0x9b,
	0x48, 0xf9, 0x55, 0x2e, 0x26, 0x06, 0x58, 0x54, 0xcb, 0x89, 0xaa, 0x95, 0xc6, 0xac, 0x31, 0xf8,
	0x27, 0x0f, 0x73, 0x85, 0x60, 0x9b, 0x5c, 0x3f, 0x7c, 0x18, 0xcb, 0x50, 0xb3, 0x84, 0x69, 0x66,
	0xc9, 0x8e, 0xc2, 0x71, 0x55, 0x2a, 0x59, 0xee, 0xe6, 0x52, 0x64, 0x42, 0xa7, 0x0d, 0x77, 0xf4,
	0xb3, 0x07, 0x07, 0xf3, 0xe6, 0x52, 0x64, 0x06, 0x7d, 0x9e, 0x60, 0xae, 0xb9, 0xae, 0x3d, 0x27,
	0x70, 0xc2, 0xe1, 0xf4, 0x24, 0x6a, 0x63, 0x22, 0x13, 0x13, 0x6d, 0x62, 0xa2, 0x0f, 0x06, 0x5d,
	0x72, 0x2c, 0xe9, 0xbd, 0x8d, 0x3c, 0x01, 0x97, 0x25, 0x19, 0xcf, 0xbd, 0x47, 0x81, 0x13, 0x0e,
	0x68, 0xb3, 0x21, 0x63, 0xe8, 0x55, 0xa5, 0xf0, 0x7a, 0x46, 0x5b, 0x2f, 0x09, 0x81, 0x3d, 0x21,
	0x57, 0xd2, 0xdb, 0x33, 0x92, 0x59, 0x93, 0x33, 0x18, 0x60, 0x56, 0x08, 0x59, 0x23, 0x2a, 0xcf,
	0x35, 0xf5, 0x83, 0x8e, 0xfa, 0x94, 0xe5, 0x2b, 0xa4, 0xad, 0x85, 0xbc, 0x05, 0xd7, 0xf4, 0xd4,
	0xdb, 0xdf, 0xe9, 0xbd, 0x5e, 0x73, 0xb4, 0xc1, 0xc9, 0x4b, 0x38, 0x5c, 0xca, 0x2a, 0x4f, 0x30,
	0xb9, 0xad, 0x91, 0x95, 0xde, 0x81, 0x39, 0xd3, 0xd0, 0x6a, 0x5f, 0x90, 0x95, 0xe4, 0x18, 0x46,
	0x29, 0x2f, 0x79, 0xbe, 0xba, 0x55, 0x9a, 0xe9, 0x4a, 0x79, 0xfd, 0xc0, 0x09, 0xfb, 0xf4, 0xb0,
	0x11, 0x6f, 0x8c, 0x46, 0x9e, 0xc2, 0xbe, 0xba, 0xe3, 0x42, 0x28, 0x6f, 0x10, 0xf4, 0xc2, 0x01,
	0xb5, 0x3b, 0xf2, 0x0e, 0xfa, 0x9b, 0x47, 0xf4, 0xc0, 0x1c, 0xed, 0xb8, 0xe3, 0x68, 0x57, 0x16,
	0xa5, 0xf7, 0xa6, 0xa3, 0x0b, 0x20, 0xf6, 0x89, 0x66, 0x42, 0x50, 0x54, 0x85, 0xcc, 0x15, 0x92,
	0x08, 0x06, 0xcd, 0x34, 0x72, 0x54, 0x9e, 0x13, 0xf4, 0xc2, 0xe1, 0x74, 0x1c, 0xd9, 0xf9, 0x8c,
	0x2c, 0x4f, 0x5b, 0x64, 0xfa, 0xc3, 0x85, 0xc7, 0x56, 0xbe, 0x69, 0x86, 0x9a, 0x9c, 0xc3, 0x68,
	0x5e, 0x22, 0xd3, 0xb8, 0x99, 0x80, 0xad, 0x00, 0xff, 0x79, 0xe7, 0x04, 0x90, 0x33, 0x18, 0x52,
	0x64, 0xc9, 0xc6, 0xdf, 0x4d, 0xfb, 0x5b, 0xf1, 0xe4, 0x12, 0x46, 0xad, 0x9f, 0xa3, 0x22, 0x5d,
	0xaf, 0x76, 0x99, 0x15, 0xba, 0xde, 0x0e, 0x39, 0x75, 0xc8, 0x27, 0x18, 0xaf, 0x63, 0x66, 0x42,
	0xb4, 0x49, 0x5d, 0xb3, 0x7b, 0xcd, 0x56, 0x3c, 0x67, 0x9a, 0xcb, 0xdc, 0x7f, 0xf6, 0x67, 0xdc,
	0xef, 0x3d, 0x3e, 0x87, 0xd1, 0xc7, 0x22, 0xf9, 0x97, 0x06, 0x5d, 0xc1, 0xe8, 0x02, 0x05, 0xb6,
	0x09, 0x3b, 0x5a, 0xb4, 0x23, 0x6e, 0x09, 0xee, 0x3c, 0xc5, 0xf8, 0x8e, 0xbc, 0xee, 0xe0, 0xde,
	0x9b, 0x1f, 0xdc, 0x70, 0x14, 0xbf, 0x55, 0xa8, 0xb4, 0x1f, 0xfd, 0x2d, 0x6e, 0x2f, 0x9e, 0x82,
	0xfb, 0x99, 0xe9, 0x38, 0xfd, 0xcf, 0x75, 0x4e, 0x9d, 0xc5, 0xbe, 0xf9, 0xfe, 0xe6, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0x7a, 0x27, 0xc7, 0x7b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	CreateCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*protobuf.Id, error)
	ReadCompany(ctx context.Context, in *protobuf.Id, opts ...grpc.CallOption) (*Company, error)
	ReadCompanies(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (CompanyService_ReadCompaniesClient, error)
	ReadAllCompanies(ctx context.Context, in *protobuf.Pagination, opts ...grpc.CallOption) (*CompanyAllResponse, error)
	UpdateCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*protobuf.Id, error)
	DeleteCompany(ctx context.Context, in *protobuf.Id, opts ...grpc.CallOption) (*protobuf.Id, error)
	Check(ctx context.Context, in *protobuf.HealthCheckRequest, opts ...grpc.CallOption) (*protobuf.HealthCheckResponse, error)
	Watch(ctx context.Context, in *protobuf.HealthCheckRequest, opts ...grpc.CallOption) (CompanyService_WatchClient, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*protobuf.Id, error) {
	out := new(protobuf.Id)
	err := c.cc.Invoke(ctx, "/company.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ReadCompany(ctx context.Context, in *protobuf.Id, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/company.CompanyService/ReadCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ReadCompanies(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (CompanyService_ReadCompaniesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompanyService_serviceDesc.Streams[0], "/company.CompanyService/ReadCompanies", opts...)
	if err != nil {
		return nil, err
	}
	x := &companyServiceReadCompaniesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompanyService_ReadCompaniesClient interface {
	Recv() (*Company, error)
	grpc.ClientStream
}

type companyServiceReadCompaniesClient struct {
	grpc.ClientStream
}

func (x *companyServiceReadCompaniesClient) Recv() (*Company, error) {
	m := new(Company)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *companyServiceClient) ReadAllCompanies(ctx context.Context, in *protobuf.Pagination, opts ...grpc.CallOption) (*CompanyAllResponse, error) {
	out := new(CompanyAllResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/ReadAllCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *Company, opts ...grpc.CallOption) (*protobuf.Id, error) {
	out := new(protobuf.Id)
	err := c.cc.Invoke(ctx, "/company.CompanyService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *protobuf.Id, opts ...grpc.CallOption) (*protobuf.Id, error) {
	out := new(protobuf.Id)
	err := c.cc.Invoke(ctx, "/company.CompanyService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Check(ctx context.Context, in *protobuf.HealthCheckRequest, opts ...grpc.CallOption) (*protobuf.HealthCheckResponse, error) {
	out := new(protobuf.HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) Watch(ctx context.Context, in *protobuf.HealthCheckRequest, opts ...grpc.CallOption) (CompanyService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompanyService_serviceDesc.Streams[1], "/company.CompanyService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &companyServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompanyService_WatchClient interface {
	Recv() (*protobuf.HealthCheckResponse, error)
	grpc.ClientStream
}

type companyServiceWatchClient struct {
	grpc.ClientStream
}

func (x *companyServiceWatchClient) Recv() (*protobuf.HealthCheckResponse, error) {
	m := new(protobuf.HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	CreateCompany(context.Context, *Company) (*protobuf.Id, error)
	ReadCompany(context.Context, *protobuf.Id) (*Company, error)
	ReadCompanies(*protobuf.Empty, CompanyService_ReadCompaniesServer) error
	ReadAllCompanies(context.Context, *protobuf.Pagination) (*CompanyAllResponse, error)
	UpdateCompany(context.Context, *Company) (*protobuf.Id, error)
	DeleteCompany(context.Context, *protobuf.Id) (*protobuf.Id, error)
	Check(context.Context, *protobuf.HealthCheckRequest) (*protobuf.HealthCheckResponse, error)
	Watch(*protobuf.HealthCheckRequest, CompanyService_WatchServer) error
}

// UnimplementedCompanyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (*UnimplementedCompanyServiceServer) CreateCompany(ctx context.Context, req *Company) (*protobuf.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) ReadCompany(ctx context.Context, req *protobuf.Id) (*Company, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) ReadCompanies(req *protobuf.Empty, srv CompanyService_ReadCompaniesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadCompanies not implemented")
}
func (*UnimplementedCompanyServiceServer) ReadAllCompanies(ctx context.Context, req *protobuf.Pagination) (*CompanyAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllCompanies not implemented")
}
func (*UnimplementedCompanyServiceServer) UpdateCompany(ctx context.Context, req *Company) (*protobuf.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) DeleteCompany(ctx context.Context, req *protobuf.Id) (*protobuf.Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) Check(ctx context.Context, req *protobuf.HealthCheckRequest) (*protobuf.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedCompanyServiceServer) Watch(req *protobuf.HealthCheckRequest, srv CompanyService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ReadCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ReadCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/ReadCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ReadCompany(ctx, req.(*protobuf.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ReadCompanies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(protobuf.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompanyServiceServer).ReadCompanies(m, &companyServiceReadCompaniesServer{stream})
}

type CompanyService_ReadCompaniesServer interface {
	Send(*Company) error
	grpc.ServerStream
}

type companyServiceReadCompaniesServer struct {
	grpc.ServerStream
}

func (x *companyServiceReadCompaniesServer) Send(m *Company) error {
	return x.ServerStream.SendMsg(m)
}

func _CompanyService_ReadAllCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ReadAllCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/ReadAllCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ReadAllCompanies(ctx, req.(*protobuf.Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Company)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*Company))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*protobuf.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).Check(ctx, req.(*protobuf.HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(protobuf.HealthCheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompanyServiceServer).Watch(m, &companyServiceWatchServer{stream})
}

type CompanyService_WatchServer interface {
	Send(*protobuf.HealthCheckResponse) error
	grpc.ServerStream
}

type companyServiceWatchServer struct {
	grpc.ServerStream
}

func (x *companyServiceWatchServer) Send(m *protobuf.HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "ReadCompany",
			Handler:    _CompanyService_ReadCompany_Handler,
		},
		{
			MethodName: "ReadAllCompanies",
			Handler:    _CompanyService_ReadAllCompanies_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _CompanyService_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadCompanies",
			Handler:       _CompanyService_ReadCompanies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _CompanyService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/company/service.proto",
}
