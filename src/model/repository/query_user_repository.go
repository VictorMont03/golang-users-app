package repository

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity/converter"
	"github.com/VictorMont03/golang-users-app/src/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) QueryUser(params url.Values) ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init query user repository")

	collectionName := os.Getenv("MONGODB_COLLECTION_USERS_NAME")
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	var userList []model.UserDomainInterface

	builder := utils.NewQueryBuilder().ApplyWhere(params)
	query := builder.Build()

	cursor, err := collection.Find(context.Background(), query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Users not found with"
			logger.Error(errorMessage, err, zap.String("journey", "repository QueryUser"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		} else {
			logger.Error("Error getting user Repository/Entity", err, zap.String("journey", "repository QueryUser"))
			return nil, rest_err.NewInternalServerError("Error trying to get user by id")
		}
	}

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&userEntity)

		log.Println(userEntity)

		if err != nil {
			errorMessage := "Error trying to decode user entity"
			logger.Error(errorMessage, err, zap.String("journey", "repository QueryUser"))
			return nil, rest_err.NewInternalServerError(errorMessage)
		}

		userList = append(userList, converter.ConvertEntityToDomain(*userEntity))
	}

	logger.Info("End get user repository", zap.String("journey", "repository QueryUser"))

	return userList, nil
}
