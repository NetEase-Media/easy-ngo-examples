package vo

type BaseResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess() *BaseResponse {
	return &BaseResponse{
		Code:    200,
		Message: "success",
	}
}

func ResponseSuccessWithData(d interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    200,
		Data:    d,
		Message: "success",
	}
}

func ResponseFailed() *BaseResponse {
	return &BaseResponse{
		Code:    400,
		Message: "failed",
	}
}

func ResponseFailedWithCode(c int32) *BaseResponse {
	return &BaseResponse{
		Code:    c,
		Message: "failed",
	}
}
