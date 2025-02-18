package animal_domain

type AnimalDomainInterface interface {
	GetID() string
	GetName() string
	GetBreedID() string
	GetBirthDate() string
	GetGender() string
	GetOwnerID() string

	SetID(string)
}

func NewAnimalDomain(name, breed_id, birthdate, gender, owner_id string) AnimalDomainInterface {
	return &animalDomain{
		name:      name,
		breed_id:  breed_id,
		birthdate: birthdate,
		gender:    gender,
		owner_id:  owner_id,
	}
}
