package response

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"monorepo/pkg/utils"
)

type ICodeError interface {
	Code() int32
	Message() string
	Error() string
	Stack() string
}

type WrappedResponse struct {
	Code       int32  `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	ServerTime int64  `json:"server_time"`
}

func SendJson(ctx *gin.Context, code int, data any, err ICodeError) {
	resp := &WrappedResponse{
		Data:       data,
		Code:       int32(code),
		Message:    "success",
		ServerTime: time.Now().Unix(),
	}
	if !utils.IsNil(err) {
		resp.Code = err.Code()
		resp.Message = err.Message()
	}
	ctx.JSON(code, resp)
}

func SendProtoBuf(ctx *gin.Context, code int, data proto.Message, err ICodeError) {
	resp := &WrappedResponse{
		Data:       data,
		Code:       int32(code),
		Message:    "success",
		ServerTime: time.Now().Unix(),
	}
	if !utils.IsNil(err) {
		resp.Code = err.Code()
		resp.Message = err.Message()
	}
	ctx.ProtoBuf(code, resp)
}
