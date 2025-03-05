package csrf

import (
	"errors"

	//"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
)

// Option is the only struct that can be used to set Options.
type Option struct {
	F func(o *Options)
}

const (
	csrfSecret     = "csrfSecret"
	csrfSalt       = "csrfSalt"
	csrfToken      = "csrfToken"
	csrfHeaderName = "X-CSRF-TOKEN"

	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

var (
	errMissingHeader = errors.New("[CSRF] missing csrf token in header")
	errMissingQuery  = errors.New("[CSRF] missing csrf token in query")
	errMissingParam  = errors.New("[CSRF] missing csrf token in param")
	errMissingForm   = errors.New("[CSRF] missing csrf token in form")
	errMissingSalt   = errors.New("[CSRF] missing salt")
	errInvalidToken  = errors.New("[CSRF] invalid token")
)

type CsrfNextHandler func(ctx *gin.Context) bool

type CsrfExtractorHandler func(ctx *gin.Context) (string, error)

type CsrfErrorHandler func(ctx *gin.Context, err error)

// Options defines the config for middleware.
type Options struct {
	// Secret used to generate token.
	//
	// Default: csrfSecret
	Secret string

	// Ignored methods will be considered no protection required.
	//
	// Optional. Default: "GET", "HEAD", "OPTIONS", "TRACE"
	IgnoreMethods []string

	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next CsrfNextHandler

	// KeyLookup is a string in the form of "<source>:<key>" that is used
	// to create an Extractor that extracts the token from the request.
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "param:<name>"
	// - "form:<name>"
	//
	// Optional. Default: "header:X-CSRF-TOKEN"
	KeyLookup string

	// ErrorFunc is executed when an error is returned from app.HandlerFunc.
	//
	// Optional. Default: func(ctx context.Context, c *app.RequestContext) { panic(c.Errors.Last()) }
	ErrorFunc CsrfErrorHandler //func( *gin.HandlerFunc, err error)

	// Extractor returns the csrf token.
	//
	// If set this will be used in place of an Extractor based on KeyLookup.
	//
	// Optional. Default will create an Extractor based on KeyLookup.
	Extractor CsrfExtractorHandler
}

func (o *Options) Apply(opts []Option) {
	for _, op := range opts {
		op.F(o)
	}
}

// OptionsDefault is the default options.
var OptionsDefault = Options{
	Secret: csrfSecret,
	// Assume that anything not defined as 'safe' by RFC7231 needs protection
	IgnoreMethods: []string{"GET", "HEAD", "OPTIONS", "TRACE"},
	Next:          nil,
	KeyLookup:     "header:" + csrfHeaderName,
	ErrorFunc: func(ctx *gin.Context, err error) {
		panic(err)
	},
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Secret:        OptionsDefault.Secret,
		IgnoreMethods: OptionsDefault.IgnoreMethods,
		Next:          OptionsDefault.Next,
		KeyLookup:     OptionsDefault.KeyLookup,
		ErrorFunc:     OptionsDefault.ErrorFunc,
	}
	options.Apply(opts)
	return options
}

// WithSecret sets secret.
func WithSecret(secret string) Option {
	return Option{
		F: func(o *Options) {
			o.Secret = secret
		},
	}
}

// WithIgnoredMethods sets methods that do not need to be protected.
func WithIgnoredMethods(methods []string) Option {
	return Option{
		F: func(o *Options) {
			o.IgnoreMethods = methods
		},
	}
}

// WithNext sets whether to skip this middleware.
func WithNext(f CsrfNextHandler) Option {
	return Option{
		F: func(o *Options) {
			o.Next = f
		},
	}
}

// WithKeyLookUp sets a string in the form of "<source>:<key>" that is used
// to create an Extractor that extracts the token from the request.
func WithKeyLookUp(lookup string) Option {
	return Option{
		F: func(o *Options) {
			o.KeyLookup = lookup
		},
	}
}

// WithErrorFunc sets ErrorFunc.
func WithErrorFunc(f CsrfErrorHandler) Option {
	return Option{
		F: func(o *Options) {
			o.ErrorFunc = f
		},
	}
}

// WithExtractor sets extractor.
func WithExtractor(f CsrfExtractorHandler) Option {
	return Option{
		F: func(o *Options) {
			o.Extractor = f
		},
	}
}
