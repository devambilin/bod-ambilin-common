package response

import (
	"encoding/json"
	"github.com/devambilin/bod-ambilin-common/constant"
	"github.com/devambilin/bod-ambilin-common/logger"
	"github.com/devambilin/bod-ambilin-common/model"
	"github.com/devambilin/bod-ambilin-common/proto"
	"net/http"
)

func ObjectToJson(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Infof("error marshal : %v", err.Error())
	}
	return jsonData
}

func SuccessResponseObject(data interface{}) *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusOK,
		Response: ObjectToJson(&model.BaseResponseObject{
			Code:    http.StatusOK,
			Message: constant.Success,
			Data:    data,
		}),
	}
}

func SuccessResponseList(data []interface{}) *proto.BaseResponse {
	return &proto.BaseResponse{
		HttpStatus: http.StatusOK,
		Response: ObjectToJson(model.BaseResponseList{
			Code:    http.StatusOK,
			Message: constant.Success,
			Data:    data,
		}),
	}
}
