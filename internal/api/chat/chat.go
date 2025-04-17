package chat

import (
	"net/http"

	"github.com/Edd-v2/rpi-go-message/dto"
	"github.com/Edd-v2/rpi-go-message/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Logger
}

func NewHandler(log *logrus.Logger) *Handler {
	return &Handler{log: log}
}

func (h *Handler) StartChatHandler(c *gin.Context) {
	var req dto.StartChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warnf("[chat] invalid start chat input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userId").(string)

	chat, err := service.StartPrivateChat(userID, req.TargetID, h.log)
	if err != nil {
		h.log.Errorf("[chat] failed to start chat: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not start chat"})
		return
	}

	resp := dto.ChatResponse{
		ID:      chat.ID.Hex(),
		Members: []string{chat.Members[0].Hex(), chat.Members[1].Hex()},
		IsGroup: false,
	}

	c.JSON(http.StatusOK, resp)
}
