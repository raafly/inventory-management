package model

import "time"

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