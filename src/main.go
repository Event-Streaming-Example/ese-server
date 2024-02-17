package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ese/server/application"
	"ese/server/data"
)

func shutDownResources(client *data.RedisClient, server *application.Server) {
	log.Println("Attempting shutting down resources gracefully")
	server.Stop()
	client.Destroy()
	log.Println("Resources shutdown gracefully")
}

func main() {
	PORT := os.Getenv("SERVER_PORT")

	config := ProvideConfig()
	redisClient := ProvideRedisClient(config.Redis)
	server := ProvideServer(PORT, &redisClient)

	go func() {
		server.Start()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutDownResources(&redisClient, &server)

}
