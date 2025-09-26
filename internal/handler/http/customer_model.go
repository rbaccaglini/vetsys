package http

import "vetsys/internal/domain/entity"

// CreateCustomerRequest represents the input DTO (HTTP Body) for creating a new customer.
type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// ToDomainEntity convert the DTO to a domain entity
func (r *CreateCustomerRequest) ToDomainEntity() *entity.Customer {
	return &entity.Customer{
		Name:    r.Name,
		Email:   r.Email,
		Phone:   r.Phone,
		Address: r.Address,
	}
}
