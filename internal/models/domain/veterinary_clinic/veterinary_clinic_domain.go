package veterinary_clinic_domain

import "github.com/rbaccaglini/vetsys/internal/models/domain/common"

type veterinaryClinicDomain struct {
	id        string
	cnpj      string
	tradeName string
	legalName string
	address   common.Address // chave: "cep", "city", "street", "number", "neighborhood"
	phones    []common.Phone // um marcado como principal
	emails    []common.Email // um marcado como principal
}

func (v *veterinaryClinicDomain) GetID() string {
	return v.id
}

func (v *veterinaryClinicDomain) SetID(id string) {
	v.id = id
}

func (v *veterinaryClinicDomain) GetCNPJ() string {
	return v.cnpj
}

func (v *veterinaryClinicDomain) GetTradeName() string {
	return v.tradeName
}

func (v *veterinaryClinicDomain) GetLegalName() string {
	return v.legalName
}

func (v *veterinaryClinicDomain) GetAddress() common.Address {
	return v.address
}

func (v *veterinaryClinicDomain) GetPhones() []common.Phone {
	return v.phones
}

func (v *veterinaryClinicDomain) GetEmails() []common.Email {
	return v.emails
}
