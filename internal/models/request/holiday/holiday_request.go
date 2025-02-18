package holiday_request

type HolidayRequest struct {
	Date        string `json:"date" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCyclical  bool   `json:"is_cyclical" binding:"required"`
}
