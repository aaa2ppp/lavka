package courier

import (
	"encoding/json"
	"fmt"

	"lavka/internal/lib"
)

type Type int

const (
	_ Type = iota
	FootCourier
	BikeCourier
	AutoCourier
)

const (
	s_foot = "FOOT"
	s_bike = "BIKE"
	s_auto = "AUTO"
	q_foot = `"FOOT"`
	q_bike = `"BIKE"`
	q_auto = `"AUTO"`
)

func (t Type) IsZero() bool {
	return t == 0
}

func ParseCourierType(s string) (Type, error) {
	switch s {
	case s_foot:
		return FootCourier, nil
	case s_bike:
		return BikeCourier, nil
	case s_auto:
		return AutoCourier, nil
	default:
		return 0, fmt.Errorf("%s unknown courier type", s)
	}
}

func (t Type) String() string {
	switch t {
	case FootCourier:
		return s_foot
	case BikeCourier:
		return s_bike
	case AutoCourier:
		return s_auto
	}
	return fmt.Sprintf("CourierType(%d)", t)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Type) UnmarshalJSON(b []byte) error {
	switch lib.UnsafeString(b) {
	case `null`:
	case q_foot:
		*t = FootCourier
	case q_bike:
		*t = BikeCourier
	case q_auto:
		*t = AutoCourier
	default:
		return fmt.Errorf("%s unknown courier type", b)
	}
	return nil
}

// MarshalJSON implements json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	switch t {
	case 0:
		return []byte(`null`), nil
	case FootCourier:
		return []byte(q_foot), nil
	case BikeCourier:
		return []byte(q_bike), nil
	case AutoCourier:
		return []byte(q_auto), nil
	}
	return nil, fmt.Errorf("invalid value %v", t)
}

var _ json.Unmarshaler = (*Type)(nil)
var _ json.Marshaler = Type(0)
