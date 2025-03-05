package request

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
)

type BaseRequest[P any, H any] struct {
	ctx *gin.Context

	headers *H
	params  *P
}

func NewBaseRequest[P any, H any](context *gin.Context) *BaseRequest[P, H] {
	r := &BaseRequest[P, H]{
		ctx: context,
	}
	return r
}

func ParseParams[P any, H any](context *gin.Context) (*BaseRequest[P, H], error) {
	req := NewBaseRequest[P, H](context)
	err := req.bindHeaders()
	if err != nil {
		return nil, errors.Wrap(err, "request.ParseParams.bindHeaders failed")
	}
	err = req.bindParams()
	if err != nil {
		return nil, errors.Wrap(err, "request.ParseParams.bindParams failed")
	}
	return req, nil
}

func (r *BaseRequest[P, H]) bindParams() error {
	params := new(P)

	// 绑定body 中json参数
	method := strings.ToLower(r.ctx.Request.Method)
	if method == "post" {
		err := r.ctx.ShouldBindBodyWith(params, binding.JSON)
		if err != nil {
			return errors.Wrap(err, "bind json params failed")
		}
	}

	// 绑定url中query参数，会覆盖body中解出来的参数
	err := r.ctx.ShouldBindQuery(params)
	if err != nil {
		return errors.Wrap(err, "bind query params failed")
	}

	r.params = params
	return nil
}

func (r *BaseRequest[P, H]) bindHeaders() error {
	headers := new(H)
	err := r.ctx.ShouldBindHeader(headers)
	if err != nil {
		return errors.Wrap(err, "bind request headers failed")
	}
	r.headers = headers
	return nil
}

func (r *BaseRequest[P, H]) GetParams() *P {
	return r.params
}

func (r *BaseRequest[T, H]) GetHeaders() *H {
	return r.headers
}

type Request[P any] struct {
	*BaseRequest[P, Headers]
}

func NewRequest[P any](ctx *gin.Context) (*Request[P], error) {
	req, err := ParseParams[P, Headers](ctx)
	newReq := &Request[P]{
		BaseRequest: req,
	}
	return newReq, err
}
