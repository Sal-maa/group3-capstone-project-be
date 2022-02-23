package entity

import "time"

type Activity struct {
	Id           int       `json:"id" form:"id"`
	Category     string    `json:"category" form:"category"`
	AssetName    string    `json:"asset_name" form:"asset_name"`
	AssetImage   string    `json:"asset_image" form:"asset_image"`
	StockLeft    int       `json:"stock_left" form:"stock_left"`
	UserName     string    `json:"user_name" form:"user_name"`
	RequestDate  time.Time `json:"request_date" form:"request_date"`
	ReturnDate   time.Time `json:"return_date" form:"return_date"`
	ActivityType string    `json:"activity_type" form:"activity_type"`
	Status       string    `json:"status" form:"status"`
	Description  string    `json:"note" form:"note"`
}

type ActivitySimplified struct {
	Id          int       `json:"id" form:"id"`
	AssetImage  string    `json:"asset_image" form:"asset_image"`
	AssetName   string    `json:"asset_name" form:"asset_name"`
	Status      string    `json:"status" form:"status"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
}

type UpdateActivityStatus struct {
	Status string `json:"status" form:"status"`
}
