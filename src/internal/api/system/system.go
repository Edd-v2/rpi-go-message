package system

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Logger
}

func NewHandler(log *logrus.Logger) *Handler {
	return &Handler{log: log}
}

func (h *Handler) HealthHandler(c *gin.Context) {

}

func (h *Handler) ReadyHandler(c *gin.Context) {

}

func (h *Handler) MetricsHandler(c *gin.Context) {

}
