package holiday_domain

type HolidayDomainInterface interface {
	GetID() string
	GetDate() string
	GetHolidayType() string
	GetDescription() string
	GetIsCyclical() bool

	SetID(string)
}

func NewHolidayDomain(date, holidayType, description string, isCyclical bool) HolidayDomainInterface {
	return &holidayDomain{
		date:        date,
		holidayType: holidayType,
		description: description,
		isCyclical:  isCyclical,
	}
}
