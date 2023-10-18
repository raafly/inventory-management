package model

type UserSignUp struct {
	Id			int		`json:"id" validate:"required"`
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
	Id			int		`json:"id" validate:"required"`
	Username	string	`json:"username" validate:"required"`
	Email		string	`json:"email" validate:"required"`
}