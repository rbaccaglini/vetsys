package customer

import (
	"context"
	"vetsys/internal/domain/entity"
	"vetsys/internal/domain/repository"
)

// CreateCustomerInput is the input DTO for the Usecase.
type CreateCustomerInput struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// CreateCustomerOutput is the output DTO of the Usecase.
type CreateCustomerOutput struct {
	ID string
}

// CreateCustomerUsecase is the struct for our use case.
type CreateCustomerUsecase struct {
	Repo repository.CustomerRepository // Dependency on the Interface
}

// NewCreateCustomerUsecase is the constructor.
func NewCreateCustomerUsecase(r repository.CustomerRepository) *CreateCustomerUsecase {
	return &CreateCustomerUsecase{Repo: r}
}

// Execute orchestrates the creation of a new customer.
func (uc *CreateCustomerUsecase) Execute(ctx context.Context, input CreateCustomerInput) (*CreateCustomerOutput, error) {
	// 1. Create the Domain Entity from the Usecase input DTO
	customer := &entity.Customer{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Address: input.Address,
	}

	// 2. Call the Repository Interface (without knowing it's MongoDB)
	newCustomer, err := uc.Repo.Save(customer)
	if err != nil {
		return nil, err
	}

	// 3. Return the Usecase output DTO
	return &CreateCustomerOutput{
		ID: newCustomer.ID,
	}, nil
}
