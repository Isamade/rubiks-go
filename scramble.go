package main

import (
	"math/rand/v2"
)

func scrambleCube(pieces []Piece, movesCount int) ([]string, []Piece) {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	possibleMoves := [12]string{"U", "U'", "D", "D'", "R", "R'", "L", "L'", "F", "F'", "B", "B'"}

	rotationSequence := make([]string, movesCount)

	for i := 0; i < movesCount; i++ {
		index := rand.IntN(12)
		rotationSequence[i] = possibleMoves[index]
		switch possibleMoves[index] {
		case "U":
			newState = TopClockwise(newState)
		case "U'":
			newState = TopCounterClockwise(newState)
		case "D":
			newState = BottomClockwise(newState)
		case "D'":
			newState = BottomCounterClockwise(newState)
		case "R":
			newState = RightClockwise(newState)
		case "R'":
			newState = RightCounterClockwise(newState)
		case "L":
			newState = LeftClockwise(newState)
		case "L'":
			newState = LeftCounterClockwise(newState)
		case "F":
			newState = FrontClockwise(newState)
		case "F'":
			newState = FrontCounterClockwise(newState)
		case "B":
			newState = BackClockwise(newState)
		case "B'":
			newState = BackCounterClockwise(newState)
		}
	}
	return (rotationSequence, newState)
}
