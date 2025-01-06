package orderRepo

import (
	"context"
	"database/sql"
	"strings"

	"lavka/internal/model"
	"lavka/internal/model/daytime"
	"lavka/internal/repo/helper"
)

type xhelper struct{ *helper.Helper }

func newHelper(ctx context.Context, op string) xhelper {
	return xhelper{helper.New(ctx, "courierRepo", op)}
}

func (x xhelper) makeOrderDto(
	orderID int64,
	weight float64,
	regions int,
	deliveryHours string,
	cost int,
	completedTime sql.NullTime,
) (zero model.OrderDto, _ error) {

	resp := model.OrderDto{
		OrderID: orderID,
		Weight:  weight,
		Regions: regions,
		Cost:    cost,
	}

	if completedTime.Valid {
		resp.CompletedTime = model.NullTime{Time: completedTime.Time}
	}

	for _, s := range strings.Split(deliveryHours, ",") {
		period, err := daytime.ParsePeriod(s)
		if err != nil {
			x.Log().Error("can't parse delivery_hours", "error", err, "s", s)
			return zero, model.ErrInternalError
		}
		resp.DeliveryHours = append(resp.DeliveryHours, period)
	}

	return resp, nil
}
