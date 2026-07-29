package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectMongo() error {

	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DATABASE_NAME")

	if uri == "" {
		return errors.New("MONGO_URI no configurada")
	}

	if dbName == "" {
		return errors.New("DATABASE_NAME no configurada")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().
			ApplyURI(uri),
	)

	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error al pingear MongoDB: %w", err)
	}

	Client = client
	Database = client.Database(dbName)
	fmt.Println("Base de datos:", dbName)

	collections, err := Database.ListCollectionNames(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Colecciones:", collections)
	}
	return nil
}
