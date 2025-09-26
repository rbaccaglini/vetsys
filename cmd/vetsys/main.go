package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// Importa as camadas
	domainRepo "vetsys/internal/domain/repository"
	httpHandler "vetsys/internal/handler/http"
	infraMongo "vetsys/internal/infra/repository/mongodb"
	customerUC "vetsys/internal/usecase/customer"
)

func main() {
	// 1. Configurar Infraestrutura (MongoDB)
	mockClient, mockDBName := infraMongo.ConnectToMongoDB("mongodb://admin:password@localhost:27017", "vetsysdb") // Função de conexão real

	// 2. Criar a Implementação CONCRETA do Repositório (Infra)
	customerRepoImpl := infraMongo.NewCustomerRepository(mockClient, mockDBName)

	var customerRepo domainRepo.CustomerRepository = customerRepoImpl

	// 3. Criar os Usecases, INJETANDO a Implementação CONCRETA na Interface
	createUC := customerUC.NewCreateCustomerUsecase(customerRepo)
	getByIDUC := customerUC.NewGetCustomerByIDUsecase(customerRepo)

	// 4. Criar o Handler HTTP, INJETANDO os Usecases
	customerHandler := httpHandler.NewCustomerHandler(createUC, getByIDUC)

	// 5. Configurar o Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", r)

	// Registrar as rotas do Customer
	customerHandler.RegisterRoutes(r)

	// 6. Iniciar o servidor
	log.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", r)
}
