package courierRepo

import (
	"context"
	"database/sql"
	"errors"

	"lavka/internal/model"
)

// GetCourierById implements courierController.Service.
func (r CourierRepo) GetCourierById(ctx context.Context, courierID uint64) (model.CourierDto, error) {
	x := newHelper(ctx, "GetCourierById")

	const q = `SELECT courier_type, regions, working_hours FROM courier WHERE courier_id=$1`

	var (
		courierType  string
		regions      string
		workingHours string
	)

	if err := r.db.QueryRowContext(ctx, q, courierID).
		Scan(&courierType, &regions, &workingHours); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			x.Log().Debug("can't query", "error", err, "q", q, "courierID", courierID)
			return model.CourierDto{}, model.ErrNotFound
		}

		x.Log().Error("can't query", "error", err, "q", q, "courierID", courierID)
		return model.CourierDto{}, model.ErrInternalError
	}

	return x.makeCourierDto(int64(courierID), courierType, regions, workingHours)
}
