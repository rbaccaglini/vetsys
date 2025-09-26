package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// Import the layers
	domainRepo "vetsys/internal/domain/repository"
	httpHandler "vetsys/internal/handler/http"
	infraMongo "vetsys/internal/infra/repository/mongodb"
	customerUC "vetsys/internal/usecase/customer"
)

func main() {
	// 1. Configure Infrastructure (MongoDB)
	mockClient, mockDBName := infraMongo.ConnectToMongoDB("mongodb://admin:password@localhost:27017", "vetsysdb") // Função de conexão real

	// 2. Create the CONCRETE Implementation of the Repository (Infra)
	customerRepoImpl := infraMongo.NewCustomerRepository(mockClient, mockDBName)

	var customerRepo domainRepo.CustomerRepository = customerRepoImpl

	// 3. Create the Usecases, INJECTING the CONCRETE Implementation into the Interface
	createUC := customerUC.NewCreateCustomerUsecase(customerRepo)
	getByIDUC := customerUC.NewGetCustomerByIDUsecase(customerRepo)

	// 4. Create the HTTP Handler, INJECTING the Usecases
	customerHandler := httpHandler.NewCustomerHandler(createUC, getByIDUC)

	// 5. Configure the Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/v1", r)

	// Register Customer routes
	customerHandler.RegisterRoutes(r)

	// 6. Start the server
	log.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", r)
}
