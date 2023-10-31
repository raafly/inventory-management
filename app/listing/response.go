package listing

func ToUserResponse(user User) UserResponse {
	return UserResponse{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
	}
}

func ToUserResponses(user []User) []UserResponse {
	var userResponse []UserResponse
	for _, users := range user {
		userResponse = append(userResponse, ToUserResponse(users))
	}

	return userResponse
}


func ToItemResponse(item Item) ItemResponse {
	return ItemResponse{
		Id: item.Id,
		Name: item.Name,
		Category: item.Category,
		Quantity: item.Quantity,
	}
}

func ToItemResponses(item []Item) []ItemResponse {
	var itemResponse []ItemResponse
	for _, items := range item {
		itemResponse = append(itemResponse, ToItemResponse(items))
	}

	return itemResponse
}