package veterinarian_domain

import "github.com/rbaccaglini/vetsys/internal/models/domain/common"

type VeterinarianDomainInterface interface {
	GetID() string
	GetName() string
	GetCRMV() string
	GetPhones() []common.Phone
	GetEmails() []common.Email
	GetWorkplaces() []string

	SetID(string)
}

func NewVeterinarianDomain(
	name, crmv string,
	phones []common.Phone,
	emails []common.Email,
	clinics []string) VeterinarianDomainInterface {
	return &veterinarianDomain{
		name:       name,
		crmv:       crmv,
		phones:     phones,
		emails:     emails,
		workplaces: clinics,
	}
}
