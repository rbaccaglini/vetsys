package customer

import (
	"context"
	"vetsys/internal/domain/repository"
)

// GetCustomerByIDOutput é o DTO de saída do Usecase (similar ao de criação, mas completo)
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

// Execute busca um cliente pelo ID.
func (uc *GetCustomerByIDUsecase) Execute(ctx context.Context, id string) (*GetCustomerByIDOutput, error) {
	// 1. Chama a Interface de Repositório para buscar
	customer, err := uc.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 2. Retorna o DTO de saída do Usecase (mapeando a Entidade pura)
	return &GetCustomerByIDOutput{
		ID:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}, nil
}
