package sample_type_domain

type sampleTypeDomain struct {
	id   string
	name string
}

func (s *sampleTypeDomain) GetID() string {
	return s.id
}

func (s *sampleTypeDomain) SetID(id string) {
	s.id = id
}

func (s *sampleTypeDomain) GetName() string {
	return s.name
}
