// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: organization/v1/organization.proto

package organizationv1connect

import (
	v1 "app/gen/organization/v1"
	v11 "app/gen/tools/v1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// OrganizationName is the fully-qualified name of the Organization service.
	OrganizationName = "organization.v1.Organization"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrganizationCreateOrganizationProcedure is the fully-qualified name of the Organization's
	// CreateOrganization RPC.
	OrganizationCreateOrganizationProcedure = "/organization.v1.Organization/CreateOrganization"
	// OrganizationDetailOrganizationProcedure is the fully-qualified name of the Organization's
	// DetailOrganization RPC.
	OrganizationDetailOrganizationProcedure = "/organization.v1.Organization/DetailOrganization"
	// OrganizationListOrganizationProcedure is the fully-qualified name of the Organization's
	// ListOrganization RPC.
	OrganizationListOrganizationProcedure = "/organization.v1.Organization/ListOrganization"
	// OrganizationUpdateOrganizationProcedure is the fully-qualified name of the Organization's
	// UpdateOrganization RPC.
	OrganizationUpdateOrganizationProcedure = "/organization.v1.Organization/UpdateOrganization"
	// OrganizationDeleteOrganizationProcedure is the fully-qualified name of the Organization's
	// DeleteOrganization RPC.
	OrganizationDeleteOrganizationProcedure = "/organization.v1.Organization/DeleteOrganization"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	organizationServiceDescriptor                  = v1.File_organization_v1_organization_proto.Services().ByName("Organization")
	organizationCreateOrganizationMethodDescriptor = organizationServiceDescriptor.Methods().ByName("CreateOrganization")
	organizationDetailOrganizationMethodDescriptor = organizationServiceDescriptor.Methods().ByName("DetailOrganization")
	organizationListOrganizationMethodDescriptor   = organizationServiceDescriptor.Methods().ByName("ListOrganization")
	organizationUpdateOrganizationMethodDescriptor = organizationServiceDescriptor.Methods().ByName("UpdateOrganization")
	organizationDeleteOrganizationMethodDescriptor = organizationServiceDescriptor.Methods().ByName("DeleteOrganization")
)

// OrganizationClient is a client for the organization.v1.Organization service.
type OrganizationClient interface {
	CreateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	DetailOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	ListOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganizationList], error)
	UpdateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	DeleteOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v11.Empty], error)
}

// NewOrganizationClient constructs a client for the organization.v1.Organization service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrganizationClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrganizationClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &organizationClient{
		createOrganization: connect.NewClient[v1.RequestOrganization, v1.ResponseOrganization](
			httpClient,
			baseURL+OrganizationCreateOrganizationProcedure,
			connect.WithSchema(organizationCreateOrganizationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		detailOrganization: connect.NewClient[v1.ParamsOrganization, v1.ResponseOrganization](
			httpClient,
			baseURL+OrganizationDetailOrganizationProcedure,
			connect.WithSchema(organizationDetailOrganizationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listOrganization: connect.NewClient[v1.ParamsOrganization, v1.ResponseOrganizationList](
			httpClient,
			baseURL+OrganizationListOrganizationProcedure,
			connect.WithSchema(organizationListOrganizationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateOrganization: connect.NewClient[v1.RequestOrganization, v1.ResponseOrganization](
			httpClient,
			baseURL+OrganizationUpdateOrganizationProcedure,
			connect.WithSchema(organizationUpdateOrganizationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteOrganization: connect.NewClient[v1.ParamsOrganization, v11.Empty](
			httpClient,
			baseURL+OrganizationDeleteOrganizationProcedure,
			connect.WithSchema(organizationDeleteOrganizationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// organizationClient implements OrganizationClient.
type organizationClient struct {
	createOrganization *connect.Client[v1.RequestOrganization, v1.ResponseOrganization]
	detailOrganization *connect.Client[v1.ParamsOrganization, v1.ResponseOrganization]
	listOrganization   *connect.Client[v1.ParamsOrganization, v1.ResponseOrganizationList]
	updateOrganization *connect.Client[v1.RequestOrganization, v1.ResponseOrganization]
	deleteOrganization *connect.Client[v1.ParamsOrganization, v11.Empty]
}

// CreateOrganization calls organization.v1.Organization.CreateOrganization.
func (c *organizationClient) CreateOrganization(ctx context.Context, req *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return c.createOrganization.CallUnary(ctx, req)
}

// DetailOrganization calls organization.v1.Organization.DetailOrganization.
func (c *organizationClient) DetailOrganization(ctx context.Context, req *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return c.detailOrganization.CallUnary(ctx, req)
}

// ListOrganization calls organization.v1.Organization.ListOrganization.
func (c *organizationClient) ListOrganization(ctx context.Context, req *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganizationList], error) {
	return c.listOrganization.CallUnary(ctx, req)
}

// UpdateOrganization calls organization.v1.Organization.UpdateOrganization.
func (c *organizationClient) UpdateOrganization(ctx context.Context, req *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return c.updateOrganization.CallUnary(ctx, req)
}

// DeleteOrganization calls organization.v1.Organization.DeleteOrganization.
func (c *organizationClient) DeleteOrganization(ctx context.Context, req *connect.Request[v1.ParamsOrganization]) (*connect.Response[v11.Empty], error) {
	return c.deleteOrganization.CallUnary(ctx, req)
}

// OrganizationHandler is an implementation of the organization.v1.Organization service.
type OrganizationHandler interface {
	CreateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	DetailOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	ListOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganizationList], error)
	UpdateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error)
	DeleteOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v11.Empty], error)
}

// NewOrganizationHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrganizationHandler(svc OrganizationHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	organizationCreateOrganizationHandler := connect.NewUnaryHandler(
		OrganizationCreateOrganizationProcedure,
		svc.CreateOrganization,
		connect.WithSchema(organizationCreateOrganizationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationDetailOrganizationHandler := connect.NewUnaryHandler(
		OrganizationDetailOrganizationProcedure,
		svc.DetailOrganization,
		connect.WithSchema(organizationDetailOrganizationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationListOrganizationHandler := connect.NewUnaryHandler(
		OrganizationListOrganizationProcedure,
		svc.ListOrganization,
		connect.WithSchema(organizationListOrganizationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationUpdateOrganizationHandler := connect.NewUnaryHandler(
		OrganizationUpdateOrganizationProcedure,
		svc.UpdateOrganization,
		connect.WithSchema(organizationUpdateOrganizationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	organizationDeleteOrganizationHandler := connect.NewUnaryHandler(
		OrganizationDeleteOrganizationProcedure,
		svc.DeleteOrganization,
		connect.WithSchema(organizationDeleteOrganizationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/organization.v1.Organization/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrganizationCreateOrganizationProcedure:
			organizationCreateOrganizationHandler.ServeHTTP(w, r)
		case OrganizationDetailOrganizationProcedure:
			organizationDetailOrganizationHandler.ServeHTTP(w, r)
		case OrganizationListOrganizationProcedure:
			organizationListOrganizationHandler.ServeHTTP(w, r)
		case OrganizationUpdateOrganizationProcedure:
			organizationUpdateOrganizationHandler.ServeHTTP(w, r)
		case OrganizationDeleteOrganizationProcedure:
			organizationDeleteOrganizationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrganizationHandler returns CodeUnimplemented from all methods.
type UnimplementedOrganizationHandler struct{}

func (UnimplementedOrganizationHandler) CreateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("organization.v1.Organization.CreateOrganization is not implemented"))
}

func (UnimplementedOrganizationHandler) DetailOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("organization.v1.Organization.DetailOrganization is not implemented"))
}

func (UnimplementedOrganizationHandler) ListOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v1.ResponseOrganizationList], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("organization.v1.Organization.ListOrganization is not implemented"))
}

func (UnimplementedOrganizationHandler) UpdateOrganization(context.Context, *connect.Request[v1.RequestOrganization]) (*connect.Response[v1.ResponseOrganization], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("organization.v1.Organization.UpdateOrganization is not implemented"))
}

func (UnimplementedOrganizationHandler) DeleteOrganization(context.Context, *connect.Request[v1.ParamsOrganization]) (*connect.Response[v11.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("organization.v1.Organization.DeleteOrganization is not implemented"))
}
