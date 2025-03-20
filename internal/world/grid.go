package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (w *World) DrawGrid(cellSize float32) {
	floor := w.State.Floor
	// Calculate the number of cells based on the floor's size
	width := int32(floor.Size.X / cellSize)
	height := int32(floor.Size.Y / cellSize)

	floorStart := rl.Vector3{
		X: floor.Position.X + floor.Size.X/2,
		Y: floor.Position.Y + 0.1,
		Z: floor.Position.Z + floor.Size.Y/2,
	}

	floorEnd := rl.Vector3{
		X: floor.Position.X - floor.Size.X/2,
		Y: floor.Position.Y + 0.1,
		Z: floor.Position.Z - floor.Size.Y/2,
	}

	// Draw horizontal grid lines
	for i := int32(0); i <= height; i++ {
		start := rl.Vector3{
			X: floorStart.X,
			Y: floorStart.Y,
			Z: floorStart.Z - float32(i)*cellSize,
		}
		end := rl.Vector3{
			X: floorEnd.X,
			Y: floorEnd.Y,
			Z: floorStart.Z - float32(i)*cellSize,
		}
		rl.DrawLine3D(start, end, rl.Black)
	}

	// Draw vertical grid lines
	for i := int32(0); i <= width; i++ {
		start := rl.Vector3{
			X: floorStart.X - float32(i)*cellSize,
			Y: floorStart.Y,
			Z: floorStart.Z,
		}
		end := rl.Vector3{
			X: floorStart.X - float32(i)*cellSize,
			Y: floorEnd.Y,
			Z: floorEnd.Z,
		}
		rl.DrawLine3D(start, end, rl.Black)
	}
}

func (w *World) drawAxisHelpers(length float32) {
	position := rl.Vector3{Y: 1}

	// X-axis (Red)
	rl.DrawLine3D(position, rl.Vector3{X: position.X + length, Y: position.Y, Z: position.Z}, rl.Red)

	// Y-axis (Green)
	rl.DrawLine3D(position, rl.Vector3{X: position.X, Y: position.Y + length, Z: position.Z}, rl.Green)

	// Z-axis (Blue)
	rl.DrawLine3D(position, rl.Vector3{X: position.X, Y: position.Y, Z: position.Z + length}, rl.Blue)
}
