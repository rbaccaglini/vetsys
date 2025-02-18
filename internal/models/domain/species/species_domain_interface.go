package species_domain

type SpeciesDomainInterface interface {
	GetID() string
	GetName() string

	SetID(string)
}

func NewSpeciesDomain(name string) SpeciesDomainInterface {
	return &speciesDomain{name: name}
}
