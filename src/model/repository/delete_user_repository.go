package repository

import (
	"context"
	"os"

	"github.com/VictorMont03/golang-users-app/src/config/logger"
	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init delete user repository")

	collectionName := os.Getenv("MONGODB_COLLECTION_USERS_NAME")
	collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		logger.Error("User id is not a valid hex", err)
		return rest_err.NewInternalServerError("User id is not a valid hex")
	}

	_, err = collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: userIdHex}})

	if err != nil {
		logger.Error("Error updating user Repository/Entity", err)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("End update user repository", zap.String("user_id", userId))

	return nil
}
