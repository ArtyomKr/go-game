package utils

import "math"

const GridSize = 10

// SnapToGrid is a helper function to snap a value to grid
func SnapToGrid(value, gridSize float32) float32 {
	return (float32(math.Floor(float64(value/gridSize))) * gridSize) + (gridSize / 2)
}
