package view

import (
	"github.com/VictorMont03/golang-users-app/src/controller/model/response"
	"github.com/VictorMont03/golang-users-app/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

type UserDomainResponse struct {
	ID    string
	Email string
	Name  string
	Age   int
}
