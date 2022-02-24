package entity

import "time"

type Borrow struct {
	Id          int `json:"id" form:"id"`
	User        User
	Asset       Asset
	Activity    string    `json:"activity" form:"activity"`
	RequestTime time.Time `json:"request_time" form:"request_time"`
	ReturnTime  time.Time `json:"return_time" form:"return_time"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}
