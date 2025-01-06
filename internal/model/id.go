package model

import (
	"math/rand/v2"
)

// Максимальное целое число, которое можно безопасно использовать в JavaScript (2^53 - 1)
const MaxID = (1 << 53) - 1

type ID = int64

func NewID() ID {
	return ID(rand.Int64N(MaxID) + 1)
}
