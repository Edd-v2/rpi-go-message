package service

import (
	"errors"

	"github.com/Edd-v2/rpi-go-message/dto"
	"github.com/Edd-v2/rpi-go-message/internal/middleware/auth"
	"github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/internal/repository"
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

func LoginUser(input dto.LoginRequest, log *logrus.Logger) (string, error) {
	user, err := repository.FindUserByPhone(input.Phone)
	if err != nil {
		log.Warnf("[service] User not found: %s", input.Phone)
		return "", errors.New("invalid credentials")
	}

	if !auth.CheckPassword(user.Password, input.Password) {
		log.Warnf("[service] Invalid password for: %s", input.Phone)
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID.Hex())
	if err != nil {
		log.Errorf("[service] Failed to generate token: %v", err)
		return "", err
	}

	return token, nil
}
