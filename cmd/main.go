package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gndw/starting-golang/internals/resources"
)

func main() {

	// startup context
	ctx := context.Background()

	// initiating resources
	resource, err := resources.Init(ctx)
	if err != nil {
		log.Fatalf("[FATAL] startup failed with err: %v", err.Error())
	}

	// starting http server
	err = resource.HttpServerService.Start(ctx)
	if err != nil {
		log.Fatalf("[FATAL] http-server failed with err: %v", err.Error())
	}

	// Channel to listen for signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	<-stop

	// Shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = resource.HttpServerService.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("[FATAL] http-server shutdown failed with err: %v", err.Error())
	}
	log.Println("Server gracefully stopped")
}
