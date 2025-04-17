package api

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

func (h *Handler) MeHandler(c *gin.Context) {

}

func (h *Handler) SearchHandler(c *gin.Context) {

}
