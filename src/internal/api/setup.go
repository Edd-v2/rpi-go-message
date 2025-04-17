package api

import (
	"github.com/Edd-v2/rpi-go-message/internal/api/auth"
	"github.com/Edd-v2/rpi-go-message/internal/api/group"
	"github.com/Edd-v2/rpi-go-message/internal/api/system"
	"github.com/Edd-v2/rpi-go-message/internal/api/user"
	"github.com/Edd-v2/rpi-go-message/internal/middleware"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, log *logrus.Logger) {
	api := router.Group("/api")

	authHandler := auth.NewHandler(log)
	userHandler := user.NewHandler(log)
	groupHandler := group.NewHandler(log)

	// Auth
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.RegisterHandler)
		authGroup.POST("/login", authHandler.LoginHandler)
	}

	// User (JWT protected)
	userGroup := api.Group("/user")
	userGroup.Use(middleware.JWTAuthMiddleware(log))
	{
		userGroup.GET("/me", userHandler.MeHandler)
		userGroup.GET("/search", userHandler.SearchHandler)
	}

	// Group (JWT protected)
	groupGroup := api.Group("/group")
	groupGroup.Use(middleware.JWTAuthMiddleware(log))
	{
		groupGroup.POST("/create", groupHandler.CreateHandler)
		groupGroup.POST("/:id/invite", groupHandler.InviteHandler)
		groupGroup.GET("/:id/messages", groupHandler.GetMessagesHandler)
	}

	// System (public)
	api.GET("/healthz", system.HealthHandler)
	api.GET("/readyz", system.ReadyHandler)
	api.GET("/metrics", system.MetricsHandler)

	log.Debug("[READY] Api SetupRoutes done...")
}
