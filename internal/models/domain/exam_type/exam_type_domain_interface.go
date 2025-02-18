package exam_type_domain

type ExamTypeDomainInterface interface {
	GetID() string
	GetName() string
	GetDeliveryTime() int
	GetPrice() float64
	GetSampleTypes() []string

	SetID(string)
}

func NewExamTypeDomain(name string, deliveryTime int, price float64, sampleTypes []string) ExamTypeDomainInterface {
	return &examTypeDomain{
		name:         name,
		deliveryTime: deliveryTime,
		price:        price,
		sampleTypes:  sampleTypes,
	}
}
