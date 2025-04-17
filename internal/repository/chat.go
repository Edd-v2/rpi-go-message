package repository

import (
	"context"
	"time"

	"github.com/Edd-v2/rpi-go-message/internal/db"
	db_model "github.com/Edd-v2/rpi-go-message/internal/model/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOrCreatePrivateChat(userID, targetID string, log *logrus.Logger) (*db_model.Chat, error) {
	coll := db.MongoClient.Database("rpi").Collection("chats")

	userObjID, _ := primitive.ObjectIDFromHex(userID)
	targetObjID, _ := primitive.ObjectIDFromHex(targetID)

	filter := bson.M{
		"is_group": false,
		"members": bson.M{
			"$all": []primitive.ObjectID{userObjID, targetObjID},
		},
	}

	var chat db_model.Chat
	err := coll.FindOne(context.TODO(), filter).Decode(&chat)
	if err == nil {
		log.Infof("[repository] existing chat found")
		return &chat, nil
	}

	// not found → create
	chat = db_model.Chat{
		IsGroup:   false,
		Members:   []primitive.ObjectID{userObjID, targetObjID},
		CreatedAt: time.Now(),
	}

	res, err := coll.InsertOne(context.TODO(), chat)
	if err != nil {
		return nil, err
	}

	chat.ID = res.InsertedID.(primitive.ObjectID)
	log.Infof("[repository] new chat created")
	return &chat, nil
}

func GetChatsByUserID(userID string) ([]*db_model.Chat, error) {
	coll := db.MongoClient.Database("rpi").Collection("chats")
	objID, _ := primitive.ObjectIDFromHex(userID)

	cursor, err := coll.Find(context.TODO(), bson.M{
		"members": objID,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var chats []*db_model.Chat
	if err := cursor.All(context.TODO(), &chats); err != nil {
		return nil, err
	}
	return chats, nil
}

func FindMessageByID(id string) (*db_model.Message, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	coll := db.MongoClient.Database("rpi").Collection("messages")

	var msg db_model.Message
	err = coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
