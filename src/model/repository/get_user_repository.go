package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) GetUser(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init get user repository")

	collectionName := os.Getenv("MONGODB_COLLECTION_USERS_NAME")
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	parseId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: parseId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with id %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "repository GetUser"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		} else {
			logger.Error("Error getting user Repository/Entity", err, zap.String("journey", "repository GetUser"))
			return nil, rest_err.NewInternalServerError("Error trying to get user by id")
		}
	}

	logger.Info("End get user repository", zap.String("journey", "repository GetUser"), zap.String("user_id", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
