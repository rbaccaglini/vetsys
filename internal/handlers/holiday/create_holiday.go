package holiday_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	holiday_request "github.com/rbaccaglini/vetsys/internal/models/request/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
	validation_err "github.com/rbaccaglini/vetsys/pkg/utils/validation/validator_err"
	"github.com/rbaccaglini/vetsys/pkg/utils/validation/validator_request"
	"go.uber.org/zap"
)

func (h *holidayHandlerInterface) CreateHoliday(c *gin.Context) {
	journey := zap.String("journey", "FindHolidayByDate")
	logger.Info("Init Find holidys by date", journey)

	var holidayRequest holiday_request.HolidayRequest

	if err := c.ShouldBindJSON(&holidayRequest); err != nil {
		msg := "There are some incorrect fields"
		logger.Error(msg, err, journey)
		errRest := validation_err.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, errRest)
		return
	}

	if err := validator_request.IsValidDate(holidayRequest.Date); err != nil {
		logger.Error(err.Message, err, journey)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := validator_request.IsValidHolidayType(holidayRequest.Type); err != nil {
		logger.Error("", err, journey)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	domain := holiday_domain.NewHolidayDomain(
		holidayRequest.Date,
		holidayRequest.Type,
		holidayRequest.Description,
		holidayRequest.IsCyclical,
	)

	id, err := h.service.CreateHoliday(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, id)
}
