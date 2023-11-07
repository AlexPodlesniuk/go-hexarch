package main

import (
	"context"
	"user-acquisition/ports"
	"user-acquisition/service"
)

func main() {
	ctx := context.Background()
	application := service.NewApp(ctx)
	api := ports.NewApi(application)
	api.Start()
}
