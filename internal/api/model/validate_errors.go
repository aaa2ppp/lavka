package model

import "fmt"

func ErrCannotBeEmpty(name string) error {
	return fmt.Errorf("%s cannot be empty", name)
}

func ErrMustBeGreaterThanZero(name string) error {
	return fmt.Errorf("%s must be greater than 0", name)
}

func ErrItemMustBeGreaterThanZero(name string, i int) error {
	return fmt.Errorf("%s[%d] must be greater than 0", name, i)
}

func ErrMustBeOnOf(name string, enum []string) error {
	return fmt.Errorf("%s must be one of %v", name, enum)
}

func ErrMustBeDateTime(name string) error {
	return fmt.Errorf("%s must be date-time in RFC3339 format", name)
}

func ErrMustBeDayTimePeriod(name string) error {
	return fmt.Errorf("%s must be time period in HH:MM-HH:MM format", name)
}

func ErrItemMustBeDayTimePeriod(name string, i int) error {
	return fmt.Errorf("%s[%d] must be time period in HH:MM-HH:MM format", name, i)
}

func ErrItemError(name string, i int, err error) error {
	return fmt.Errorf("%s[%d] %w", name, i, err)
}
