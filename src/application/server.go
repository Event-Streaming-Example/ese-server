package application

import (
	"context"
	"log"
	"net/http"
	"time"

	"ese/server/data"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	RedisClient data.RedisClient
	Server      http.Server
	Address     string
}

func ProvideServer(redisClient *data.RedisClient, address string) Server {
	router := provideRouter(redisClient, address)
	server := http.Server{
		Addr:    address,
		Handler: &router,
	}
	return Server{
		RedisClient: *redisClient,
		Server:      server,
		Address:     address,
	}
}

func provideRouter(redisClient *data.RedisClient, address string) gin.Engine {
	prometheus.MustRegister(healthCounter)
	prometheus.MustRegister(addEventCounter)
	prometheus.MustRegister(addEventsCounter)
	prometheus.MustRegister(getEventsCounter)

	router := gin.Default()
	controller := Controller{
		Storage: redisClient,
		Address: address,
	}

	router.GET("/", controller.Health)
	router.GET("/events", controller.GetEventLogs)
	router.POST("/event", controller.AddEvent)
	router.POST("/events", controller.AddEvents)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return *router
}

func (server *Server) Start() {
	log.Println("Starting Server on", server.Address)
	if err := server.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (server *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exited successfully")
}
