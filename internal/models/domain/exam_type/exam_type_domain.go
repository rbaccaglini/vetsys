package exam_type_domain

type examTypeDomain struct {
	id           string
	name         string
	deliveryTime int // em horas úteis
	price        float64
	sampleTypes  []string // lista de IDs de tipos de amostras permitidas
}

func (e *examTypeDomain) GetID() string {
	return e.id
}

func (e *examTypeDomain) SetID(id string) {
	e.id = id
}

func (e *examTypeDomain) GetName() string {
	return e.name
}

func (e *examTypeDomain) GetDeliveryTime() int {
	return e.deliveryTime
}

func (e *examTypeDomain) GetPrice() float64 {
	return e.price
}

func (e *examTypeDomain) GetSampleTypes() []string {
	return e.sampleTypes
}
