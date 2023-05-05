// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/dashboard.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DashboardService service

func NewDashboardServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DashboardService service

type DashboardService interface {
	// for client grpc
	GetMasterConfig(ctx context.Context, in *MasterConfigRequest, opts ...client.CallOption) (*MasterConfigResponse, error)
	GetAdmins(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
	GetAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
	CreateAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
	UpdateAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
	DeleteAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
	// page dashboard
	GetSummary(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDataDaily(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDataExecutive(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDataOperational(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetTopAgents(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// combo box
	FindAgents(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	FindOfficeBranches(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	FindOfficeRegionals(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// promo
	GetListPromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDetailPromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CreatePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	UpdatePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeletePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DownloadExecutive(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DownloadOperational(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// CRUD Dashboard
	GetUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDetailUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CreateUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	UpdateUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeleteUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CheckEmail(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CheckPersonalNumber(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	ChangePhoto(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeletePhoto(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// CRUD Role
	GetListMenu(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetListOffice(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetListRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetListRoleStatic(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetDetailRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CreateRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	UpdateRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeleteRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// Activity
	GetHistoryActivity(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type dashboardService struct {
	c    client.Client
	name string
}

func NewDashboardService(name string, c client.Client) DashboardService {
	return &dashboardService{
		c:    c,
		name: name,
	}
}

func (c *dashboardService) GetMasterConfig(ctx context.Context, in *MasterConfigRequest, opts ...client.CallOption) (*MasterConfigResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetMasterConfig", in)
	out := new(MasterConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetAdmins(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetAdmins", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetAdmin", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CreateAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CreateAdmin", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) UpdateAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.UpdateAdmin", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DeleteAdmin(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DeleteAdmin", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetSummary(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetSummary", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDataDaily(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDataDaily", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDataExecutive(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDataExecutive", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDataOperational(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDataOperational", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetTopAgents(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetTopAgents", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) FindAgents(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.FindAgents", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) FindOfficeBranches(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.FindOfficeBranches", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) FindOfficeRegionals(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.FindOfficeRegionals", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetListPromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetListPromo", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDetailPromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDetailPromo", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CreatePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CreatePromo", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) UpdatePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.UpdatePromo", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DeletePromo(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DeletePromo", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DownloadExecutive(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DownloadExecutive", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DownloadOperational(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DownloadOperational", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetUserAdmin", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDetailUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDetailUserAdmin", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CreateUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CreateUserAdmin", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) UpdateUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.UpdateUserAdmin", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DeleteUserAdmin(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DeleteUserAdmin", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CheckEmail(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CheckEmail", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CheckPersonalNumber(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CheckPersonalNumber", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) ChangePhoto(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.ChangePhoto", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DeletePhoto(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DeletePhoto", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetListMenu(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetListMenu", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetListOffice(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetListOffice", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetListRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetListRole", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetListRoleStatic(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetListRoleStatic", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetDetailRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetDetailRole", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) CreateRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.CreateRole", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) UpdateRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.UpdateRole", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) DeleteRole(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.DeleteRole", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardService) GetHistoryActivity(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "DashboardService.GetHistoryActivity", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DashboardService service

type DashboardServiceHandler interface {
	// for client grpc
	GetMasterConfig(context.Context, *MasterConfigRequest, *MasterConfigResponse) error
	GetAdmins(context.Context, *AdminRequest, *AdminResponse) error
	GetAdmin(context.Context, *AdminRequest, *AdminResponse) error
	CreateAdmin(context.Context, *AdminRequest, *AdminResponse) error
	UpdateAdmin(context.Context, *AdminRequest, *AdminResponse) error
	DeleteAdmin(context.Context, *AdminRequest, *AdminResponse) error
	// page dashboard
	GetSummary(context.Context, *BaseRequest, *BaseResponse) error
	GetDataDaily(context.Context, *BaseRequest, *BaseResponse) error
	GetDataExecutive(context.Context, *BaseRequest, *BaseResponse) error
	GetDataOperational(context.Context, *BaseRequest, *BaseResponse) error
	GetTopAgents(context.Context, *BaseRequest, *BaseResponse) error
	// combo box
	FindAgents(context.Context, *BaseRequest, *BaseResponse) error
	FindOfficeBranches(context.Context, *BaseRequest, *BaseResponse) error
	FindOfficeRegionals(context.Context, *BaseRequest, *BaseResponse) error
	// promo
	GetListPromo(context.Context, *BaseRequest, *BaseResponse) error
	GetDetailPromo(context.Context, *BaseRequest, *BaseResponse) error
	CreatePromo(context.Context, *BaseRequest, *BaseResponse) error
	UpdatePromo(context.Context, *BaseRequest, *BaseResponse) error
	DeletePromo(context.Context, *BaseRequest, *BaseResponse) error
	DownloadExecutive(context.Context, *BaseRequest, *BaseResponse) error
	DownloadOperational(context.Context, *BaseRequest, *BaseResponse) error
	// CRUD Dashboard
	GetUserAdmin(context.Context, *BaseRequest, *BaseResponse) error
	GetDetailUserAdmin(context.Context, *BaseRequest, *BaseResponse) error
	CreateUserAdmin(context.Context, *BaseRequest, *BaseResponse) error
	UpdateUserAdmin(context.Context, *BaseRequest, *BaseResponse) error
	DeleteUserAdmin(context.Context, *BaseRequest, *BaseResponse) error
	CheckEmail(context.Context, *BaseRequest, *BaseResponse) error
	CheckPersonalNumber(context.Context, *BaseRequest, *BaseResponse) error
	ChangePhoto(context.Context, *BaseRequest, *BaseResponse) error
	DeletePhoto(context.Context, *BaseRequest, *BaseResponse) error
	// CRUD Role
	GetListMenu(context.Context, *BaseRequest, *BaseResponse) error
	GetListOffice(context.Context, *BaseRequest, *BaseResponse) error
	GetListRole(context.Context, *BaseRequest, *BaseResponse) error
	GetListRoleStatic(context.Context, *BaseRequest, *BaseResponse) error
	GetDetailRole(context.Context, *BaseRequest, *BaseResponse) error
	CreateRole(context.Context, *BaseRequest, *BaseResponse) error
	UpdateRole(context.Context, *BaseRequest, *BaseResponse) error
	DeleteRole(context.Context, *BaseRequest, *BaseResponse) error
	// Activity
	GetHistoryActivity(context.Context, *BaseRequest, *BaseResponse) error
}

func RegisterDashboardServiceHandler(s server.Server, hdlr DashboardServiceHandler, opts ...server.HandlerOption) error {
	type dashboardService interface {
		GetMasterConfig(ctx context.Context, in *MasterConfigRequest, out *MasterConfigResponse) error
		GetAdmins(ctx context.Context, in *AdminRequest, out *AdminResponse) error
		GetAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error
		CreateAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error
		UpdateAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error
		DeleteAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error
		GetSummary(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDataDaily(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDataExecutive(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDataOperational(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetTopAgents(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		FindAgents(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		FindOfficeBranches(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		FindOfficeRegionals(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetListPromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDetailPromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CreatePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		UpdatePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeletePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DownloadExecutive(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DownloadOperational(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDetailUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CreateUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		UpdateUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeleteUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CheckEmail(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CheckPersonalNumber(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		ChangePhoto(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeletePhoto(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetListMenu(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetListOffice(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetListRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetListRoleStatic(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetDetailRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CreateRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		UpdateRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeleteRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetHistoryActivity(ctx context.Context, in *BaseRequest, out *BaseResponse) error
	}
	type DashboardService struct {
		dashboardService
	}
	h := &dashboardServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DashboardService{h}, opts...))
}

type dashboardServiceHandler struct {
	DashboardServiceHandler
}

func (h *dashboardServiceHandler) GetMasterConfig(ctx context.Context, in *MasterConfigRequest, out *MasterConfigResponse) error {
	return h.DashboardServiceHandler.GetMasterConfig(ctx, in, out)
}

func (h *dashboardServiceHandler) GetAdmins(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.DashboardServiceHandler.GetAdmins(ctx, in, out)
}

func (h *dashboardServiceHandler) GetAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.DashboardServiceHandler.GetAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) CreateAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.DashboardServiceHandler.CreateAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) UpdateAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.DashboardServiceHandler.UpdateAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) DeleteAdmin(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.DashboardServiceHandler.DeleteAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) GetSummary(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetSummary(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDataDaily(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDataDaily(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDataExecutive(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDataExecutive(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDataOperational(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDataOperational(ctx, in, out)
}

func (h *dashboardServiceHandler) GetTopAgents(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetTopAgents(ctx, in, out)
}

func (h *dashboardServiceHandler) FindAgents(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.FindAgents(ctx, in, out)
}

func (h *dashboardServiceHandler) FindOfficeBranches(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.FindOfficeBranches(ctx, in, out)
}

func (h *dashboardServiceHandler) FindOfficeRegionals(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.FindOfficeRegionals(ctx, in, out)
}

func (h *dashboardServiceHandler) GetListPromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetListPromo(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDetailPromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDetailPromo(ctx, in, out)
}

func (h *dashboardServiceHandler) CreatePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.CreatePromo(ctx, in, out)
}

func (h *dashboardServiceHandler) UpdatePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.UpdatePromo(ctx, in, out)
}

func (h *dashboardServiceHandler) DeletePromo(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DeletePromo(ctx, in, out)
}

func (h *dashboardServiceHandler) DownloadExecutive(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DownloadExecutive(ctx, in, out)
}

func (h *dashboardServiceHandler) DownloadOperational(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DownloadOperational(ctx, in, out)
}

func (h *dashboardServiceHandler) GetUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetUserAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDetailUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDetailUserAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) CreateUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.CreateUserAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) UpdateUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.UpdateUserAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) DeleteUserAdmin(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DeleteUserAdmin(ctx, in, out)
}

func (h *dashboardServiceHandler) CheckEmail(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.CheckEmail(ctx, in, out)
}

func (h *dashboardServiceHandler) CheckPersonalNumber(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.CheckPersonalNumber(ctx, in, out)
}

func (h *dashboardServiceHandler) ChangePhoto(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.ChangePhoto(ctx, in, out)
}

func (h *dashboardServiceHandler) DeletePhoto(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DeletePhoto(ctx, in, out)
}

func (h *dashboardServiceHandler) GetListMenu(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetListMenu(ctx, in, out)
}

func (h *dashboardServiceHandler) GetListOffice(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetListOffice(ctx, in, out)
}

func (h *dashboardServiceHandler) GetListRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetListRole(ctx, in, out)
}

func (h *dashboardServiceHandler) GetListRoleStatic(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetListRoleStatic(ctx, in, out)
}

func (h *dashboardServiceHandler) GetDetailRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetDetailRole(ctx, in, out)
}

func (h *dashboardServiceHandler) CreateRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.CreateRole(ctx, in, out)
}

func (h *dashboardServiceHandler) UpdateRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.UpdateRole(ctx, in, out)
}

func (h *dashboardServiceHandler) DeleteRole(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.DeleteRole(ctx, in, out)
}

func (h *dashboardServiceHandler) GetHistoryActivity(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.DashboardServiceHandler.GetHistoryActivity(ctx, in, out)
}
