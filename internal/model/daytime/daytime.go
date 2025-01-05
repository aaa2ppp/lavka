package daytime

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lavka/internal/lib"
	"strconv"
)

type DayTime uint16

func New(hours, minutes int) DayTime {
	t := (hours*60 + minutes) % (24 * 60)
	if t < 0 {
		t += 24 * 60
	}
	return DayTime(t)
}

func (t DayTime) Hours() int {
	return int(t / 60)
}

func (t DayTime) Minutes() int {
	return int(t % 60)
}

type Period struct {
	Start  DayTime `json:"-"`
	Finish DayTime `json:"-"`
}

func ParsePeriod(s string) (Period, error) {
	start, finish, err := lib.ParseDayTimePeriod(s)
	if err != nil {
		return Period{}, err
	}
	return Period{DayTime(start), DayTime(finish)}, nil
}

func (p Period) String() string {
	return fmt.Sprintf("%02d:%02d-%02d:%02d", p.Start.Hours(), p.Start.Minutes(),
		p.Finish.Hours(), p.Finish.Minutes())
}

// MarshalJSON implements json.Marshaler.
func (p Period) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, 13))
	_, err := fmt.Fprintf(b, `"%02d:%02d-%02d:%02d"`, p.Start.Hours(), p.Start.Minutes(),
		p.Finish.Hours(), p.Finish.Minutes())
	return b.Bytes(), err
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *Period) UnmarshalJSON(b []byte) error {

	s := lib.UnsafeString(b)
	if s == "null" {
		return nil // no-op
	}

	s, err := strconv.Unquote(s)
	if err != nil {
		return err
	}

	start, finish, err := lib.ParseDayTimePeriod(s)
	if err != nil {
		return err
	}

	*p = Period{DayTime(start), DayTime(finish)}
	return err
}

var _ json.Marshaler = Period{}
var _ json.Unmarshaler = &Period{}
