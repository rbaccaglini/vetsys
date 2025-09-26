package mongodb

import (
	"time"

	"vetsys/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerDAO (Data Access Object) represents the structure of the document as it is stored in MongoDB.
type CustomerDAO struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Address string `bson:"address"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// ToDomainEntity converts the Persistence DAO into a Domain Entity.
func (dao *CustomerDAO) ToDomainEntity() *entity.Customer {
	return &entity.Customer{
		ID:        dao.ID.Hex(),
		Name:      dao.Name,
		Email:     dao.Email,
		Phone:     dao.Phone,
		Address:   dao.Address,
		CreatedAt: dao.CreatedAt,
	}
}

// FromDomainEntity converts a Domain Entity into a Persistence DAO.
// It is used when saving or updating in MongoDB.
func FromDomainEntity(e *entity.Customer) *CustomerDAO {
	objID, _ := primitive.ObjectIDFromHex(e.ID)

	return &CustomerDAO{
		ID:        objID,
		Name:      e.Name,
		Email:     e.Email,
		Phone:     e.Phone,
		Address:   e.Address,
		CreatedAt: e.CreatedAt,
	}
}
