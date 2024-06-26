package main

import (
	"github/nothiaki/idempotency/pkg/handlers/orders"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/orders", orders.Many)
  r.GET("/orders/:id", orders.Unique)
  r.POST("/orders", orders.Create)

  r.Run()
}

