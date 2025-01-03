package lib

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseDayTimePeriod(s string) (start, finish int, err error) {

	if len(s) != 11 && s[5] != '-' {
		return 0, 0, errors.New("want HH:MM-HH:MM")
	}

	start, err = ParseDayTime(s[:5])
	if err != nil {
		return 0, 0, fmt.Errorf("start: %v", err)
	}

	finish, err = ParseDayTime(s[6:])
	if err != nil {
		return 0, 0, fmt.Errorf("finish: %v", err)
	}

	if start > finish {
		return 0, 0, errors.New("start must be earlier finish")
	}

	return start, finish, nil
}

func ParseDayTime(s string) (int, error) {

	if len(s) != 5 || s[2] != ':' {
		return 0, errors.New("want HH:MM")
	}

	hh, err := strconv.ParseUint(s[:2], 10, 64)
	if err != nil {
		return 0, errors.New("hours must be unsigned integer")
	}
	if !(0 <= hh && hh <= 23) {
		return 0, errors.New("hours must be in [0..23]")
	}

	mm, err := strconv.ParseUint(s[3:], 10, 64)
	if err != nil {
		return 0, errors.New("minutes must be unsigned integer")
	}
	if !(0 <= mm && mm <= 59) {
		return 0, errors.New("minutes must be in [0..59]")
	}

	return int(hh*60 + mm), nil
}
