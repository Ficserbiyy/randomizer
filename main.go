package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin router with default middleware
	r := gin.Default()

	r.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
