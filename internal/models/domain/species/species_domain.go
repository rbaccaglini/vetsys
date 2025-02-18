package species_domain

type speciesDomain struct {
	id   string
	name string
}

func (s *speciesDomain) GetID() string {
	return s.id
}

func (s *speciesDomain) SetID(id string) {
	s.id = id
}

func (s *speciesDomain) GetName() string {
	return s.name
}
