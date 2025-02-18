package sample_type_domain

type SampleTypeDomainInterface interface {
	GetID() string
	GetName() string

	SetID(string)
}

func NewSampleTypeDomain(name string) SampleTypeDomainInterface {
	return &sampleTypeDomain{name: name}
}
