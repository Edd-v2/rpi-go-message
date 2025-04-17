package auth

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

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warnf("[auth] Invalid register request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	h.log.Infof("[auth] Attempting to register user: %s", req.Username)

	token, err := service.RegisterUser(req, h.log)
	if err != nil {
		h.log.Errorf("[auth] Registration failed: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.log.Infof("[auth] User %s registered successfully", req.Username)
	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.log.Warnf("[auth] Invalid login input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := service.LoginUser(req, h.log)
	if err != nil {
		h.log.Warnf("[auth] Login failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	h.log.Infof("[auth] Login successful for phone: %s", req.Phone)
	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}
