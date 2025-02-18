package owner_domain

import (
	"github.com/rbaccaglini/vetsys/internal/models/domain/common"
)

type OwnerDomainInterface interface {
	GetID() string
	GetName() string
	GetEmail() []common.Email
	GetPhone() []common.Phone
	GetCPF() string
	GetAddress() common.Address

	SetID(string)
}

func NewOwnerDomain(
	name string,
	emails []common.Email,
	phones []common.Phone,
	cpf string,
	address common.Address) OwnerDomainInterface {
	return &ownerDomain{
		name:    name,
		emails:  emails,
		phones:  phones,
		cpf:     cpf,
		address: address,
	}
}
