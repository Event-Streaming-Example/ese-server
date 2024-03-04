package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"ese.server/redis"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Server      *http.Server
	RedisClient *redis.RedisClient
}

func (a *App) Initialize(config Config, port string) {
	address := fmt.Sprintf(":%s", port)
	a.setupRedisClient(config.Redis)
	a.setupCounters()

	router := gin.Default()
	a.setupRouters(router, address)
	a.setupMiddleware(router)

	a.Server = &http.Server{
		Addr:    address,
		Handler: router,
	}
}

func (a *App) setupRedisClient(config redis.Config) {
	redisClient := &redis.RedisClient{}
	redisClient.Initialize(config)
	a.RedisClient = redisClient
}

func (a *App) setupCounters() {
	prometheus.MustRegister(healthCounter)
	prometheus.MustRegister(addEventCounter)
	prometheus.MustRegister(addEventsCounter)
	prometheus.MustRegister(getEventsCounter)
}

func (a *App) setupMiddleware(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length"}
	router.Use(cors.New(config))
}

func (a *App) setupRouters(router *gin.Engine, address string) {
	handler := Handler{
		RedisClient: a.RedisClient,
		Address:     address,
	}

	router.GET("/", handler.GetHealth)
	router.GET("/events", handler.GetEvents)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.POST("/event", handler.AddEvent)
	router.POST("/events", handler.AddEvents)
}

func (a *App) Start(address string) {
	log.Println("Starting Server on : ", address)
	if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server : ", err)
	}
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	a.RedisClient.Destroy()
	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exited successfully")
}
