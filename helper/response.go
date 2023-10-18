package helper

import (
	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/model"
)

func ToUserResponse(user entity.User) model.UserResponse {
	return model.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
	}
}

func ToUserResponses(user []entity.User) []model.UserResponse {
	var userResponse []model.UserResponse
	for _, users := range user {
		userResponse = append(userResponse, ToUserResponse(users))
	}

	return userResponse
}


func ToItemResponse(item entity.Item) model.ItemResponse {
	return model.ItemResponse{
		Id: item.Id,
		Name: item.Name,
		Quantity: item.Quantity,
	}
}

func ToItemResponses(item []entity.Item) []model.ItemResponse {
	var itemResponse []model.ItemResponse
	for _, items := range item {
		itemResponse = append(itemResponse, ToItemResponse(items))
	}

	return itemResponse
}