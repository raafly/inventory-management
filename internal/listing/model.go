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
	Cpassword	string	`json:"confirm_password" validate:"required,min=8,eqfield=Password"`
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
	Id				string
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
	Category	string			`json:"category"`
	Quantity	int				`json:"quantity"`
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
	Status		string
	Data		interface{}
}

