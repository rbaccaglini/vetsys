package holiday_repo

import (
	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	"github.com/rbaccaglini/vetsys/pkg/utils/rest_err.go"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_HOLIDAY_COLLECTION = "MONGODB_HOLIDAY_COLLECTION"
)

func NewHolidayRepository(database *mongo.Database) HolidayRepository {
	return &holidayRepository{database}
}

type holidayRepository struct {
	databaseConnection *mongo.Database
}

type HolidayRepository interface {
	GetByDate(date string) ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
	GetAll() ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
	Insert(h holiday_domain.HolidayDomainInterface) (string, *rest_err.RestErr)
	Update(h holiday_domain.HolidayDomainInterface) (holiday_domain.HolidayDomainInterface, *rest_err.RestErr)
}
