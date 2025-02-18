package breed_domain

type BreedDomainInterface interface {
	GetID() string
	GetName() string
	GetSpeciesID() string

	SetID(string)
}

func NewBreedDomain(name, speciesID string) BreedDomainInterface {
	return &breedDomain{name: name, speciesID: speciesID}
}
