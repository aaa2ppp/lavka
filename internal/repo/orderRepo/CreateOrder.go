package orderRepo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"lavka/internal/model"
)

func (r OrderRepo) CreateOrder(ctx context.Context, req []model.CreateOrderDto) ([]model.OrderDto, error) {
	x := newHelper(ctx, "CreateOrder")

	q := `INSERT INTO "order" (order_id, weight, regions, delivery_hours, cost, completed_time)` +
		` VALUES (%s)`

	resp := make([]model.OrderDto, 0, len(req))

	for _, it := range req {
		resp = append(resp, model.OrderDto{
			OrderID:       model.NewID(),
			Weight:        it.Weight,
			Regions:       it.Regions,
			DeliveryHours: it.DeliveryHours,
			Cost:          it.Cost,
		})
	}

	var (
		idx          int
		placeHolders []string
		values       []any
	)

	for _, it := range resp {
		placeHolders = append(placeHolders, fmt.Sprintf("$%d,$%d,$%d,$%d,$%d,$%d",
			idx+1, idx+2, idx+3, idx+4, idx+5, idx+6))
		idx += 6

		var (
			deliveryHours []string
		)

		for _, v := range it.DeliveryHours {
			deliveryHours = append(deliveryHours, v.String())
		}

		values = append(values,
			it.OrderID,
			it.Weight,
			it.Regions,
			strings.Join(deliveryHours, ","),
			it.Cost,
			time.Time{}, // zero time
		)
	}

	q = fmt.Sprintf(q, strings.Join(placeHolders, "),("))

	if _, err := r.db.ExecContext(ctx, q, values...); err != nil {
		x.Log().Error("can't query", "error", err, "q", q, "values", values)
		return nil, model.ErrInternalError
	}

	return resp, nil
}
