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

// category 

type Category struct {
	Id				string
	Name			string
	Description		string
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