package main

import (
	"github.com/gin-gonic/gin"
	"github.com/japnoor/daelog/internal/config"
	"github.com/japnoor/daelog/internal/db"
	"github.com/japnoor/daelog/internal/handlers"
	"github.com/japnoor/daelog/internal/middleware"
	"github.com/japnoor/daelog/internal/services"
	"github.com/japnoor/daelog/pkg/constants"
	"github.com/japnoor/daelog/pkg/logger"
	"github.com/japnoor/daelog/pkg/response"
)

func main() {
	logger.Init()

	cfg := config.Load()
	mongodb := db.Connect(cfg.MongoURI)

	eventService := services.NewEventService(mongodb)
	eventHandler := handlers.NewEventHandler(eventService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.RequestLogger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		response.OK(c, constants.MsgHealthOK)
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			response.OK(c, constants.MsgHealthOK)
		})
		v1.POST("/events", eventHandler.Create)
		v1.GET("/events", eventHandler.GetByDate)
	}

	logger.Info("server starting").Str("port", cfg.Port).Send()
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Fatal("server failed to start", err)
	}
}
