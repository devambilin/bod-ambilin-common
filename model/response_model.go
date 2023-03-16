package model

type BaseResponseObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type BaseResponseList struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data,omitempty"`
}
