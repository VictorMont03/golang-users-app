package model

type userDomain struct {
	ID       string
	name     string
	password string
	email    string
	age      int8
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}
