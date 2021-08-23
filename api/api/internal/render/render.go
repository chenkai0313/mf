package render

import (
	"github.com/gin-gonic/gin"
)

// Response represents API response model.
type Response struct {
	Status string `json:"status"`

	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`

	Data interface{} `json:"data"`
}

// Success returns standard success response json.
func Success(data interface{}) Response {
	if data == nil {
		data = gin.H{}
	}

	return Response{
		Status: "success",
		Data:   data,
	}
}

// Error returns standard error response json.
func Error(errCode int, errMsg string) Response {
	return Response{
		Status:    "error",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
	}
}

// IsFail tells if the response is failed and set to status 'error'.
func (resp Response) IsFail() bool {
	return resp.Status == "error"
}
