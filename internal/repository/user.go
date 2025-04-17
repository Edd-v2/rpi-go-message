package repository

import (
	"context"

	"github.com/Edd-v2/rpi-go-message/internal/db"
	db_model "github.com/Edd-v2/rpi-go-message/internal/model/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *db_model.User) error {
	coll := db.MongoClient.Database("rpi").Collection("users")
	_, err := coll.InsertOne(context.Background(), user)
	return err
}

func FindUserByID(id string) (*db_model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	coll := db.MongoClient.Database("rpi").Collection("users")
	var user db_model.User
	err = coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
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

func SearchUsers(username, phone, excludeUserId string) ([]*db_model.User, error) {
	coll := db.MongoClient.Database("rpi").Collection("users")

	filter := bson.M{}

	if username != "" {
		filter["username"] = bson.M{"$regex": primitive.Regex{Pattern: username, Options: "i"}}
	}
	if phone != "" {
		filter["phone"] = bson.M{"$regex": primitive.Regex{Pattern: phone, Options: "i"}}
	}

	if excludeUserId != "" {
		if objID, err := primitive.ObjectIDFromHex(excludeUserId); err == nil {
			filter["_id"] = bson.M{"$ne": objID}
		}
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []*db_model.User
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
