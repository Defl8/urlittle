package main

import (
	"github.com/gin-gonic/gin"
	"github.com/urlittle/pkg/server"
)

func main() {
	router := gin.Default()
	testEndpoint := EndpointNoArgs{"/test", getTest}
	router.GET(testEndpoint.Route, testEndpoint.Method)

	// This sets the trusted IPs, not really sure how this works
	router.SetTrustedProxies([]string{"localhost", "::1"})

	router.Run("localhost:42069")
}
