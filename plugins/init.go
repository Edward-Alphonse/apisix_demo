package plugins

import (
	"apisix_demo/plugins/auth"
	"apisix_demo/plugins/says"
)

func Init() {
	says.Init()
	auth.Init()
}
