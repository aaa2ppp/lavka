package orderRepo

import (
	"context"
	"database/sql"

	"lavka/internal/model"
)

func (r OrderRepo) GetOrders(ctx context.Context, limit, offset int) ([]model.OrderDto, error) {
	x := newHelper(ctx, "GetOrders")

	const q = `SELECT order_id, weight, regions, delivery_hours, cost, completed_time` +
		` FROM "order" LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		x.Log().Error("can't query", "error", err, "q", q, "limit", limit, "offset", offset)
		return nil, model.ErrInternalError
	}

	resp := make([]model.OrderDto, 0, limit)

	var (
		orderID       int64
		weight        float64
		regions       int
		deliveryHours string
		cost          int
		completedTime sql.NullTime
	)

	for rows.Next() {

		if err := rows.Scan(&orderID, &weight, &regions, &deliveryHours, &cost, &completedTime); err != nil {
			x.Log().Error("can't scan", "error", err)
			return nil, model.ErrInternalError
		}

		order, err := x.makeOrderDto(orderID, weight, regions, deliveryHours, cost, completedTime)
		if err != nil {
			return nil, err
		}

		resp = append(resp, order)
	}

	return resp, nil
}
