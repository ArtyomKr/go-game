package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func rotateCamera(camera *rl.Camera3D, angle float32) {
	// Rotate around Y axis
	dir := rl.Vector3Subtract(camera.Target, camera.Position)

	cosA := float32(math.Cos(float64(angle * rl.Deg2rad)))
	sinA := float32(math.Sin(float64(angle * rl.Deg2rad)))

	newDirX := dir.X*cosA - dir.Z*sinA
	newDirZ := dir.X*sinA + dir.Z*cosA

	dir.X = newDirX
	dir.Z = newDirZ

	camera.Target = rl.Vector3Add(camera.Position, dir)
}
