package listing

import "time"

type User struct {
	Id			string
	Username	string
	Email		string
	Password	string
	Cpassword	string
	Created_at	string
}

type Item struct {
	Id				int
	Name			string
	Description		string
	Category		string
	Quantity		int
	Status 			bool
	Created_at		time.Time
}	

type Category struct {
	Id				string
	Name			string
	Description		string
}

type History struct {
	Id			int
	ItemId		int
	Action		bool
	Quantity	int
	UpdatedAt	time.Time
}

type WebResponse struct {
	Code		int
	Data		interface{}
}

type ErrorResponse struct {
	Code	int
	Message	interface{}
}