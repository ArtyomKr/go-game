package camera

import rl "github.com/gen2brain/raylib-go/raylib"

func zoomCamera(camera *rl.Camera3D, amount float32, minZoom, maxZoom float32) {
	// Get direction from camera to target
	direction := rl.Vector3Subtract(camera.Position, camera.Target)

	// Get current distance
	distance := rl.Vector3Length(direction)

	// Calculate new distance
	newDistance := distance - amount

	// Clamp to min/max zoom
	if newDistance < minZoom {
		newDistance = minZoom
	}
	if newDistance > maxZoom {
		newDistance = maxZoom
	}

	// Normalize direction and scale to new distance
	direction = rl.Vector3Scale(rl.Vector3Normalize(direction), newDistance)

	// Update camera position
	camera.Position = rl.Vector3Add(camera.Target, direction)
}
