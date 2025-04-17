package cmd

import (
	"github.com/Edd-v2/rpi-go-message/config"
	"github.com/Edd-v2/rpi-go-message/internal/api"
	"github.com/Edd-v2/rpi-go-message/internal/db"
	"github.com/Edd-v2/rpi-go-message/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	err := config.LoadConfiguration()
	if err != 0 {
		panic(err)
	}

	log := logger.SetupLogger(config.AppConfig.ServerConfig.Mode)
	log.Info("[START] Messaging Backend...")

	db.InitMongo(config.AppConfig.Mongo, log)

	initServer(log)
	log.Info("[READY] Messaging Backend API Ready")
}

func initServer(log *logrus.Logger) error {
	router := gin.Default()
	api.SetupRoutes(router, log)

	log.Info("[READY] Messaging Backend API Ready")
	return router.Run(":" + config.AppConfig.ServerConfig.Port)
}
