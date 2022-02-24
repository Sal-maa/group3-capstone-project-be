package entity

import "time"

type Asset struct {
	Id         int    `json:"id" form:"id"`
	CodeAsset  string `json:"code_asset" form:"code_asset"`
	CategoryId int    `json:"category_id" form:"category_id"`
	Image      string `json:"image" form:"image"`
	Name       string `json:"name" form:"name"`
	Short_Name string `json:"short_name" form:"short_name"`

	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
	Quantity    int       `json:"quantity" form:"quantity"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
	Category    Category
}

type CreateAsset struct {
	Image       string `json:"image" form:"image"`
	CodeAsset   string `json:"code_asset" form:"code_asset"`
	Name        string `json:"name" form:"name"`
	Short_Name  string `json:"short_name" form:"short_name"`
	Status      string `json:"status" form:"status"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type UpdateAsset struct {
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Short_Name  string `json:"short_name" form:"short_name"`
	Status      string `json:"status" form:"status"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type TotalData struct {
	TotalPage int `json:"total_page" form:"total_page"`
}
type AssetSimplified struct {
	Id             int    `json:"id" form:"id"`
	CodeAsset      string `json:"code_asset" form:"code_asset"`
	Image          string `json:"image" form:"image"`
	Name           string `json:"name" form:"name"`
	Short_Name     string `json:"short_name" form:"short_name"`
	Status         string `json:"status" form:"status"`
	Description    string `json:"description" form:"description"`
	Quantity       int    `json:"quantity" form:"quantity"`
	UserCount      int    `json:"user_count" form:"user_count"`
	StockAvailable int    `json:"stock_available" form:"stock_available"`
	CategoryName   string `json:"category_name" form:"category_name"`
	TotalData      TotalData
}
