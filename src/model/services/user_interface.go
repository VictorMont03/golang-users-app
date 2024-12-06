package services

import (
	"net/url"

	"github.com/VictorMont03/golang-users-app/src/config/rest_err"
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	GetUser(string) (model.UserDomainInterface, *rest_err.RestErr)
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	QueryUser(url.Values) ([]model.UserDomainInterface, *rest_err.RestErr)
	LoginUser(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
