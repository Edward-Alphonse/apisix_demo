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
	"monorepo/service/http/hello_world/client"
	"monorepo/service/rpc/hello_word/kitex_gen/monorepo/rpc/hello_world"
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
	client.Init([]string{"127.0.0.1:2379"})

	port := cfg.App.Port
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		resp, err := client.HelloWorldClient.GetHelloWorld(c, &hello_world.GetHelloWorldRequest{
			Name: "ryan he",
		})
		if err != nil {
			log.Printf("get hello world err: %v", err)
			c.JSON(200, nil)
			return
		}
		c.JSON(200, gin.H{
			"name": resp.GetResult_(),
		})
	})

	srv := &http.Server{
		Addr:    ":" + port,
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
