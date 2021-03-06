// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: company.proto

package protobufpb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	// Add a new owner for a company.
	AddCEO(ctx context.Context, in *CompanyAddCEOReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// Add a new company.
	AddCompany(ctx context.Context, in *CompanyAddReq, opts ...grpc.CallOption) (*CompanyAddRes, error)
	// Update a company profile.
	UpdateCompany(ctx context.Context, in *CompanyUpdateReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get a company profile by its code.
	GetCompanyByCode(ctx context.Context, in *GetCompanyByCodeReq, opts ...grpc.CallOption) (*GetCompanyByCodeRes, error)
	// List all companies.
	ListCompany(ctx context.Context, in *CompanyListCompanyReq, opts ...grpc.CallOption) (*CompanyListCompanyRes, error)
	// Get my profile
	GetMyProfile(ctx context.Context, in *CompanyGetMyProfileReq, opts ...grpc.CallOption) (*CompanyGetMyProfileRes, error)
	// Get a company profile based on user's token.
	GetMyCompany(ctx context.Context, in *CompanyGetMyCompanyReq, opts ...grpc.CallOption) (*CompanyGetMyCompanyRes, error)
	// Add a new department.
	AddDepartment(ctx context.Context, in *CompanyAddDepartmentReq, opts ...grpc.CallOption) (*CompanyAddDepartmentRes, error)
	// Add a new group.
	AddGroup(ctx context.Context, in *CompanyAddGroupReq, opts ...grpc.CallOption) (*CompanyAddGroupRes, error)
	// List all departments of a company.
	ListDepartment(ctx context.Context, in *CompanyListDepartmentReq, opts ...grpc.CallOption) (*CompanyListDepartmentRes, error)
	// List all groups of a department.
	ListGroup(ctx context.Context, in *CompanyListGroupReq, opts ...grpc.CallOption) (*CompanyListGroupRes, error)
	// List all employees of a company.
	ListEmployee(ctx context.Context, in *CompanyListEmployeeReq, opts ...grpc.CallOption) (*CompanyListEmployeeRes, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) AddCEO(ctx context.Context, in *CompanyAddCEOReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/AddCEO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) AddCompany(ctx context.Context, in *CompanyAddReq, opts ...grpc.CallOption) (*CompanyAddRes, error) {
	out := new(CompanyAddRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/AddCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *CompanyUpdateReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetCompanyByCode(ctx context.Context, in *GetCompanyByCodeReq, opts ...grpc.CallOption) (*GetCompanyByCodeRes, error) {
	out := new(GetCompanyByCodeRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/GetCompanyByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListCompany(ctx context.Context, in *CompanyListCompanyReq, opts ...grpc.CallOption) (*CompanyListCompanyRes, error) {
	out := new(CompanyListCompanyRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/ListCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetMyProfile(ctx context.Context, in *CompanyGetMyProfileReq, opts ...grpc.CallOption) (*CompanyGetMyProfileRes, error) {
	out := new(CompanyGetMyProfileRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/GetMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetMyCompany(ctx context.Context, in *CompanyGetMyCompanyReq, opts ...grpc.CallOption) (*CompanyGetMyCompanyRes, error) {
	out := new(CompanyGetMyCompanyRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/GetMyCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) AddDepartment(ctx context.Context, in *CompanyAddDepartmentReq, opts ...grpc.CallOption) (*CompanyAddDepartmentRes, error) {
	out := new(CompanyAddDepartmentRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/AddDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) AddGroup(ctx context.Context, in *CompanyAddGroupReq, opts ...grpc.CallOption) (*CompanyAddGroupRes, error) {
	out := new(CompanyAddGroupRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/AddGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListDepartment(ctx context.Context, in *CompanyListDepartmentReq, opts ...grpc.CallOption) (*CompanyListDepartmentRes, error) {
	out := new(CompanyListDepartmentRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/ListDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListGroup(ctx context.Context, in *CompanyListGroupReq, opts ...grpc.CallOption) (*CompanyListGroupRes, error) {
	out := new(CompanyListGroupRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/ListGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListEmployee(ctx context.Context, in *CompanyListEmployeeReq, opts ...grpc.CallOption) (*CompanyListEmployeeRes, error) {
	out := new(CompanyListEmployeeRes)
	err := c.cc.Invoke(ctx, "/protobuf.CompanyService/ListEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility
type CompanyServiceServer interface {
	// Add a new owner for a company.
	AddCEO(context.Context, *CompanyAddCEOReq) (*empty.Empty, error)
	// Add a new company.
	AddCompany(context.Context, *CompanyAddReq) (*CompanyAddRes, error)
	// Update a company profile.
	UpdateCompany(context.Context, *CompanyUpdateReq) (*empty.Empty, error)
	// Get a company profile by its code.
	GetCompanyByCode(context.Context, *GetCompanyByCodeReq) (*GetCompanyByCodeRes, error)
	// List all companies.
	ListCompany(context.Context, *CompanyListCompanyReq) (*CompanyListCompanyRes, error)
	// Get my profile
	GetMyProfile(context.Context, *CompanyGetMyProfileReq) (*CompanyGetMyProfileRes, error)
	// Get a company profile based on user's token.
	GetMyCompany(context.Context, *CompanyGetMyCompanyReq) (*CompanyGetMyCompanyRes, error)
	// Add a new department.
	AddDepartment(context.Context, *CompanyAddDepartmentReq) (*CompanyAddDepartmentRes, error)
	// Add a new group.
	AddGroup(context.Context, *CompanyAddGroupReq) (*CompanyAddGroupRes, error)
	// List all departments of a company.
	ListDepartment(context.Context, *CompanyListDepartmentReq) (*CompanyListDepartmentRes, error)
	// List all groups of a department.
	ListGroup(context.Context, *CompanyListGroupReq) (*CompanyListGroupRes, error)
	// List all employees of a company.
	ListEmployee(context.Context, *CompanyListEmployeeReq) (*CompanyListEmployeeRes, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (UnimplementedCompanyServiceServer) AddCEO(context.Context, *CompanyAddCEOReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCEO not implemented")
}
func (UnimplementedCompanyServiceServer) AddCompany(context.Context, *CompanyAddReq) (*CompanyAddRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCompany not implemented")
}
func (UnimplementedCompanyServiceServer) UpdateCompany(context.Context, *CompanyUpdateReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) GetCompanyByCode(context.Context, *GetCompanyByCodeReq) (*GetCompanyByCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyByCode not implemented")
}
func (UnimplementedCompanyServiceServer) ListCompany(context.Context, *CompanyListCompanyReq) (*CompanyListCompanyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompany not implemented")
}
func (UnimplementedCompanyServiceServer) GetMyProfile(context.Context, *CompanyGetMyProfileReq) (*CompanyGetMyProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (UnimplementedCompanyServiceServer) GetMyCompany(context.Context, *CompanyGetMyCompanyReq) (*CompanyGetMyCompanyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCompany not implemented")
}
func (UnimplementedCompanyServiceServer) AddDepartment(context.Context, *CompanyAddDepartmentReq) (*CompanyAddDepartmentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDepartment not implemented")
}
func (UnimplementedCompanyServiceServer) AddGroup(context.Context, *CompanyAddGroupReq) (*CompanyAddGroupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedCompanyServiceServer) ListDepartment(context.Context, *CompanyListDepartmentReq) (*CompanyListDepartmentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartment not implemented")
}
func (UnimplementedCompanyServiceServer) ListGroup(context.Context, *CompanyListGroupReq) (*CompanyListGroupRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroup not implemented")
}
func (UnimplementedCompanyServiceServer) ListEmployee(context.Context, *CompanyListEmployeeReq) (*CompanyListEmployeeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEmployee not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_AddCEO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyAddCEOReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddCEO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/AddCEO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddCEO(ctx, req.(*CompanyAddCEOReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_AddCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/AddCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddCompany(ctx, req.(*CompanyAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*CompanyUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetCompanyByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyByCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompanyByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/GetCompanyByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompanyByCode(ctx, req.(*GetCompanyByCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/ListCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListCompany(ctx, req.(*CompanyListCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyGetMyProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/GetMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetMyProfile(ctx, req.(*CompanyGetMyProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetMyCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyGetMyCompanyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetMyCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/GetMyCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetMyCompany(ctx, req.(*CompanyGetMyCompanyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_AddDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyAddDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/AddDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddDepartment(ctx, req.(*CompanyAddDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyAddGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/AddGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).AddGroup(ctx, req.(*CompanyAddGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListDepartmentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/ListDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListDepartment(ctx, req.(*CompanyListDepartmentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/ListGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListGroup(ctx, req.(*CompanyListGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListEmployeeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.CompanyService/ListEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListEmployee(ctx, req.(*CompanyListEmployeeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCEO",
			Handler:    _CompanyService_AddCEO_Handler,
		},
		{
			MethodName: "AddCompany",
			Handler:    _CompanyService_AddCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "GetCompanyByCode",
			Handler:    _CompanyService_GetCompanyByCode_Handler,
		},
		{
			MethodName: "ListCompany",
			Handler:    _CompanyService_ListCompany_Handler,
		},
		{
			MethodName: "GetMyProfile",
			Handler:    _CompanyService_GetMyProfile_Handler,
		},
		{
			MethodName: "GetMyCompany",
			Handler:    _CompanyService_GetMyCompany_Handler,
		},
		{
			MethodName: "AddDepartment",
			Handler:    _CompanyService_AddDepartment_Handler,
		},
		{
			MethodName: "AddGroup",
			Handler:    _CompanyService_AddGroup_Handler,
		},
		{
			MethodName: "ListDepartment",
			Handler:    _CompanyService_ListDepartment_Handler,
		},
		{
			MethodName: "ListGroup",
			Handler:    _CompanyService_ListGroup_Handler,
		},
		{
			MethodName: "ListEmployee",
			Handler:    _CompanyService_ListEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "company.proto",
}
