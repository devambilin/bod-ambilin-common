package response

import (
	"github.com/devambilin/bod-ambilin-common/constant"
	"github.com/devambilin/bod-ambilin-common/model"
	"github.com/devambilin/bod-ambilin-common/proto"
	"net/http"
)

func SessionExpire() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusUnauthorized,
		Response: ObjectToJson(&model.BaseResponseObject{
			Code:    http.StatusUnauthorized,
			Message: constant.Unauthorized,
		}),
	}
}

func InternalServerError() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusInternalServerError,
		Response: ObjectToJson(&model.BaseResponseObject{
			Code:    http.StatusInternalServerError,
			Message: constant.FailedInternalServer,
		}),
	}
}

func ParameterIsNotEmpty() *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusBadRequest,
		Response: ObjectToJson(&model.BaseResponseObject{
			Code:    http.StatusBadRequest,
			Message: constant.ParameterIsBlank,
		}),
	}
}
