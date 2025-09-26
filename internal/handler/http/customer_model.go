package http

import "vetsys/internal/domain/entity"

// CreateCustomerRequest representa o DTO de entrada (HTTP Body) para a criação de um novo cliente.
type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// ToDomainEntity converter o DTO em uma entidade de domínio
func (r *CreateCustomerRequest) ToDomainEntity() *entity.Customer {
	return &entity.Customer{
		Name:    r.Name,
		Email:   r.Email,
		Phone:   r.Phone,
		Address: r.Address,
	}
}
