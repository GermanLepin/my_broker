package connection

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://mongo:27017"
)

func ConnectToMongo() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connecting:", err)
		return nil
	}

	log.Println("connected to mongo!")

	return conn
}
