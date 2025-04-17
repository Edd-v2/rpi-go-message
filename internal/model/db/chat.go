package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string               `bson:"name,omitempty" json:"name,omitempty"` // solo per gruppi
	IsGroup     bool                 `bson:"is_group" json:"is_group"`
	Members     []primitive.ObjectID `bson:"members" json:"members"` // User IDs
	CreatedBy   primitive.ObjectID   `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
	LastMessage *primitive.ObjectID  `bson:"last_message,omitempty" json:"last_message,omitempty"` // ref ultimo messaggio
}
