package entity

import "time"

type Item struct {
	Id				int
	Name			string
	Description		string
	Quantity		int
	In				time.Time
	Out				time.Time
	Created_at		time.Time
}