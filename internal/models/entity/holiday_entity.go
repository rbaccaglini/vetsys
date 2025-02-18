package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HolidayEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Date        int64              `bson:"date,omitempty"`
	Type        string             `bson:"type,omitempty"`
	Description string             `bson:"description,omitempty"`
	IsCyclical  bool               `bson:"is_cyclical,omitempty"`
}
