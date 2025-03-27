package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Edward-Alphonse/logora"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

func Init() {
	err := plugin.RegisterPlugin(&Auth{})
	if err != nil {
		log.Fatalf("failed to register plugin say: %s", err)
	}
}

// Say is a demo to show how to return data directly instead of proxying
// it to the upstream.
type Auth struct {
	// Embed the default plugin here,
	// so that we don't need to reimplement all the methods.
	plugin.DefaultPlugin
}

type AuthConf struct {
	Body string `json:"body"`
}

func (p *Auth) Name() string {
	return "auth"
}

func (p *Auth) ParseConf(in []byte) (interface{}, error) {
	logora.Info("auth parse conf", logora.Field{
		"in": string(in),
	})
	conf := AuthConf{}
	if len(in) == 0 {
		return conf, nil
	}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

func (p *Auth) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	path := r.Path()
	logora.Info("auth start", logora.Field{
		"path":   string(path),
		"id":     r.ID(),
		"header": r.Header(),
	})
	body := `{"code": 1, "message": "request filter success"}`
	_, err := w.Write([]byte(body))
	if err != nil {
		log.Errorf("failed to write: %s", err)
	}
}
