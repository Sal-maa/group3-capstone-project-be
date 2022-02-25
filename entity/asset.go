package entity

import "time"

type Asset struct {
	Id          int       `json:"id" form:"id"`
	Code        string    `json:"code" form:"code"`
	CategoryId  int       `json:"category_id" form:"category_id"`
	Name        string    `json:"name" form:"name"`
	Image       string    `json:"image" form:"image"`
	ShortName   string    `json:"short_name" form:"short_name"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}

type CreateAsset struct {
	Name             string `json:"name" form:"name"`
	Category         string `json:"category" form:"category"`
	Description      string `json:"description" form:"description"`
	Quantity         int    `json:"quantity" form:"quantity"`
	UnderMaintenance bool   `json:"under_maintenance" form:"under_maintenance"`
}

type UpdateAsset struct {
	UnderMaintenance bool `json:"under_maintenance" form:"under_maintenance"`
}

type AssetSimplified struct {
	Category       string `json:"category" form:"category"`
	Name           string `json:"name" form:"name"`
	Image          string `json:"image" form:"image"`
	ShortName      string `json:"short_name" form:"short_name"`
	Description    string `json:"description" form:"description"`
	UserCount      int    `json:"user_count" form:"user_count"`
	StockAvailable int    `json:"stock_available" form:"stock_available"`
}

type Statistics struct {
	TotalAsset       int `json:"total_asset" form:"total_asset"`
	UnderMaintenance int `json:"under_maintenance" form:"under_maintenance"`
	Borrowed         int `json:"borrowed" form:"borrowed"`
	Available        int `json:"available" form:"available"`
}
