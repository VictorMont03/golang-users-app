package converter

import (
	"github.com/VictorMont03/golang-users-app/src/model"
	"github.com/VictorMont03/golang-users-app/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
