package validator_request

import (
	"fmt"
	"strings"
	"time"

	"github.com/rbaccaglini/vetsys/pkg/utils/rest_err.go"
)

const DATE_REQUEST_FORMAT = "2006-01-02"

var HOLIDAY_TYPE_OPTIONS = []string{"international", "national", "state", "municipal"}

func IsValidDate(date string) *rest_err.RestErr {
	msg := fmt.Sprintf("Date (%s) with incorrect format", date)
	_, err := time.Parse(DATE_REQUEST_FORMAT, date)
	if err != nil {
		err := rest_err.NewBadRequestError(fmt.Sprintf("%s, error=%s", msg, err.Error()))
		return err
	}
	return nil
}

func IsValidHolidayType(ht string) *rest_err.RestErr {
	isValidType := false
	for _, v := range HOLIDAY_TYPE_OPTIONS {
		if v == strings.ToLower(ht) {
			isValidType = true
		}
	}
	if !isValidType {
		msg := fmt.Sprintf("Request with incorrect holiday type (%s) ", ht)
		err := rest_err.NewBadRequestError(msg)
		return err
	}
	return nil
}
