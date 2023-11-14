package listing

import "time"

// user

type User struct {
	Id			string
	Username	string
	Email		string
	Password	string
	Cpassword	string
	Created_at	string
}

type UserSignUp struct {
	Username	string	`json:"username" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=8"`
	Cpassword	string	`json:"confirmPassword" validate:"required,min=8,eqfield=Password"`
}

type UserSignIn struct {
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=8"`
}

type UserUpdate struct {
	Username	string	`json:"username" required:"username"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	Id			string	`json:"id" validate:"required"`
	Username	string	`json:"username" validate:"required"`
	Email		string	`json:"email" validate:"required"`
}

// item

type Item struct {
	Id				int
	Name			string
	Description		string
	Category		string
	Quantity		int
	Status 			bool
	Created_at		time.Time
}	

type ItemCreate struct {
	Name		string		`json:"name" validate:"required"`
	Category	string		`json:"category" validate:"required"`
	Quantity	int			`json:"quantity" validate:"required"`
}

type ItemUpdate struct {
	Id			int			`json:"id" validate:"required"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Status		bool		`json:"status"`
	Quantity	int			`json:"quantity"`
}

type ItemResponse struct {
	Id			int			`json:"id"`
	Name		string			`json:"name"`
	Description	string			`json:"description"`
	Category	string			`json:"category"`
	Quantity	int				`json:"quantity"`
	Status		bool			`json:"status"`
	Created_at	time.Time		`json:"createdAt"`
}

// category 

type Category struct {
	Id				string
	Name			string
	Description		string
}

type CategoryCreate struct {
	Id				string	`json:"id" validate:"required"`
	Name			string	`json:"name" validate:"required"`
	Description		string	`json:"description"`
}

type CategoryUpdate struct {
	Id 				string	`json:"id" validate:"required"`
	Description		string	`json:"description" validate:"required"`
}


// web

type WebResponse struct {
	Code		int
	Data		interface{}
}

type ErrorResponse struct {
	Code	int
	Message	interface{}
}