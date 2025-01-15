package courierRepo

import (
	"context"

	"lavka/internal/model"
)

// GetCouriers implements courierController.Service.
func (r CourierRepo) GetCouriers(ctx context.Context, limit int, offset int) ([]model.CourierDto, error) {
	x := newHelper(ctx, "GetCouriers")

	const q = `SELECT courier_id, courier_type, regions, working_hours FROM courier LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		x.Log().Error("can't query", "error", err, "q", q, "limit", limit, "offset", offset)
		return nil, model.ErrInternalError
	}
	defer rows.Close()

	resp := make([]model.CourierDto, 0, limit)

	var (
		courierID    int64
		courierType  string
		regions      string
		workingHours string
	)

	for rows.Next() {
		if err := rows.Scan(&courierID, &courierType, &regions, &workingHours); err != nil {
			x.Log().Error("can't scan row", "error", err)
			return nil, model.ErrInternalError
		}

		it, err := x.makeCourierDto(courierID, courierType, regions, workingHours)
		if err != nil {
			return nil, err
		}

		resp = append(resp, it)
	}

	if err := rows.Err(); err != nil {
		x.Log().Error("can't get next row", "error", err)
		return nil, model.ErrInternalError
	}

	return resp, nil
}
