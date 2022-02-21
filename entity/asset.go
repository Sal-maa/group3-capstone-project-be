package entity

import "time"

type Asset struct {
	Id          int       `json:"id" form:"id"`
	CategoryId  int       `json:"category_id" form:"category_id"`
	Image       string    `json:"image" form:"image"`
	Name        string    `json:"name" form:"name"`
	Status      string    `json:"status" form:"status"`
	Description string    `json:"description" form:"description"`
	Quantity    int       `json:"quantity" form:"quantity"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}

type CreateAsset struct {
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Status      string `json:"status" form:"status"`
	Address     string `json:"address" form:"address"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type UpdateAsset struct {
	Image       string `json:"image" form:"image"`
	Name        string `json:"name" form:"name"`
	Status      string `json:"status" form:"status"`
	Address     string `json:"address" form:"address"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}
