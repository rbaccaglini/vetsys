package converter

import (
	"time"
)

func ConvertStringDateToTimestamp(dateStr string) (int64, error) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}
	date := parsedTime.UnixMilli()
	return date, nil
}

func ConvertDateTimestampToString(unixMilli int64) string {
	t := time.UnixMilli(unixMilli).UTC()
	formattedDate := t.Format("2006-01-02")
	return formattedDate
}
