package group

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

func (h *Handler) CreateHandler(c *gin.Context) {

}

func (h *Handler) InviteHandler(c *gin.Context) {

}

func (h *Handler) GetMessagesHandler(c *gin.Context) {

}
