package entity

import "time"

type ActivitySimplified struct {
	Id          int       `json:"id" form:"id"`
	Image       string    `json:"image" form:"image"`
	Name        string    `json:"name" form:"name"`
	Status      string    `json:"status" form:"status"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
}
