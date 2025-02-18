package veterinary_clinic_domain

import "github.com/rbaccaglini/vetsys/internal/models/domain/common"

type VeterinaryClinicDomainInterface interface {
	GetID() string
	GetCNPJ() string
	GetTradeName() string
	GetLegalName() string
	GetAddress() common.Address
	GetPhones() []common.Phone
	GetEmails() []common.Email

	SetID(string)
}

func NewVeterinaryClinicDomain(
	cnpj, tradeName, legalName string,
	address common.Address,
	phones []common.Phone,
	emails []common.Email,
) VeterinaryClinicDomainInterface {
	return &veterinaryClinicDomain{
		cnpj:      cnpj,
		tradeName: tradeName,
		legalName: legalName,
		address:   address,
		phones:    phones,
		emails:    emails,
	}
}
