package bootstrap

import (
	"context"
	"log"
	"os"
	"time"

	"go-gin-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB instance
var DB *mongo.Database

// InitiliseMongo func creates connection with mongodb
func InitiliseMongo() (*mongo.Database, context.CancelFunc) {
	mongoDb, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.URI))
	if err != nil {
		log.Fatalf("Error while connecting to mongodb: %s\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = mongoDb.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting to mongodb: %s\n", err)
		os.Exit(1)
	}

	log.Println("Mongodb Connected")

	DB = mongoDb.Database(config.Config.Mongo.DB)
	return DB, cancel
}
