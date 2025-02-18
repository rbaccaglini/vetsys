package holiday_domain

type holidayDomain struct {
	id          string
	date        string
	holidayType string
	description string
	isCyclical  bool
}

func (hd *holidayDomain) GetID() string {
	return hd.id
}

func (hd *holidayDomain) SetID(id string) {
	hd.id = id
}

func (hd *holidayDomain) GetDate() string {
	return hd.date
}

func (hd *holidayDomain) GetHolidayType() string {
	return hd.holidayType
}

func (hd *holidayDomain) GetDescription() string {
	return hd.description
}

func (hd *holidayDomain) GetIsCyclical() bool {
	return hd.isCyclical
}
