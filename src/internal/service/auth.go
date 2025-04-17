package service

import (
	"errors"

	"github.com/Edd-v2/rpi-go-message/internal/dto"
	"github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/internal/repository"
	"github.com/Edd-v2/rpi-go-message/src/internal/middleware/auth"
	"github.com/sirupsen/logrus"
)

func RegisterUser(input dto.RegisterRequest, log *logrus.Logger) (string, error) {
	log.Infof("[service] Registering user with phone: %s", input.Phone)

	existing, _ := repository.FindUserByPhone(input.Phone)
	if existing != nil {
		log.Warnf("[service] User already exists with phone: %s", input.Phone)
		return "", errors.New("user already exists")
	}

	hashed, err := auth.HashPassword(input.Password)
	if err != nil {
		log.Errorf("[service] Failed to hash password: %v", err)
		return "", err
	}

	user := &db.User{
		Username: input.Username,
		Phone:    input.Phone,
		Password: hashed,
	}

	err = repository.CreateUser(user)
	if err != nil {
		log.Errorf("[service] Failed to create user in DB: %v", err)
		return "", err
	}

	token, err := auth.GenerateToken(user.ID.Hex())
	if err != nil {
		log.Errorf("[service] Failed to generate JWT: %v", err)
		return "", err
	}

	log.Infof("[service] User %s created successfully", user.Username)
	return token, nil
}
