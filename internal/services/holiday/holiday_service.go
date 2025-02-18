package holiday_service

import (
	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/rest_err.go"
)

func (s *holidayDomainService) FindAllHoliday() ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {

	list, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *holidayDomainService) FindHolidayByDate(date string) ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (s *holidayDomainService) CreateHoliday(holiday holiday_domain.HolidayDomainInterface) (string, *rest_err.RestErr) {
	return s.repo.Insert(holiday)
}

func (s *holidayDomainService) UpdateHoliday(holiday holiday_domain.HolidayDomainInterface) (holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
