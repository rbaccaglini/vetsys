package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB initializes the connection to the MongoDB client.
// If the connection fails, it should stop the application or return an error.
func ConnectToMongoDB(uri string, dbName string) (*mongo.Client, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	// Optional: Checks if the connection was established
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Falha no Ping ao MongoDB: %v", err)
	}

	log.Println("Conexão com MongoDB estabelecida com sucesso.")

	// We return the client and the database name (which is information from the infrastructure)
	return client, dbName
}
