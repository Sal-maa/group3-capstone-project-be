package payment

import "database/sql"

type PaymentRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}
