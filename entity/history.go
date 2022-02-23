package entity

import "time"

type UserRequestHistory struct {
	Id          int       `json:"id" form:"id"`
	Category    string    `json:"category" form:"category"`
	AssetName   string    `json:"asset_name" form:"asset_name"`
	AssetImage  string    `json:"asset_image" form:"asset_image"`
	StockLeft   int       `json:"stock_left" form:"stock_left"`
	UserName    string    `json:"user_name" form:"user_name"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
	ReturnDate  time.Time `json:"return_date" form:"return_date"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
}

type UserRequestHistorySimplified struct {
	Id           int       `json:"id" form:"id"`
	Category     string    `json:"category" form:"category"`
	AssetName    string    `json:"asset_name" form:"asset_name"`
	AssetImage   string    `json:"asset_image" form:"asset_image"`
	RequestDate  time.Time `json:"request_date" form:"request_date"`
	ActivityType string    `json:"activity_type" form:"activity_type"`
}

type AssetUsageHistory struct {
	Id          int       `json:"id" form:"id"`
	Category    string    `json:"category" form:"category"`
	AssetName   string    `json:"asset_name" form:"asset_name"`
	AssetImage  string    `json:"asset_image" form:"asset_image"`
	UserName    string    `json:"user_name" form:"user_name"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
	Status      string    `json:"status" form:"status"`
}
