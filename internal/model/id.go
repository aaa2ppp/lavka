package model

import (
	"math"
	"math/rand/v2"
)

type ID = int64

func NewID() ID {
	return ID(rand.Int64N(math.MaxInt64) + 1)
}
