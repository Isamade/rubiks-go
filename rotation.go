package main

// Piece represents a single piece on the Rubik's cube with its colors
type Piece struct {
	Colors   []string `json:"colors"`
	Position []int    `json:"position"`
}

// pieces represents the state of the entire cube

// changeColors is a helper function to rotate colors on a piece
func changeColorsTopClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[5]
	result[5] = result[0]
	result[0] = result[4]
	result[4] = result[1]
	result[1] = temp

	return result
}

// TopClockwise rotates the top face of the cube clockwise
func TopClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on top face
	locationsMatrix := make([][]int, 0)
	for i := -3; i < 0; i++ {
		locationsVector := make([]int, 0)
		for j := 1; j < 4; j++ {
			locationsVector = append(locationsVector, 9*j+i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsTopClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsTopCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[5]
	result[5] = result[1]
	result[1] = result[4]
	result[4] = result[0]
	result[0] = temp

	return result
}

// TopCounterClockwise rotates the top face of the cube counter clockwise
func TopCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on top face
	locationsMatrix := make([][]int, 0)
	for i := -3; i < 0; i++ {
		locationsVector := make([]int, 0)
		for j := 1; j < 4; j++ {
			locationsVector = append(locationsVector, 9*j+i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsTopCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsBottomClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[5]
	result[5] = result[1]
	result[1] = result[4]
	result[4] = result[0]
	result[0] = temp

	return result
}

// BottomClockwise rotates the bottom face of the cube clockwise
func BottomClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on bottom face
	locationsMatrix := make([][]int, 0)
	for i := 2; i > -1; i-- {
		locationsVector := make([]int, 0)
		for j := 0; j < 3; j++ {
			locationsVector = append(locationsVector, 9*j+i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsBottomClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsBottomCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[5]
	result[5] = result[0]
	result[0] = result[4]
	result[4] = result[1]
	result[1] = temp

	return result
}

// BottomCounterClockwise rotates the bottom face of the cube counter clockwise
func BottomCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on bottom face
	locationsMatrix := make([][]int, 0)
	for i := 2; i > -1; i-- {
		locationsVector := make([]int, 0)
		for j := 0; j < 3; j++ {
			locationsVector = append(locationsVector, 9*j+i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsBottomCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsRightClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[2]
	result[2] = result[0]
	result[0] = result[3]
	result[3] = result[1]
	result[1] = temp

	return result
}

// RightClockwise rotates the right face of the cube clockwise
func RightClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on right face
	locationsMatrix := make([][]int, 0)
	for i := 0; i > -3; i-- {
		locationsVector := make([]int, 0)
		for j := -1; j > -4; j-- {
			locationsVector = append(locationsVector, 3*(9+i)+j)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsRightClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsRightCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[2]
	result[2] = result[1]
	result[1] = result[3]
	result[3] = result[0]
	result[0] = temp

	return result
}

// RightCounterClockwise rotates the right face of the cube counter clockwise
func RightCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on right face
	locationsMatrix := make([][]int, 0)
	for i := 0; i > -3; i-- {
		locationsVector := make([]int, 0)
		for j := -1; j > -4; j-- {
			locationsVector = append(locationsVector, 3*(9+i)+j)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsRightCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsLeftClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[2]
	result[2] = result[1]
	result[1] = result[3]
	result[3] = result[0]
	result[0] = temp

	return result
}

// LeftClockwise rotates the left face of the cube clockwise
func LeftClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on left face
	locationsMatrix := make([][]int, 0)
	for i := 2; i > -1; i-- {
		locationsVector := make([]int, 0)
		for j := 0; j < 3; j++ {
			locationsVector = append(locationsVector, 3*i+j)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsLeftClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsLeftCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[2]
	result[2] = result[0]
	result[0] = result[3]
	result[3] = result[1]
	result[1] = temp

	return result
}

// LeftCounterClockwise rotates the left face of the cube counter clockwise
func LeftCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on left face
	locationsMatrix := make([][]int, 0)
	for i := 2; i > -1; i-- {
		locationsVector := make([]int, 0)
		for j := 0; j < 3; j++ {
			locationsVector = append(locationsVector, 3*i+j)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsLeftCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsFrontClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[4]
	result[4] = result[2]
	result[2] = result[5]
	result[5] = result[3]
	result[3] = temp

	return result
}

// FrontClockwise rotates the front face of the cube clockwise
func FrontClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on front face
	locationsMatrix := make([][]int, 0)
	for i := 0; i > -3; i-- {
		locationsVector := make([]int, 0)
		for j := 1; j < 4; j++ {
			locationsVector = append(locationsVector, 9*j + 3*i -1)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsFrontClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsFrontCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[4]
	result[4] = result[3]
	result[3] = result[5]
	result[5] = result[2]
	result[2] = temp

	return result
}

// FrontClockwise rotates the front face of the cube clockwise
func FrontCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on front face
	locationsMatrix := make([][]int, 0)
	for i := 0; i > -3; i-- {
		locationsVector := make([]int, 0)
		for j := 1; j < 4; j++ {
			locationsVector = append(locationsVector, 9*j + 3*i -1)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsFrontCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsBackClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[4]
	result[4] = result[3]
	result[3] = result[5]
	result[5] = result[2]
	result[2] = temp

	return result
}

// BackClockwise rotates the front face of the cube clockwise
func BackClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on front face
	locationsMatrix := make([][]int, 0)
	for i := -1; i > -4; i-- {
		locationsVector := make([]int, 0)
		for j := 3; j > 0; j-- {
			locationsVector = append(locationsVector, (9*j)+3*i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third columns
	for i := 0; i < 3; i++ {
		rotationMatrix[i][0], rotationMatrix[i][2] = rotationMatrix[i][2], rotationMatrix[i][0]
	}

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsBackClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}


// changeColors is a helper function to rotate colors on a piece
func changeColorsBackCounterClockwise(colors []string) []string {
	result := make([]string, len(colors))
	copy(result, colors)

	temp := result[4]
	result[4] = result[2]
	result[2] = result[5]
	result[5] = result[3]
	result[3] = temp

	return result
}

// BackClockwise rotates the front face of the cube clockwise
func BackCounterClockwise(pieces []Piece) []Piece {
	// Create a deep copy of pieces
	newState := make([]Piece, len(pieces))
	for k, v := range pieces {
		colorsCopy := make([]string, len(v.Colors))
		copy(colorsCopy, v.Colors)
		positionCopy := make([]int, len(v.Position))
		copy(positionCopy, v.Position)
		newState[k] = Piece{Colors: colorsCopy, Position: positionCopy}
	}

	// Build the locations matrix for the pieces on front face
	locationsMatrix := make([][]int, 0)
	for i := -1; i > -4; i-- {
		locationsVector := make([]int, 0)
		for j := 3; j > -1; j-- {
			locationsVector = append(locationsVector, (9*j)+3*i)
		}
		locationsMatrix = append(locationsMatrix, locationsVector)
	}

	// Create the rotation matrix by first transposing locationsMatrix
	rotationMatrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		rotationMatrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			rotationMatrix[i][j] = locationsMatrix[j][i]
		}
	}

	// Complete the rotation matrix by swapping first and third rows
	rotationMatrix[0], rotationMatrix[2] = rotationMatrix[2], rotationMatrix[0]

	// Update newState based on rotation
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Find new position of the piece at locationsMatrix[i][j]
			var newI, newJ int
			found := false
			for m := 0; m < 3 && !found; m++ {
				for n := 0; n < 3 && !found; n++ {
					if rotationMatrix[m][n] == locationsMatrix[i][j] {
						newI, newJ = m, n
						found = true
					}
				}
			}

			// Update the piece's position by changing colors in newState
			sourcePiece := pieces[locationsMatrix[i][j]]
			newState[locationsMatrix[newI][newJ]] = Piece{
				Colors:   changeColorsBackCounterClockwise(sourcePiece.Colors),
				Position: pieces[locationsMatrix[newI][newJ]].Position,
			}
		}
	}

	return newState
}
