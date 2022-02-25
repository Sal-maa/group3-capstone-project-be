package entity

import "time"

type CreateBorrow struct {
	EmployeeName string    `json:"employee_name" form:"employee_name"`
	Category     string    `json:"category" form:"category"`
	AssetName    int       `json:"asset_name" form:"asset_name"`
	Activity     string    `json:"activity" form:"activity"`
	RequestTime  time.Time `json:"request_time" form:"request_time"`
	ReturnTime   time.Time `json:"return_time" form:"return_time"`
	Description  string    `json:"description" form:"description"`
}

type UpdateBorrow struct {
	Status string `json:"status" form:"status"`
}

type Procure struct {
	Id          int `json:"id" form:"id"`
	User        User
	Category    int       `json:"category" form:"category"`
	Image       string    `json:"image" form:"image"`
	Activity    string    `json:"activity" form:"activity"`
	RequestTime time.Time `json:"request_time" form:"request_time"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}

type CreateProcure struct {
	Category    string    `json:"category" form:"category"`
	Image       string    `json:"image" form:"image"`
	Activity    string    `json:"activity" form:"activity"`
	RequestTime time.Time `json:"request_time" form:"request_time"`
	Description string    `json:"description" form:"description"`
}

type UpdateProcure struct {
	Status string `json:"status" form:"status"`
}

type RequestResponse struct {
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
