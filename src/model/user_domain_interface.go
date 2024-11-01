package model

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetEmail() string
	GetAge() int8
	GetID() string

	EncryptPassword()
	SetID(string)
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
