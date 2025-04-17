package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    primitive.ObjectID `bson:"chat_id"`
	SenderID  primitive.ObjectID `bson:"sender_id"`
	Content   string             `bson:"content"`
	Type      string             `bson:"type"` // "text", "image" (later)
	Timestamp time.Time          `bson:"timestamp"`
}
