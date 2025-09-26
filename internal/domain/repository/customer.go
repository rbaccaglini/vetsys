package repository

import (
	"vetsys/internal/domain/entity"
)

// CustomerRepository é a interface que define o contrato de persistência para a entidade Customer.
// Esta interface é usada e conhecida SOMENTE pelos Usecases.
type CustomerRepository interface {
	// Save persiste ou atualiza um novo Customer.
	Save(customer *entity.Customer) (*entity.Customer, error)

	// FindByID busca um Customer pelo seu identificador único.
	FindByID(id string) (*entity.Customer, error)

	// FindAll retorna uma lista de Customers.
	FindAll() ([]*entity.Customer, error)

	// Delete remove um Customer pelo ID.
	Delete(id string) error
}
