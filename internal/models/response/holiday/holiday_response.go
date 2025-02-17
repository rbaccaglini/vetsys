package holiday_response

type HolidayResponse struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	Type        string `json:"type"`
	Description string `json:"description"`
	IsCyclical  bool   `json:"is_cyclical"`
}
