package grains

import "errors"

const totalChessSquares = 64

//Square calculates the total number of grains on given square
func Square(input int) (uint64, error) {
	if input <= 0 || input > totalChessSquares {
		return 0, errors.New("invalid input")
	}
	return countOnSquare(input), nil
}

//Total calculates the total number of grains on chessboard
func Total() uint64 {
	return (countOnSquare(totalChessSquares + 1)) - 1
}

func countOnSquare(input int) uint64 {
	return 1 << (input - 1)
}
