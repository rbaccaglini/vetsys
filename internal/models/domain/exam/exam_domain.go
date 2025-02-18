package exam_domain

type examDomain struct {
	id             string
	veterinarianID string
	clinicID       string
	ownerID        string
	animalID       string
	examTypeID     string
	sampleTypeID   string
	deliveryDate   string
}

func (e *examDomain) GetID() string {
	return e.id
}

func (e *examDomain) SetID(id string) {
	e.id = id
}

func (e *examDomain) GetVeterinarianID() string {
	return e.veterinarianID
}

func (e *examDomain) GetClinicID() string {
	return e.clinicID
}

func (e *examDomain) GetOwnerID() string {
	return e.ownerID
}

func (e *examDomain) GetAnimalID() string {
	return e.animalID
}

func (e *examDomain) GetExamTypeID() string {
	return e.examTypeID
}

func (e *examDomain) GetSampleTypeID() string {
	return e.sampleTypeID
}

func (e *examDomain) GetDeliveryDate() string {
	return e.deliveryDate
}
