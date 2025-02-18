package holiday_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *holidayHandlerInterface) UpdateHoliday(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
