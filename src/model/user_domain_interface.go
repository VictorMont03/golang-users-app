package model

import "github.com/VictorMont03/golang-users-app/src/config/rest_err"

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetEmail() string
	GetAge() int8
	GetID() string

	EncryptPassword()
	SetID(string)
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email string,
	password string,
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name:     name,
		password: password,
		email:    email,
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUseLoginDomain(
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
