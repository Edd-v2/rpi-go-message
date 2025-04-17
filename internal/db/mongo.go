package db

import (
	"context"
	"time"

	"github.com/Edd-v2/rpi-go-message/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo(cfg config.MongoConfig, log *logrus.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.WithError(err).Fatal("MongoDB connection failed")
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.WithError(err).Fatal("MongoDB ping failed")
	}

	MongoClient = client
	log.Info("MongoDB connection established")
}
