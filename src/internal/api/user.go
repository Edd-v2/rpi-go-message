package api

import (
	"github.com/Edd-v2/rpi-go-message/src/dto"
	"github.com/Edd-v2/rpi-go-message/src/internal/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user, err := service.RegisterUser(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dto.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
	})
}
