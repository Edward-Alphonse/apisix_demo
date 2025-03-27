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

package main

import (
	"flag"

	"github.com/Edward-Alphonse/logora"
	"github.com/Edward-Alphonse/logora/writers"
	"github.com/apache/apisix-go-plugin-runner/pkg/runner"

	"apisix_demo/internal/config"
	"apisix_demo/plugins"
)

var (
	configPath = flag.String("config", "./config/config.yaml", "User Account Service address")
)

func main() {
	flag.Parse()
	plugins.Init()

	cfg := config.Init(*configPath)
	runnerCfg := runner.RunnerConfig{}
	opts := make([]logora.LogOption, 0)
	opts = append(opts, logora.FileLog(cfg.Logs))
	logora.Register(opts...)
	logora.Info("start main")
	runnerCfg.LogLevel = writers.GetZapCoreLevel(cfg.Logs.Level)
	runnerCfg.LogOutput = writers.GetFileWriterSyncer(cfg.Logs)
	runner.Run(runnerCfg)
}
