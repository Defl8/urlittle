package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EndpointNoArgs struct {
	Route  string
	Method func(c *gin.Context)
}

func getTest(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the test endpoint.")
}

func main() {
	router := gin.Default()
	testEndpoint := EndpointNoArgs{"/test", getTest}
	router.GET(testEndpoint.Route, testEndpoint.Method)

	// This sets the trusted IPs, not really sure how this works
	router.SetTrustedProxies([]string{"localhost", "::1"})

	router.Run("localhost:42069")
}
