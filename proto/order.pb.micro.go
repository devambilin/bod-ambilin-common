// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OrderService service

func NewOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderService service

type OrderService interface {
	// for clients grpc
	GetOrders(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	GetOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	UpdateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	DeleteOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error)
	// for customer
	CustomerSubmitOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerDetailOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerConfirmOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerGetListOrderProcess(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerChangeStatus(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerCancelOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerCheckOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// for agent
	AgentGetListProcessAndHistoryOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentAcceptOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentPickupOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentGoingOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentConfirmCodeOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentUpdateAmountOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentGetFeeTransactionOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) GetOrders(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrders", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) GetOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.GetOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CreateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CreateOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) UpdateOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.UpdateOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) DeleteOrder(ctx context.Context, in *OrderRequest, opts ...client.CallOption) (*OrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.DeleteOrder", in)
	out := new(OrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerSubmitOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerSubmitOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerDetailOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerDetailOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerConfirmOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerConfirmOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerGetListOrderProcess(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerGetListOrderProcess", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerChangeStatus(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerChangeStatus", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerCancelOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerCancelOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) CustomerCheckOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.CustomerCheckOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentGetListProcessAndHistoryOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentGetListProcessAndHistoryOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentAcceptOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentAcceptOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentPickupOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentPickupOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentGoingOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentGoingOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentConfirmCodeOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentConfirmCodeOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentUpdateAmountOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentUpdateAmountOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) AgentGetFeeTransactionOrder(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "OrderService.AgentGetFeeTransactionOrder", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceHandler interface {
	// for clients grpc
	GetOrders(context.Context, *OrderRequest, *OrderResponse) error
	GetOrder(context.Context, *OrderRequest, *OrderResponse) error
	CreateOrder(context.Context, *OrderRequest, *OrderResponse) error
	UpdateOrder(context.Context, *OrderRequest, *OrderResponse) error
	DeleteOrder(context.Context, *OrderRequest, *OrderResponse) error
	// for customer
	CustomerSubmitOrder(context.Context, *BaseRequest, *BaseResponse) error
	CustomerDetailOrder(context.Context, *BaseRequest, *BaseResponse) error
	CustomerConfirmOrder(context.Context, *BaseRequest, *BaseResponse) error
	CustomerGetListOrderProcess(context.Context, *BaseRequest, *BaseResponse) error
	CustomerChangeStatus(context.Context, *BaseRequest, *BaseResponse) error
	CustomerCancelOrder(context.Context, *BaseRequest, *BaseResponse) error
	CustomerCheckOrder(context.Context, *BaseRequest, *BaseResponse) error
	// for agent
	AgentGetListProcessAndHistoryOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentAcceptOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentPickupOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentGoingOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentConfirmCodeOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentUpdateAmountOrder(context.Context, *BaseRequest, *BaseResponse) error
	AgentGetFeeTransactionOrder(context.Context, *BaseRequest, *BaseResponse) error
}

func RegisterOrderServiceHandler(s server.Server, hdlr OrderServiceHandler, opts ...server.HandlerOption) error {
	type orderService interface {
		GetOrders(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		GetOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		CreateOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		UpdateOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		DeleteOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error
		CustomerSubmitOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerDetailOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerConfirmOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerGetListOrderProcess(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerChangeStatus(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerCancelOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerCheckOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentGetListProcessAndHistoryOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentAcceptOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentPickupOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentGoingOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentConfirmCodeOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentUpdateAmountOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentGetFeeTransactionOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error
	}
	type OrderService struct {
		orderService
	}
	h := &orderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderService{h}, opts...))
}

type orderServiceHandler struct {
	OrderServiceHandler
}

func (h *orderServiceHandler) GetOrders(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.GetOrders(ctx, in, out)
}

func (h *orderServiceHandler) GetOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.GetOrder(ctx, in, out)
}

func (h *orderServiceHandler) CreateOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.CreateOrder(ctx, in, out)
}

func (h *orderServiceHandler) UpdateOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.UpdateOrder(ctx, in, out)
}

func (h *orderServiceHandler) DeleteOrder(ctx context.Context, in *OrderRequest, out *OrderResponse) error {
	return h.OrderServiceHandler.DeleteOrder(ctx, in, out)
}

func (h *orderServiceHandler) CustomerSubmitOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerSubmitOrder(ctx, in, out)
}

func (h *orderServiceHandler) CustomerDetailOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerDetailOrder(ctx, in, out)
}

func (h *orderServiceHandler) CustomerConfirmOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerConfirmOrder(ctx, in, out)
}

func (h *orderServiceHandler) CustomerGetListOrderProcess(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerGetListOrderProcess(ctx, in, out)
}

func (h *orderServiceHandler) CustomerChangeStatus(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerChangeStatus(ctx, in, out)
}

func (h *orderServiceHandler) CustomerCancelOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerCancelOrder(ctx, in, out)
}

func (h *orderServiceHandler) CustomerCheckOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.CustomerCheckOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentGetListProcessAndHistoryOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentGetListProcessAndHistoryOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentAcceptOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentAcceptOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentPickupOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentPickupOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentGoingOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentGoingOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentConfirmCodeOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentConfirmCodeOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentUpdateAmountOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentUpdateAmountOrder(ctx, in, out)
}

func (h *orderServiceHandler) AgentGetFeeTransactionOrder(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.OrderServiceHandler.AgentGetFeeTransactionOrder(ctx, in, out)
}
