package response

import (
	"github.com/devambilin/bod-ambilin-common/proto"
	"net/http"
)

func SessionExpire() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusUnauthorized,
		Response:   nil,
	}
}

func InternalServerError() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusInternalServerError,
		Response:   nil,
	}
}

func ParameterIsNotEmpty() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusBadRequest,
		Response:   nil,
	}
}
