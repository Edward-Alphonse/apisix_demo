package csrf

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"math/rand"
	"net/textproto"
	"strings"
	"time"

	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// New validates CSRF token.
func New(opts ...Option) gin.HandlerFunc {
	cfg := NewOptions(opts...)
	selectors := strings.Split(cfg.KeyLookup, ":")

	if len(selectors) != 2 {
		panic(errors.New("[CSRF] KeyLookup must in the form of <source>:<key>"))
	}

	if cfg.Extractor == nil {
		// By default, we extract from a header
		cfg.Extractor = CsrfFromHeader(textproto.CanonicalMIMEHeaderKey(selectors[1]))

		switch selectors[0] {
		//case "form":
		//	cfg.Extractor = CsrfFromForm(selectors[1])
		case "query":
			cfg.Extractor = CsrfFromQuery(selectors[1])
		case "param":
			cfg.Extractor = CsrfFromParam(selectors[1])
		}
	}

	return func(ctx *gin.Context) {
		// Don't execute middleware if Next returns true
		//if cfg.Next != nil && cfg.Next(ctx) {
		//	ctx.Next()
		//	return
		//}
		//
		//session := sessions.Default(ctx)
		//ctx.Set(csrfSecret, cfg.Secret)
		//
		//if isIgnored(cfg.IgnoreMethods, ctx.Request.Method) {
		//	ctx.Next()
		//	return
		//}
		//
		//salt, ok := session.Get(csrfSalt).(string)
		//if !ok || len(salt) == 0 {
		//	cfg.ErrorFunc(ctx, errMissingSalt)
		//	return
		//}
		//
		//token, err := cfg.Extractor(ctx)
		//if err != nil {
		//	cfg.ErrorFunc(ctx, err)
		//	return
		//}
		//
		//if tokenize(cfg.Secret, salt) != token {
		//	cfg.ErrorFunc(ctx, errInvalidToken)
		//	return
		//}

		ctx.Next()
	}
}

// GetToken returns a CSRF token.
func GetToken(c *gin.Context) string {
	//session := sessions.Default(c)
	//secret := c.MustGet(csrfSecret).(string)
	//
	//if t, ok := c.Get(csrfToken); ok {
	//	return t.(string)
	//}
	//
	//salt, ok := session.Get(csrfSalt).(string)
	//if !ok {
	//	salt = randStr(16)
	//	session.Set(csrfSalt, salt)
	//	_ = session.Save()
	//}
	//token := tokenize(secret, salt)
	//c.Set(csrfToken, token)
	//
	//return token
	return ""
}

// tokenize generates token through secret and salt.
func tokenize(secret, salt string) string {
	h := sha256.New()
	_, _ = io.WriteString(h, salt+"-"+secret)
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return hash
}

// isIgnored determines whether the method is ignored.
func isIgnored(arr []string, value string) bool {
	ignore := false

	for _, v := range arr {
		if v == value {
			ignore = true
			break
		}
	}

	return ignore
}

var src = rand.NewSource(time.Now().UnixNano())

// randStr generates random string.
func randStr(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return sb.String()
}
