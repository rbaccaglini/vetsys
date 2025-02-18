package owner_domain

import (
	"github.com/rbaccaglini/vetsys/internal/models/domain/common"
)

type ownerDomain struct {
	id      string
	name    string
	emails  []common.Email
	phones  []common.Phone
	cpf     string
	address common.Address // chave: "cep", "city", "street", "number", "neighborhood"
}

func (o *ownerDomain) GetID() string {
	return o.id
}

func (o *ownerDomain) SetID(id string) {
	o.id = id
}

func (o *ownerDomain) GetName() string {
	return o.name
}

func (o *ownerDomain) GetEmail() []common.Email {
	return o.emails
}

func (o *ownerDomain) GetPhone() []common.Phone {
	return o.phones
}

func (o *ownerDomain) GetCPF() string {
	return o.cpf
}

func (o *ownerDomain) GetAddress() common.Address {
	return o.address
}
