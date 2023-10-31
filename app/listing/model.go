package listing

import "time"

// user

type User struct {
	Id			string
	Username	string
	Email		string
	Password	string
	Created_at	string
}

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

// item

type Item struct {
	Id				string
	Name			string
	Description		string
	Category		int
	Quantity		int
	In				time.Time
	Out				time.Time
	Created_at		time.Time
}



type ItemCreate struct {
	Name		string		`json:"name" validate:"required"`
	Category	int			`json:"category" validate:"required"`
	Quantity	int			`json:"quantity" validate:"required"`
}

type ItemUpdate struct {
	Name		string		`json:"name" validate:"required"`
	Quantity	int			`json:"quantity" validate:"required"`
}

type ItemResponses struct {
	Id			string			`json:"id"`
	Name		string			`json:"name"`
	Description	string			`json:"description"`
	Category	int				`json:"category"`
	Quantity	int				`json:"quantity"`	
	In 			time.Time		`json:"in"`
	Out			time.Time		`json:"out"`
	Created_at	time.Time		`json:"created_at"`
}

type ItemResponse struct {
	Id			string			`json:"id"`
	Name		string			`json:"name"`
	Category	int				`json:"category"`
	Quantity	int				`json:"quantity"`
}

// web

type WebResponse struct {
	Code		int
	Status		string
	Data		interface{}
}