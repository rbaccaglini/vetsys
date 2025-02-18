package holiday_repo

import (
	"context"
	"os"

	holiday_domain "github.com/rbaccaglini/vetsys/internal/models/domain/holiday"
	"github.com/rbaccaglini/vetsys/internal/models/entity"
	"github.com/rbaccaglini/vetsys/internal/util/converter"
	"github.com/rbaccaglini/vetsys/pkg/utils/logger"
	"github.com/rbaccaglini/vetsys/pkg/utils/rest_err.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (r *holidayRepository) GetByDate(date string) ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {
	return nil, nil
}

func (r *holidayRepository) GetAll() ([]holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindAllUsers")
	logger.Info("Init Find All Users repository", journey)

	filter := bson.D{}
	collection := r.databaseConnection.Collection(os.Getenv(MONGODB_HOLIDAY_COLLECTION))
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		errMessage := "Error trying to get all holidays"
		logger.Error(errMessage, err, journey)
		return nil, rest_err.NewInternalServerError(errMessage)
	}
	defer cur.Close(context.Background())

	resp := []holiday_domain.HolidayDomainInterface{}
	for cur.Next(context.Background()) {
		userEntity := &entity.Holiday{}
		err := cur.Decode(userEntity)
		if err != nil {
			errMessage := "Error trying to get all holidays"
			logger.Error(errMessage, err, journey)
			return nil, rest_err.NewInternalServerError(errMessage)
		}
		userDomain := converter.ConverterEntityToDomain(*userEntity)
		resp = append(resp, userDomain)
	}

	return resp, nil
}

func (r *holidayRepository) Insert(h holiday_domain.HolidayDomainInterface) (string, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateNewHoliday")
	logger.Info("Init Insert Holiday Repository", journey)

	e, err := converter.ConverterDomainToEntity(h)
	if err != nil {
		return "", rest_err.NewInternalServerError(err.Error())
	}

	collection := r.databaseConnection.Collection(os.Getenv(MONGODB_HOLIDAY_COLLECTION))
	result, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return "", rest_err.NewInternalServerError(err.Error())
	}
	e.ID = result.InsertedID.(primitive.ObjectID)

	return e.ID.Hex(), nil
}

func (r *holidayRepository) Update(h holiday_domain.HolidayDomainInterface) (holiday_domain.HolidayDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
