package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	custon_error "vetsys/internal/domain"
	"vetsys/internal/domain/entity"
	domain "vetsys/internal/domain/repository"
)

// The name of our collection in MongoDB
const customerCollection = "customers"

// CustomerRepository implements the domain.CustomerRepository interface
// And it's the only one that has access to the MongoDB client.
type CustomerRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

// Ensure that CustomerRepository implements the Domain interface
var _ domain.CustomerRepository = (*CustomerRepository)(nil)

// NewCustomerRepository is the implementation constructor.
func NewCustomerRepository(client *mongo.Client, dbName string) *CustomerRepository {
	return &CustomerRepository{
		client: client,
		db:     client.Database(dbName),
	}
}

// ---------------------------------------------------------------------
// Implementing Domain Interface Methods
// ---------------------------------------------------------------------------

// Save persists or updates a new Customer.
func (r *CustomerRepository) Save(customer *entity.Customer) (*entity.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Convert the Domain Entity to the Persistence DAO (MongoDB)
	dao := FromDomainEntity(customer)

	collection := r.db.Collection(customerCollection)

	if dao.ID.IsZero() {
		// Insert
		dao.ID = primitive.NewObjectID()
		dao.CreatedAt = time.Now()
		dao.UpdatedAt = time.Now()

		_, err := collection.InsertOne(ctx, dao)
		if err != nil {
			return nil, err
		}
	} else {
		// Update
		dao.UpdatedAt = time.Now()

		filter := bson.M{"_id": dao.ID}
		// We use $set to update only the fields we want
		update := bson.M{"$set": dao}

		_, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
	}

	// 2. Convert the updated DAO (with ID and timestamps) back to the Domain Entity
	return dao.ToDomainEntity(), nil
}

// FindByID searches for a Customer by its unique identifier.
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

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, custon_error.ErrCustomerNotFound
	}
	if err != nil {
		return nil, err
	}

	// Converts DAO read from DB to Domain Entity
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
