package orderRepo

import (
	"database/sql"
)

type OrderRepo struct {
	db *sql.DB
}

func New(db *sql.DB) OrderRepo {
	return OrderRepo{db}
}
