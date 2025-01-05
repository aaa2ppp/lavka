package courierRepo

import (
	"database/sql"
)

type CourierRepo struct {
	db *sql.DB
}

func New(db *sql.DB) CourierRepo {
	return CourierRepo{db}
}
