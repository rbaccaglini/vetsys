package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Holiday struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Date        int64              `bson:"date,omitempty"`
	Type        string             `bson:"type,omitempty"`
	Description string             `bson:"description,omitempty"`
	IsCyclical  bool               `bson:"is_cyclical,omitempty"`
}

type Species struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type Breed struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	SpeciesID primitive.ObjectID `bson:"species_id"`
}

type SampleType struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

type ExamType struct {
	ID             primitive.ObjectID   `bson:"_id,omitempty"`
	Name           string               `bson:"name"`
	DeliveryTime   int                  `bson:"delivery_time"` // in business hours
	Price          float64              `bson:"price"`
	AllowedSamples []primitive.ObjectID `bson:"allowed_samples"`
}

type Address struct {
	ZIPCode      string `bson:"zip_code"`
	City         string `bson:"city"`
	Street       string `bson:"street"`
	Number       string `bson:"number"`
	Neighborhood string `bson:"neighborhood"`
}

type Contact struct {
	Phone   string `bson:"phone"`
	Email   string `bson:"email"`
	Primary bool   `bson:"primary"`
}

type VeterinaryClinic struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	CNPJ          string             `bson:"cnpj"`
	TradeName     string             `bson:"trade_name"`
	CorporateName string             `bson:"corporate_name"`
	Address       Address            `bson:"address"`
	Contacts      []Contact          `bson:"contacts"`
}

type Veterinarian struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty"`
	Name     string               `bson:"name"`
	CRMV     string               `bson:"crmv"`
	Contacts []Contact            `bson:"contacts"`
	Clinics  []primitive.ObjectID `bson:"clinics"`
}

type Owner struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	CPF      string             `bson:"cpf,omitempty"`
	Address  Address            `bson:"address,omitempty"`
	Contacts []Contact          `bson:"contacts"`
}

type Animal struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	BreedID   primitive.ObjectID `bson:"breed_id"`
	BirthDate time.Time          `bson:"birth_date"`
	Gender    string             `bson:"gender"` // Male/Female
	OwnerID   primitive.ObjectID `bson:"owner_id"`
}

type Exam struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	VeterinarianID primitive.ObjectID `bson:"veterinarian_id"`
	ClinicID       primitive.ObjectID `bson:"clinic_id"`
	OwnerID        primitive.ObjectID `bson:"owner_id"`
	AnimalID       primitive.ObjectID `bson:"animal_id"`
	ExamTypeID     primitive.ObjectID `bson:"exam_type_id"`
	SampleID       primitive.ObjectID `bson:"sample_id"`
	RequestDate    int64              `bson:"request_date"`
	DeliveryDate   int64              `bson:"delivery_date"`
}
