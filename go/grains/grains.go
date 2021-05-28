package grains

import (
	"errors"
	"math"
)

//Square calculates the total number of grains on given square
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("invalid input")
	}
	return grainsOnSquare(input), nil
}

//Total calculates the total number of grains on chessboard
func Total() uint64 {
	return (grainsOnSquare(64) * 2) - 1
}

func grainsOnSquare(input int) uint64 {
	return uint64(math.Pow(float64(2), float64(input-1)))
}
