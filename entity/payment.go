package entity

import "time"

type Payment struct {
	EwalletType string
	ExternalId  string
	Amount      float64
	BusinessId  string
	Created     string
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" form:"deleted_at"`
}

type PaymentReq struct {
	Amount float64
	Phone  string
}
