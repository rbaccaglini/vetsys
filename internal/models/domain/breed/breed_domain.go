package breed_domain

type breedDomain struct {
	id        string
	name      string
	speciesID string
}

func (b *breedDomain) GetID() string {
	return b.id
}

func (b *breedDomain) SetID(id string) {
	b.id = id
}

func (b *breedDomain) GetName() string {
	return b.name
}

func (b *breedDomain) GetSpeciesID() string {
	return b.speciesID
}
