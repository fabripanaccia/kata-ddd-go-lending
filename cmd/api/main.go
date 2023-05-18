package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kata/ddd/cmd/dependencies"
	"github.com/kata/ddd/cmd/routes"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	time.Sleep(30 * time.Second)

	r := gin.New()

	deps := dependencies.NewDeps(ctx)

	routes.Routes(r, deps)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
