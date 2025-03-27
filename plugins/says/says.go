/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package says

import (
	"encoding/json"
	"net/http"

	"github.com/Edward-Alphonse/logora"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
)

func Init() {
	err := plugin.RegisterPlugin(&Say{})
	if err != nil {
		log.Fatalf("failed to register plugin say: %s", err)
	}
}

// Say is a demo to show how to return data directly instead of proxying
// it to the upstream.
type Say struct {
	// Embed the default plugin here,
	// so that we don't need to reimplement all the methods.
	plugin.DefaultPlugin
}

type SayConf struct {
	Body string `json:"body"`
}

func (p *Say) Name() string {
	return "say"
}

func (p *Say) ParseConf(in []byte) (interface{}, error) {
	// 在 apisix-dashboard上配置的 value进行解析
	logora.Info("says parse conf", logora.Field{
		"in": string(in),
	})
	conf := SayConf{}
	err := json.Unmarshal(in, &conf)
	return conf, err
}

func (p *Say) RequestFilter(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request) {
	body := conf.(SayConf).Body
	logora.Info("says request filter", logora.Field{
		"body": body,
	})
	if len(body) == 0 {
		return
	}

	w.Header().Add("X-Resp-A6-Runner", "Go")
}
