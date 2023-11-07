package main

import (
	"context"
	"user-alerts/ports"
	"user-alerts/service"
)

func main() {
	ctx := context.Background()
	application := service.NewApp(ctx)

	events := ports.NewEvents(application)
	go events.Start()

	api := ports.NewApi(application)
	api.Start()
}
