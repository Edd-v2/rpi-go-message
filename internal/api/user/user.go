package user

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
func (h *Handler) MeHandler(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	user, err := service.GetUserByID(userId, h.log)
	if err != nil {
		h.log.Warnf("[user] User not found: %s", userId)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	resp := dto.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SearchHandler(c *gin.Context) {
	username := c.Query("username")
	phone := c.Query("phone")
	userId := c.MustGet("userId").(string)

	users, err := service.SearchUsers(username, phone, userId, h.log)
	if err != nil {
		h.log.Errorf("[user] search failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, dto.UserResponse{
			ID:       u.ID.Hex(),
			Username: u.Username,
		})
	}

	c.JSON(http.StatusOK, responses)
}
