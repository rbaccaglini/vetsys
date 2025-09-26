package mongodb

import (
	"time"

	"vetsys/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerDAO (Data Access Object) representa a estrutura do documento conforme ele é armazenado no MongoDB.
type CustomerDAO struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	Name    string `bson:"name"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Address string `bson:"address"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// ToDomainEntity converte o DAO de Persistência em uma Entidade de Domínio.
func (dao *CustomerDAO) ToDomainEntity() *entity.Customer {
	return &entity.Customer{
		// O ID BSON precisa ser convertido para string para a Entidade de Domínio
		ID:        dao.ID.Hex(),
		Name:      dao.Name,
		Email:     dao.Email,
		Phone:     dao.Phone,
		Address:   dao.Address,
		CreatedAt: dao.CreatedAt,
	}
}

// FromDomainEntity converte uma Entidade de Domínio em um DAO de Persistência.
// É usado quando vamos salvar ou atualizar no MongoDB.
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
