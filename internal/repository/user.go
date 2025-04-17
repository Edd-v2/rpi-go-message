package repository

import (
	"context"

	"github.com/Edd-v2/rpi-go-message/src/internal/db"
	db_model "github.com/Edd-v2/rpi-go-message/src/internal/model/db"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *db_model.User) error {
	coll := db.MongoClient.Database("rpi").Collection("users")
	_, err := coll.InsertOne(context.Background(), user)
	return err
}

func FindUserByPhone(phone string) (*db_model.User, error) {
	coll := db.MongoClient.Database("rpi").Collection("users")
	var user db_model.User
	err := coll.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
