package repository

import (
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepositoury{
		database,
	}
}

type userRepositoury struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
