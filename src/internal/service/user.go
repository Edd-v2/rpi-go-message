package service

import (
	"github.com/Edd-v2/rpi-go-message/src/dto"
	db_model "github.com/Edd-v2/rpi-go-message/src/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/src/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req dto.RegisterRequest) (*db_model.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	user := db_model.User{
		Username: req.Username,
		Phone:    req.Phone,
		Password: string(hashed),
	}

	if err := repository.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
