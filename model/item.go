package model

import "time"

type ItemCreate struct {
	Id			int			`json:"id" validate:"required"`
	Name		string		`json:"name" validate:"required"`
	Quantity	int			`json:"quantiy" validate:"required"`
}

type ItemUpdate struct {
	Id			int			`json:"id" validate:"required"`
	Quantity	int			`json:"quantiy" validate:"required"`
}

type ItemResponses struct {
	Id			int				`json:"id"`
	Name		string			`json:"name"`
	Description	string			`json:"description"`
	Quantity	int				`json:"quantiy"`	
	In 			time.Time		`json:"in"`
	Out			time.Time		`json:"out"`
	Created_at	time.Time		`json:"created_at"`
}

type ItemResponse struct {
	Id			int				`json:"id"`
	Name		string			`json:"name"`
	Quantity	int				`json:"quantity"`
}