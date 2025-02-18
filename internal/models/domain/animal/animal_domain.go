package animal_domain

type animalDomain struct {
	id        string
	name      string
	breed_id  string
	birthdate string
	gender    string
	owner_id  string
}

func (a *animalDomain) GetID() string {
	return a.id
}

func (a *animalDomain) SetID(id string) {
	a.id = id
}

func (a *animalDomain) GetName() string {
	return a.name
}

func (a *animalDomain) GetBreedID() string {
	return a.breed_id
}

func (a *animalDomain) GetBirthDate() string {
	return a.birthdate
}

func (a *animalDomain) GetGender() string {
	return a.gender
}

func (a *animalDomain) GetOwnerID() string {
	return a.owner_id
}
