package holiday_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	holiday_response "github.com/rbaccaglini/vetsys/internal/models/response/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
	"go.uber.org/zap"
)

func FindAllHoliday(c *gin.Context) {
	journey := zap.String("journey", "FindAllHoliday")
	logger.Info("Init Find All holidys", journey)

	var holidays []holiday_response.HolidayResponse
	var h holiday_response.HolidayResponse
	h.ID = "1"
	h.Date = "2021-01-01"
	h.IsCyclical = true
	h.Description = "test..."
	h.Type = "NATIONAL"

	holidays = append(holidays, h)

	c.JSON(http.StatusOK, holidays)
}

func FindHolidayByDate(c *gin.Context) {
	journey := zap.String("journey", "FindHolidayByDate")
	logger.Info("Init Find holidys by date", journey)

	c.JSON(http.StatusOK, nil)
}
