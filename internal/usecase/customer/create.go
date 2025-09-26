package customer

import (
	"context"
	"vetsys/internal/domain/entity"
	"vetsys/internal/domain/repository"
)

// CreateCustomerInput é o DTO de entrada para o Usecase.
type CreateCustomerInput struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// CreateCustomerOutput é o DTO de saída do Usecase.
type CreateCustomerOutput struct {
	ID string
}

// CreateCustomerUsecase é a struct de nosso caso de uso.
type CreateCustomerUsecase struct {
	Repo repository.CustomerRepository // Dependência da Interface
}

// NewCreateCustomerUsecase é o construtor.
func NewCreateCustomerUsecase(r repository.CustomerRepository) *CreateCustomerUsecase {
	return &CreateCustomerUsecase{Repo: r}
}

// Execute orquestra a criação de um novo cliente.
func (uc *CreateCustomerUsecase) Execute(ctx context.Context, input CreateCustomerInput) (*CreateCustomerOutput, error) {
	// 1. Cria a Entidade de Domínio a partir do DTO de entrada do Usecase
	customer := &entity.Customer{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Address: input.Address,
	}

	// 2. Chama a Interface de Repositório (sem saber que é MongoDB)
	newCustomer, err := uc.Repo.Save(customer)
	if err != nil {
		return nil, err
	}

	// 3. Retorna o DTO de saída do Usecase
	return &CreateCustomerOutput{
		ID: newCustomer.ID,
	}, nil
}
