package exam_domain

type ExamDomainInterface interface {
	GetID() string
	GetVeterinarianID() string
	GetClinicID() string
	GetOwnerID() string
	GetAnimalID() string
	GetExamTypeID() string
	GetSampleTypeID() string
	GetDeliveryDate() string

	SetID(string)
}

func NewExamDomain(veterinarianID, clinicID, ownerID, animalID, examTypeID, sampleTypeID, deliveryDate string) ExamDomainInterface {
	return &examDomain{
		veterinarianID: veterinarianID,
		clinicID:       clinicID,
		ownerID:        ownerID,
		animalID:       animalID,
		examTypeID:     examTypeID,
		sampleTypeID:   sampleTypeID,
		deliveryDate:   deliveryDate,
	}
}
