package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spectertian/dianputest/config"
	"github.com/spectertian/dianputest/database"
	"github.com/spectertian/dianputest/router"
)

func main() {
	database.Init()

	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.Setup(r)

	log.Printf("Server starting on %s", config.Port)
	if err := r.Run(config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
