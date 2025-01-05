package courierRepo

import (
	"context"
	"strconv"
	"strings"

	"lavka/internal/model"
	"lavka/internal/model/courier"
	"lavka/internal/model/daytime"
	"lavka/internal/repo/helper"
)

type xhelper struct{ *helper.Helper }

func newHelper(ctx context.Context, op string) xhelper {
	return xhelper{helper.New(ctx, "courierRepo", op)}
}

// makeCourierDto make and returns CourierDto structure or zero value and error
func (x xhelper) makeCourierDto(
	courierID int64,
	courierType,
	regions,
	workingHours string,
) (zero model.CourierDto, _ error) {

	resp := model.CourierDto{CourierID: courierID}

	{
		v, err := courier.ParseType(courierType)
		if err != nil {
			x.Log().Error("can't parse courier_type", "error", err, "s", courierType)
			return zero, model.ErrInternalError
		}
		resp.CourierType = v
	}

	for _, s := range strings.Split(regions, ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			x.Log().Error("can't parse regions", "error", err, "s", s)
			return zero, model.ErrInternalError
		}
		resp.Regions = append(resp.Regions, v)
	}

	for _, s := range strings.Split(workingHours, ",") {
		v, err := daytime.ParsePeriod(s)
		if err != nil {
			x.Log().Error("can't parse working_hours", "error", err, "s", s)
			return zero, model.ErrInternalError
		}
		resp.WorkingHours = append(resp.WorkingHours, v)
	}

	return resp, nil
}
