package converter

import (
	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	"github.com/rbaccaglini/vetsys/internal/models/entity"
	holiday_response "github.com/rbaccaglini/vetsys/internal/models/response/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/converter"
)

func ConvertDomainToResponse(hd holiday_domain.HolidayDomainInterface) holiday_response.HolidayResponse {
	return holiday_response.HolidayResponse{
		ID:          hd.GetID(),
		Date:        hd.GetDate(),
		Type:        hd.GetHolidayType(),
		Description: hd.GetDescription(),
		IsCyclical:  hd.GetIsCyclical(),
	}
}

func ConverterDomainToEntity(d holiday_domain.HolidayDomainInterface) (*entity.Holiday, error) {
	dc, err := converter.ConvertStringDateToTimestamp(d.GetDate())
	if err != nil {
		return nil, err
	}

	return &entity.Holiday{
		Date:        dc,
		Type:        d.GetHolidayType(),
		Description: d.GetDescription(),
		IsCyclical:  d.GetIsCyclical(),
	}, nil
}

func ConverterEntityToDomain(entity entity.Holiday) holiday_domain.HolidayDomainInterface {
	domain := holiday_domain.NewHolidayDomain(
		converter.ConvertDateTimestampToString(entity.Date),
		entity.Type,
		entity.Description,
		entity.IsCyclical,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}
