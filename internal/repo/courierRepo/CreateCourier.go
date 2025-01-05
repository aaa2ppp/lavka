package courierRepo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"lavka/internal/model"
)

// CreateCourier implements courierController.Service.
func (r CourierRepo) CreateCourier(ctx context.Context, req []model.CreateCourierDto) ([]model.CourierDto, error) {
	x := newHelper(ctx, "CreateCourier")

	q := `INSERT INTO courier (courier_id, courier_type, regions, working_hours) VALUES (%s)`

	resp := make([]model.CourierDto, 0, len(req))

	for _, it := range req {
		resp = append(resp, model.CourierDto{
			CourierID:    model.NewID(),
			CourierType:  it.CourierType,
			Regions:      it.Regions,
			WorkingHours: it.WorkingHours,
		})
	}

	var (
		idx          int
		placeHolders []string
		values       []any
	)

	for _, it := range resp {
		placeHolders = append(placeHolders, fmt.Sprintf("$%d,$%d,$%d,$%d", idx+1, idx+2, idx+3, idx+4))
		idx += 4

		var (
			workingHours []string
			regions      []string
		)

		for _, v := range it.WorkingHours {
			workingHours = append(workingHours, v.String())
		}

		for _, v := range it.Regions {
			regions = append(regions, strconv.Itoa(v))
		}

		values = append(values,
			it.CourierID,
			it.CourierType,
			strings.Join(regions, ","),
			strings.Join(workingHours, ","),
		)
	}

	q = fmt.Sprintf(q, strings.Join(placeHolders, "),("))

	if _, err := r.db.ExecContext(ctx, q, values...); err != nil {
		x.Log().Error("can't query", "error", err, "q", q, "values", values)
		return nil, model.ErrInternalError
	}

	return resp, nil
}
