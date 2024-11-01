package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
	Age      int8               `bson:"age"`
}
