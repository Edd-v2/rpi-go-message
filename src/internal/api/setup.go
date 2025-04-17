package api

import (
	"github.com/Edd-v2/rpi-go-message/src/internal/api/auth"
	"github.com/Edd-v2/rpi-go-message/src/internal/api/group"
	"github.com/Edd-v2/rpi-go-message/src/internal/api/user"

	"github.com/Edd-v2/rpi-go-message/src/internal/api/system"
	auth_middleware "github.com/Edd-v2/rpi-go-message/src/internal/middleware/auth"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, log *logrus.Logger) {
	api := router.Group("/api")

	authHandler := auth.NewHandler(log)
	userHandler := user.NewHandler(log)
	groupHandler := group.NewHandler(log)
	systemHandler := system.NewHandler(log)

	// Auth
	authGroup := api.Group("/auth")
	{
		authGroup.POST("/register", authHandler.RegisterHandler)
		authGroup.POST("/login", authHandler.LoginHandler)
	}

	// User (JWT protected)
	userGroup := api.Group("/user")
	userGroup.Use(auth_middleware.JWTAuthMiddleware(log))
	{
		userGroup.GET("/me", userHandler.MeHandler)
		userGroup.GET("/search", userHandler.MeHandler)
	}

	// Group (JWT protected)
	groupGroup := api.Group("/group")
	groupGroup.Use(auth_middleware.JWTAuthMiddleware(log))
	{
		groupGroup.POST("/create", groupHandler.CreateHandler)
		groupGroup.POST("/:id/invite", groupHandler.InviteHandler)
		groupGroup.GET("/:id/messages", groupHandler.GetMessagesHandler)
	}

	// System (public)
	api.GET("/healthz", systemHandler.HealthHandler)
	api.GET("/readyz", systemHandler.ReadyHandler)
	api.GET("/metrics", systemHandler.MetricsHandler)

	log.Debug("[READY] Api SetupRoutes done...")
}
