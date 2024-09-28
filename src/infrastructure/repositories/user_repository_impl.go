package repositories

import (
	"context"

	"github.com/nanda03dev/go-ms-template/src/domain"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/db"
	"github.com/nanda03dev/go-ms-template/src/infrastructure/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	dbs        *db.Databases
	collection *mongo.Collection
}

func NewUserRepository(dbs *db.Databases) domain.UserRepository {
	return &UserRepositoryImpl{dbs: dbs, collection: dbs.MongoDB.DB.Collection("users")}
}

func (r *UserRepositoryImpl) Save(user *domain.User) error {
	// Convert domain.User to entity.User
	userEntity := convertDomainUserToEntityUser(user)

	_, err := r.collection.InsertOne(context.TODO(), userEntity)
	return err
}

func (r *UserRepositoryImpl) FindById(id string) (*domain.User, error) {
	var user entity.User
	err := r.collection.FindOne(context.TODO(), bson.M{"userID": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return convertEntityUserToDomainUser(&user), nil
}

// Convert entity.User to domain.User
func convertEntityUserToDomainUser(eu *entity.User) *domain.User {
	return &domain.User{
		UserID:   eu.UserID,
		Name:     eu.Name,
		Password: eu.Password,
		Email:    eu.Email,
	}
}

// Convert domain.User to entity.User for MongoDB storage
func convertDomainUserToEntityUser(du *domain.User) *entity.User {
	return &entity.User{
		UserID:   du.UserID, // Generates a new ObjectID for Mongo
		Name:     du.Name,
		Password: du.Password,
		Email:    du.Email,
	}
}
