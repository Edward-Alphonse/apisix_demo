package redis

import (
	"fmt"
	"time"
)

type Config struct {
	Host         string        `json:"host" yaml:"host"`
	Port         int32         `json:"port" yaml:"port"`
	User         string        `json:"user" yaml:"user"`
	Password     string        `json:"password" yaml:"password"`
	MaxRetries   int           `json:"max_retries" yaml:"max_retries"`
	DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
	PoolSize     int32         `json:"pool_size" yaml:"pool_size"`
	MinIdleConns int           `json:"min_idle_conns" yaml:"min_idle_conns"`
}

func (r *Config) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
