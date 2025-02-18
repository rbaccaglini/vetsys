package holiday_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	holiday_response "github.com/rbaccaglini/vetsys/internal/models/response/holiday"
	"github.com/rbaccaglini/vetsys/internal/util/converter"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
	"go.uber.org/zap"
)

func (h *holidayHandlerInterface) FindAllHoliday(c *gin.Context) {
	journey := zap.String("journey", "FindAllHoliday")
	logger.Info("Init Find All holidys", journey)

	list, err := h.service.FindAllHoliday()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	listView := []holiday_response.HolidayResponse{}
	for _, h := range list {
		viewUser := converter.ConvertDomainToResponse(h)
		listView = append(listView, viewUser)
	}

	c.JSON(http.StatusOK, listView)
}

func (h *holidayHandlerInterface) FindHolidayByDate(c *gin.Context) {
	journey := zap.String("journey", "FindHolidayByDate")
	logger.Info("Init Find holidys by date", journey)

	c.JSON(http.StatusOK, nil)
}
