package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"monorepo/pkg/config"
)

type App struct {
	Port string `json:"port" yaml:"port"`
}

type Configuration struct {
	App App `json:"app" yaml:"app"`
}

var cfg *Configuration

var (
	configPath = flag.String("config", "./config/config.yaml", "User Account Service address")
)

func main() {
	flag.Parse()
	cfg = config.Init[Configuration](*configPath)

	r := gin.Default()
	r.GET("/api/:version/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		version := c.Param("version")
		c.JSON(200, gin.H{
			"user_name": username,
			"password":  password,
			"version":   version,
		})
	})

	srv := &http.Server{
		Addr:    ":" + cfg.App.Port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil { // listen and serve on 0.0.0.0:8080
			log.Printf("listen: %s\n", err)
		}
	}()

	gracefulStop(srv)
}

// 优雅关闭
func gracefulStop(srv *http.Server) {
	// 监听终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// 设置优雅关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 关闭服务并等待请求处理完成
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited gracefully")
}
