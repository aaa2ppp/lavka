package orderRepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"lavka/internal/model"
)

func (r OrderRepo) GetOrder(ctx context.Context, orderID uint64) (zero model.OrderDto, _ error) {
	x := newHelper(ctx, "GetOrder")

	const q = `SELECT weight, regions, delivery_hours, cost, completed_time` +
		` FROM "order" WHERE order_id=$1`

	var (
		weight        float64
		regions       int
		deliveryHours string
		cost          int
		completedTime time.Time
	)

	if err := r.db.QueryRowContext(ctx, q, orderID).
		Scan(&weight, &regions, &deliveryHours, &cost, &completedTime); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			x.Log().Debug("order not found", "error", err, "orderID", orderID)
			return zero, model.ErrNotFound
		}

		x.Log().Error("can't query", "error", err, "q", q, "orderID", orderID)
		return zero, model.ErrInternalError
	}

	return x.makeOrderDto(int64(orderID), weight, regions, deliveryHours, cost, completedTime)
}
