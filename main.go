package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.New()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
