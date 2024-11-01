package main

import (
	"github.com/VictorMont03/golang-users-app/src/controller/users"
	"github.com/VictorMont03/golang-users-app/src/model/repository"
	"github.com/VictorMont03/golang-users-app/src/model/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) users.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := services.NewUserDomainService(repo)
	userController := users.NewUserControllerInterface(service)

	return userController
}
