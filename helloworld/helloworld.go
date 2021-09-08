package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("starting server...")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)
	r.Run(":" + port)
}
