package main

import (
	"context"
	"log"

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
	err = resource.HttpServerService.Serve(ctx)
	if err != nil {
		log.Fatalf("[FATAL] http-server failed with err: %v", err.Error())
	}

}
