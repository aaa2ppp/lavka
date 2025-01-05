package lib

import (
	"fmt"
	"strconv"
)

func ParseDayTimePeriod(s string) (start, finish int, err error) {
	const op = "ParseDayTimePeriod"

	if len(s) != 11 && s[5] != '-' {
		return 0, 0, fmt.Errorf("%s: %q want HH:MM-HH:MM", op, s)
	}

	start, err = ParseDayTime(s[:5])
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %q start: %v", op, s, err)
	}

	finish, err = ParseDayTime(s[6:])
	if err != nil {
		return 0, 0, fmt.Errorf("%s: %q finish: %v", op, s, err)
	}

	// (!) Это не проблема парсинга. В общем случае периоды могут пересекать полночь
	// if start > finish {
	// 	return 0, 0, fmt.Errorf("%s: %q start must be earlier finish", op, s)
	// }

	return start, finish, nil
}

func ParseDayTime(s string) (int, error) {
	const op = "ParseDayTime"

	if len(s) != 5 || s[2] != ':' {
		return 0, fmt.Errorf("%s: %q want HH:MM", op, s)
	}

	hh, err := strconv.ParseUint(s[:2], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %q hours must be unsigned integer", op, s)
	}
	if !(0 <= hh && hh <= 23) {
		return 0, fmt.Errorf("%s: %q hours must be in [0..23]", op, s)
	}

	mm, err := strconv.ParseUint(s[3:], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %q minutes must be unsigned integer", op, s)
	}
	if !(0 <= mm && mm <= 59) {
		return 0, fmt.Errorf("%s: %q minutes must be in [0..59]", op, s)
	}

	return int(hh*60 + mm), nil
}
