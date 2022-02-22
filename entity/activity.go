package entity

import "time"

type Activity struct {
	Id          int       `json:"id" form:"id"`
	Category    string    `json:"category" form:"category"`
	AssetName   string    `json:"asset_name" form:"asset_name"`
	AssetImage  string    `json:"asset_image" form:"asset_image"`
	UserName    string    `json:"user_name" form:"user_name"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"note" form:"note"`
}

type ActivitySimplified struct {
	Id          int       `json:"id" form:"id"`
	AssetImage  string    `json:"asset_image" form:"asset_image"`
	AssetName   string    `json:"asset_name" form:"asset_name"`
	Status      string    `json:"status" form:"status"`
	RequestDate time.Time `json:"request_date" form:"request_date"`
}
