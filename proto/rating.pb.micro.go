// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/rating.proto

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

// Api Endpoints for RatingService service

func NewRatingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RatingService service

type RatingService interface {
	// for clients grpc
	CreateAgentRating(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error)
	CreateCustomerRating(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error)
	GetAgentRatings(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error)
	GetAgentRatingTotal(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error)
	GetCustomerRatings(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error)
	GetCustomerRatingTotal(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error)
	// for customer
	CustomerCreateAgentRating(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerGetRatings(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CustomerGetRatingTotal(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// for agent
	AgentCreateCustomerRating(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentGetRatings(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AgentGetRatingTotal(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type ratingService struct {
	c    client.Client
	name string
}

func NewRatingService(name string, c client.Client) RatingService {
	return &ratingService{
		c:    c,
		name: name,
	}
}

func (c *ratingService) CreateAgentRating(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.CreateAgentRating", in)
	out := new(AgentRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) CreateCustomerRating(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.CreateCustomerRating", in)
	out := new(UserRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) GetAgentRatings(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.GetAgentRatings", in)
	out := new(AgentRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) GetAgentRatingTotal(ctx context.Context, in *AgentRatingRequest, opts ...client.CallOption) (*AgentRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.GetAgentRatingTotal", in)
	out := new(AgentRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) GetCustomerRatings(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.GetCustomerRatings", in)
	out := new(UserRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) GetCustomerRatingTotal(ctx context.Context, in *UserRatingRequest, opts ...client.CallOption) (*UserRatingResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.GetCustomerRatingTotal", in)
	out := new(UserRatingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) CustomerCreateAgentRating(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.CustomerCreateAgentRating", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) CustomerGetRatings(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.CustomerGetRatings", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) CustomerGetRatingTotal(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.CustomerGetRatingTotal", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) AgentCreateCustomerRating(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.AgentCreateCustomerRating", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) AgentGetRatings(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.AgentGetRatings", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingService) AgentGetRatingTotal(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "RatingService.AgentGetRatingTotal", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RatingService service

type RatingServiceHandler interface {
	// for clients grpc
	CreateAgentRating(context.Context, *AgentRatingRequest, *AgentRatingResponse) error
	CreateCustomerRating(context.Context, *UserRatingRequest, *UserRatingResponse) error
	GetAgentRatings(context.Context, *AgentRatingRequest, *AgentRatingResponse) error
	GetAgentRatingTotal(context.Context, *AgentRatingRequest, *AgentRatingResponse) error
	GetCustomerRatings(context.Context, *UserRatingRequest, *UserRatingResponse) error
	GetCustomerRatingTotal(context.Context, *UserRatingRequest, *UserRatingResponse) error
	// for customer
	CustomerCreateAgentRating(context.Context, *BaseRequest, *BaseResponse) error
	CustomerGetRatings(context.Context, *BaseRequest, *BaseResponse) error
	CustomerGetRatingTotal(context.Context, *BaseRequest, *BaseResponse) error
	// for agent
	AgentCreateCustomerRating(context.Context, *BaseRequest, *BaseResponse) error
	AgentGetRatings(context.Context, *BaseRequest, *BaseResponse) error
	AgentGetRatingTotal(context.Context, *BaseRequest, *BaseResponse) error
}

func RegisterRatingServiceHandler(s server.Server, hdlr RatingServiceHandler, opts ...server.HandlerOption) error {
	type ratingService interface {
		CreateAgentRating(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error
		CreateCustomerRating(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error
		GetAgentRatings(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error
		GetAgentRatingTotal(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error
		GetCustomerRatings(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error
		GetCustomerRatingTotal(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error
		CustomerCreateAgentRating(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerGetRatings(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CustomerGetRatingTotal(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentCreateCustomerRating(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentGetRatings(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AgentGetRatingTotal(ctx context.Context, in *BaseRequest, out *BaseResponse) error
	}
	type RatingService struct {
		ratingService
	}
	h := &ratingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RatingService{h}, opts...))
}

type ratingServiceHandler struct {
	RatingServiceHandler
}

func (h *ratingServiceHandler) CreateAgentRating(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error {
	return h.RatingServiceHandler.CreateAgentRating(ctx, in, out)
}

func (h *ratingServiceHandler) CreateCustomerRating(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error {
	return h.RatingServiceHandler.CreateCustomerRating(ctx, in, out)
}

func (h *ratingServiceHandler) GetAgentRatings(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error {
	return h.RatingServiceHandler.GetAgentRatings(ctx, in, out)
}

func (h *ratingServiceHandler) GetAgentRatingTotal(ctx context.Context, in *AgentRatingRequest, out *AgentRatingResponse) error {
	return h.RatingServiceHandler.GetAgentRatingTotal(ctx, in, out)
}

func (h *ratingServiceHandler) GetCustomerRatings(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error {
	return h.RatingServiceHandler.GetCustomerRatings(ctx, in, out)
}

func (h *ratingServiceHandler) GetCustomerRatingTotal(ctx context.Context, in *UserRatingRequest, out *UserRatingResponse) error {
	return h.RatingServiceHandler.GetCustomerRatingTotal(ctx, in, out)
}

func (h *ratingServiceHandler) CustomerCreateAgentRating(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.CustomerCreateAgentRating(ctx, in, out)
}

func (h *ratingServiceHandler) CustomerGetRatings(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.CustomerGetRatings(ctx, in, out)
}

func (h *ratingServiceHandler) CustomerGetRatingTotal(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.CustomerGetRatingTotal(ctx, in, out)
}

func (h *ratingServiceHandler) AgentCreateCustomerRating(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.AgentCreateCustomerRating(ctx, in, out)
}

func (h *ratingServiceHandler) AgentGetRatings(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.AgentGetRatings(ctx, in, out)
}

func (h *ratingServiceHandler) AgentGetRatingTotal(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.RatingServiceHandler.AgentGetRatingTotal(ctx, in, out)
}
