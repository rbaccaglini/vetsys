package customer

import (
	"context"
	"vetsys/internal/domain/repository"
)

// GetCustomerByIDOutput is the output DTO of the Usecase (similar to the creation one, but complete)
type GetCustomerByIDOutput struct {
	ID      string
	Name    string
	Email   string
	Phone   string
	Address string
}

type GetCustomerByIDUsecase struct {
	Repo repository.CustomerRepository
}

func NewGetCustomerByIDUsecase(r repository.CustomerRepository) *GetCustomerByIDUsecase {
	return &GetCustomerByIDUsecase{Repo: r}
}

// Execute searches for a customer by ID.
func (uc *GetCustomerByIDUsecase) Execute(ctx context.Context, id string) (*GetCustomerByIDOutput, error) {
	// 1. Call the Repository Interface to search
	customer, err := uc.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 2. Return the Usecase output DTO (mapping the pure Entity)
	return &GetCustomerByIDOutput{
		ID:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}, nil
}
