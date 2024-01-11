package view

import (
	dto_response "github.com/vinialeixo/crud-golang/src/controller/dto/response"
	"github.com/vinialeixo/crud-golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) dto_response.UserResponse {
	return dto_response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
