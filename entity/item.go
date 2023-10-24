package entity

import "time"

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