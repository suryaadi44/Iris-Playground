package response

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   ErrorValues `json:"error,omitempty"`
}

func NewBaseResponse(message string, data interface{}, errs ErrorValues) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Data:    data,
		Error:   errs,
	}
}

func NewErrorResponse(message string, errs ...ErrorValue) *BaseResponse {
	return &BaseResponse{
		Message: message,
		Error:   NewErrorValues(errs...),
	}
}
