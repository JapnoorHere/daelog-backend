package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/japnoor/daelog/internal/config"
	"github.com/japnoor/daelog/internal/db"
)

func main() {
	cfg := config.Load()
	mongodb := db.Connect(cfg.MongoURI)
	_ = mongodb

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": "daelog backend ok"})
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "data": "daelog backend ok"})
		})
	}

	r.Run(":" + cfg.Port)
}
