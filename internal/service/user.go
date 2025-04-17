package service

import (
	db_model "github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/Edd-v2/rpi-go-message/internal/repository"
	"github.com/sirupsen/logrus"
)

func GetUserByID(id string, log *logrus.Logger) (*db_model.User, error) {
	user, err := repository.FindUserByID(id)
	if err != nil {
		log.Warnf("[service] GetUserByID failed: %v", err)
		return nil, err
	}
	return user, nil
}

func SearchUsers(username, phone, excludeUserId string, log *logrus.Logger) ([]*db_model.User, error) {
	users, err := repository.SearchUsers(username, phone, excludeUserId)
	if err != nil {
		log.Errorf("[service] SearchUsers failed: %v", err)
		return nil, err
	}
	return users, nil
}
