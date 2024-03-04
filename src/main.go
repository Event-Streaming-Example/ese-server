package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ese.server/app"
)

func shutDownResources(app *app.App) {
	log.Println("Attempting shutting down resources gracefully")
	app.Stop()
	log.Println("Resources shutdown gracefully")
}

func main() {
	PORT := os.Getenv("SERVER_PORT")
	config := app.ParseConfig()

	app := &app.App{}
	app.Initialize(config, PORT)

	go func() {
		app.Start()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutDownResources(app)

}
