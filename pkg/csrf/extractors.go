package csrf

import (
	"github.com/gin-gonic/gin"
)

// CsrfFromParam returns a function that extracts token from the url param string.
func CsrfFromParam(param string) func(ctx *gin.Context) (string, error) {
	return func(ctx *gin.Context) (string, error) {
		token := ctx.Param(param)
		if token == "" {
			return "", errMissingParam
		}
		return token, nil
	}
}

// CsrfFromForm returns a function that extracts a token from a multipart-form.
//func CsrfFromForm(param string) func(ctx *gin.Context) (string, error) {
//	return func(ctx *gin.Context) (string, error) {
//		token := ctx.FormValue(param)
//		if string(token) == "" {
//			return "", errMissingForm
//		}
//		return string(token), nil
//	}
//}

// CsrfFromHeader returns a function that extracts token from the request header.
func CsrfFromHeader(param string) func(ctx *gin.Context) (string, error) {
	return func(ctx *gin.Context) (string, error) {
		token := ctx.Request.Header.Get(param)
		if string(token) == "" {
			return "", errMissingHeader
		}
		return string(token), nil
	}
}

// CsrfFromQuery returns a function that extracts token from the query string.
func CsrfFromQuery(param string) func(ctx *gin.Context) (string, error) {
	return func(ctx *gin.Context) (string, error) {
		token := ctx.Query(param)
		if token == "" {
			return "", errMissingQuery
		}
		return token, nil
	}
}
