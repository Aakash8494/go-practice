package repository

import (
	"context"
	"myapp/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	db *mongo.Collection
}

// Constructor
func NewMongoUserRepository(db *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{db: db}
}

// Fulfilling the Contract: CreateUser
func (m *MongoUserRepository) CreateUser(user *domain.User) error {
	_, err := m.db.InsertOne(context.Background(), user)
	return err
}

// Fulfilling the Contract: GetUser
func (m *MongoUserRepository) GetUser(id string) (*domain.User, error) {
	var user domain.User
	filter := bson.M{"id": id}
	err := m.db.FindOne(context.Background(), filter).Decode(&user)
	return &user, err
}

// Placeholder for updates
func (m *MongoUserRepository) UpdateUser(user *domain.User) error { 
	return nil 
}

// Placeholder for deletes
func (m *MongoUserRepository) DeleteUser(id string) error { 
	return nil 
}
