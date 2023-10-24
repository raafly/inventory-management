package model

type UserSignUp struct {
	Username	string	`json:"username" validate:"required"`
	Email		string	`json:"email" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}

type UserSignIn struct {
	Email		string	`json:"email" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}

type UserUpdate struct {
	Email		string	`json:"email" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}

type UserResponse struct {
	Id			string	`json:"id" validate:"required"`
	Username	string	`json:"username" validate:"required"`
	Email		string	`json:"email" validate:"required"`
}