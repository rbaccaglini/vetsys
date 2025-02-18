package holiday_service

import (
	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	holiday_repo "github.com/rbaccaglini/vetsys/internal/repositories/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/rest_err.go"
)

type HolidayDomainServiceInterface interface {
	FindAllHoliday() ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
	FindHolidayByDate(date string) ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
	CreateHoliday(holiday holiday_domain.HolidayDomainInterface) (string, *rest_err.RestErr)
	UpdateHoliday(holiday holiday_domain.HolidayDomainInterface) (holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
}

func NewHolidayDomainService(userRepo holiday_repo.HolidayRepository) HolidayDomainServiceInterface {
	return &holidayDomainService{userRepo}
}

type holidayDomainService struct {
	repo holiday_repo.HolidayRepository
}
