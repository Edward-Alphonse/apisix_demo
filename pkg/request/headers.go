package request

// import (
//
//	"encoding/json"
//	"github.com/saywo-wh/saywo_pkg/utils"
//	"strings"
//
// )
type Headers struct {
	Authorization string
	ContentType   string `header:"content-type"`
}

//
//func (h *Headers) JWT() string {
//	return h.parseBearerAuth()
//}
//
//func (h *Headers) parseBearerAuth() (token string) {
//	const prefix = "Bearer "
//	auth := h.Authorization
//	if len(auth) < len(prefix) || !utils.EqualFold(auth[:len(prefix)], prefix) {
//		return ""
//	}
//	return auth[len(prefix):]
//}
//
//func (h *Headers) IsDevApp() bool {
//	return h.XDev || strings.ToLower(h.XChannel) == "dev"
//}
//
//func (p Headers) ToMap() map[string]any {
//	data := make(map[string]any)
//	bytes, err := json.Marshal(p)
//	if err != nil {
//		return data
//	}
//
//	err = json.Unmarshal(bytes, &data)
//	if err != nil {
//		return data
//	}
//	return data
//}
