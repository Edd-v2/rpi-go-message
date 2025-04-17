package service

import (
	db_model "github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/internal/repository"
	"github.com/sirupsen/logrus"
)

func StartPrivateChat(userID, targetID string, log *logrus.Logger) (*db_model.Chat, error) {
	return repository.FindOrCreatePrivateChat(userID, targetID, log)
}
