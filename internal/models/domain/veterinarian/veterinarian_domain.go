package veterinarian_domain

import "github.com/rbaccaglini/vetsys/internal/models/domain/common"

type veterinarianDomain struct {
	id         string
	name       string
	crmv       string
	phones     []common.Phone // um marcado como principal
	emails     []common.Email // um marcado como principal
	workplaces []string       // lista de IDs de clínicas previamente cadastradas
}

func (v *veterinarianDomain) GetID() string {
	return v.id
}

func (v *veterinarianDomain) SetID(id string) {
	v.id = id
}

func (v *veterinarianDomain) GetName() string {
	return v.name
}

func (v *veterinarianDomain) GetCRMV() string {
	return v.crmv
}

func (v *veterinarianDomain) GetPhones() []common.Phone {
	return v.phones
}

func (v *veterinarianDomain) GetEmails() []common.Email {
	return v.emails
}

func (v *veterinarianDomain) GetWorkplaces() []string {
	return v.workplaces
}
