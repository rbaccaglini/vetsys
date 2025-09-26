package repository

import (
	"vetsys/internal/domain/entity"
)

// CustomerRepository is the interface that defines the persistence contract for the Customer entity.
// This interface is used and known ONLY by Usecases.
type CustomerRepository interface {
	// Save persists or updates a new Customer.
	Save(customer *entity.Customer) (*entity.Customer, error)

	// FindByID searches for a Customer by its unique identifier.
	FindByID(id string) (*entity.Customer, error)

	// FindAll returns a list of Customers.
	FindAll() ([]*entity.Customer, error)

	// Delete removes a Customer by ID.
	Delete(id string) error
}
