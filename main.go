package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hebitigo/fib_api/handler"
)

func main() {
	r := gin.Default()
	r.GET("/fib", handler.FibonacciHandler)

	r.Run(":8080")
}
