package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"vetsys/internal/domain/entity"
	domain "vetsys/internal/domain/repository"
)

// O nome da nossa collection no MongoDB
const customerCollection = "customers"

// CustomerRepository implementa a interface domain.CustomerRepository
// E é a única que tem acesso ao cliente MongoDB.
type CustomerRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

// Ensure que CustomerRepository implementa a interface do Domínio
var _ domain.CustomerRepository = (*CustomerRepository)(nil)

// NewCustomerRepository é o construtor da implementação.
func NewCustomerRepository(client *mongo.Client, dbName string) *CustomerRepository {
	return &CustomerRepository{
		client: client,
		db:     client.Database(dbName),
	}
}

// ----------------------------------------------------------------------
// Implementação dos Métodos da Interface de Domínio
// ----------------------------------------------------------------------

// Save persiste ou atualiza um novo Customer.
func (r *CustomerRepository) Save(customer *entity.Customer) (*entity.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Converte a Entidade de Domínio para o DAO de Persistência (MongoDB)
	dao := FromDomainEntity(customer)

	collection := r.db.Collection(customerCollection)

	if dao.ID.IsZero() {
		// Novo registro (Insert)
		dao.ID = primitive.NewObjectID()
		dao.CreatedAt = time.Now()
		dao.UpdatedAt = time.Now()

		_, err := collection.InsertOne(ctx, dao)
		if err != nil {
			return nil, err
		}
	} else {
		// Registro existente (Update)
		dao.UpdatedAt = time.Now()

		filter := bson.M{"_id": dao.ID}
		// Usamos $set para atualizar apenas os campos que queremos
		update := bson.M{"$set": dao}

		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
	}

	// 2. Converte o DAO atualizado (com ID e timestamps) de volta para a Entidade de Domínio
	return dao.ToDomainEntity(), nil
}

// FindByID busca um Customer pelo seu identificador único.
func (r *CustomerRepository) FindByID(id string) (*entity.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	var dao CustomerDAO
	err = r.db.Collection(customerCollection).FindOne(ctx, filter).Decode(&dao)

	if err == mongo.ErrNoDocuments {
		//return nil, domain.ErrCustomerNotFound // Retorna um erro amigável/de domínio
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// Converte o DAO lido do DB para a Entidade de Domínio
	return dao.ToDomainEntity(), nil
}

func (r *CustomerRepository) FindAll() ([]*entity.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *CustomerRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
