package holiday_handler

import (
	"github.com/gin-gonic/gin"
	holiday_service "github.com/rbaccaglini/vetsys/internal/services/holiday"
)

const DATE_REQUEST_FORMAT = "2006-01-02"

var HOLIDAY_TYPE_OPTIONS = []string{"international", "national", "state", "municipal"}

type HolidayHandlerInterface interface {
	FindAllHoliday(c *gin.Context)
	FindHolidayByDate(c *gin.Context)
	CreateHoliday(c *gin.Context)
	UpdateHoliday(c *gin.Context)
}

type holidayHandlerInterface struct {
	service holiday_service.HolidayDomainServiceInterface
}

func NewUserHandlerInterface(s holiday_service.HolidayDomainServiceInterface) HolidayHandlerInterface {
	return &holidayHandlerInterface{service: s}
}
