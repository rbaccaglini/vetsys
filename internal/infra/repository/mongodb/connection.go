package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB inicializa a conexão com o cliente MongoDB.
// Se a conexão falhar, ela deve parar a aplicação ou retornar um erro.
func ConnectToMongoDB(uri string, dbName string) (*mongo.Client, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}

	// Opcional: Verifica se a conexão foi estabelecida
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Falha no Ping ao MongoDB: %v", err)
	}

	log.Println("Conexão com MongoDB estabelecida com sucesso.")

	// Retornamos o cliente e o nome do banco de dados (que é uma informação da infra)
	return client, dbName
}
