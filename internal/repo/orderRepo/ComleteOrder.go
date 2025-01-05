package orderRepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"lavka/internal/model"
)

func (r OrderRepo) ComleteOrder(ctx context.Context, req []model.CompleteOrderDto) ([]model.OrderDto, error) {
	x := newHelper(ctx, "ComleteOrder")

	const q = `UPDATE "order" SET completed_time=$1` +
		` FROM "order" AS o JOIN courier AS c USING(courier_id)` +
		` WHERE o.order_id=$2 AND o.courier_id=$3` +
		` RETURNING o.order_id, o.weight, o.regions, o.delivery_hours, o.cost, o.completed_time`

	stmt, err := r.db.PrepareContext(ctx, q)
	if err != nil {
		x.Log().Error("can't prepare query", "error", err, "q", q)
		return nil, model.ErrInternalError
	}
	defer stmt.Close()

	tx, err := r.db.Begin()
	if err != nil {
		x.Log().Error("can't begin transaction", "error", err)
		return nil, model.ErrInternalError
	}
	defer tx.Rollback()

	stmt = tx.StmtContext(ctx, stmt)

	resp := []model.OrderDto{}

	var (
		orderID       int64
		weight        float64
		regions       int
		deliveryHours string
		cost          int
		completedTime time.Time
	)

	for _, p := range req {

		if err := stmt.QueryRowContext(ctx, p.CompleteTime, p.OrderID, p.CourierID).
			Scan(&orderID, &weight, &regions, &deliveryHours, &cost, &completedTime); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				x.Log().Debug("order not found", "error", err, "q", q, "p", p)
				return nil, model.ErrBadRequest
			}

			x.Log().Error("can't query", "error", err, "q", q, "p", p)
			return nil, model.ErrInternalError
		}

		order, err := x.makeOrderDto(orderID, weight, regions, deliveryHours, cost, completedTime)
		if err != nil {
			return nil, err
		}

		resp = append(resp, order)
	}

	if err := tx.Commit(); err != nil {
		x.Log().Error("can't commit transaction", "error", err)
		return nil, model.ErrInternalError
	}

	return resp, nil
}
