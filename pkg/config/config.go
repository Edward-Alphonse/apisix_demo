package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Init[P any](path string) *P {
	v := viper.New()
	v.SetConfigFile(path)
	// 查找并读取配置文件
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			v.SetDefault("port", "8080")
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	v.SetDefault("port", "9090")
	cfg := new(P)
	err := v.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
