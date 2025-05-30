package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type resp struct {
	StatusCode int
	Body       string
}

func main() {
	router := gin.Default()

	// This sets the trusted IPs, not really sure how this works
	router.SetTrustedProxies([]string{"localhost", "::1"})

	router.Run("localhost:42069")
}
