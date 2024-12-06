package repository

import (
	"context"
	"os"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init update user repository")

	collectionName := os.Getenv("MONGODB_COLLECTION_USERS_NAME")
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		logger.Error("User id is not a valid hex", err)
		return rest_err.NewInternalServerError("User id is not a valid hex")
	}

	_, err = collection.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: userIdHex}}, bson.D{{Key: "$set", Value: value}})

	if err != nil {
		logger.Error("Error updating user Repository/Entity", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("End update user repository", zap.String("user_id", userId))

	return nil
}
