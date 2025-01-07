package camera

import rl "github.com/gen2brain/raylib-go/raylib"

func moveCamera(camera *rl.Camera3D, right, forward float32) {
	// Calculate forward and right vectors
	forwardVec := rl.Vector3Subtract(camera.Target, camera.Position)
	forwardVec.Y = 0
	forwardVec = rl.Vector3Normalize(forwardVec)

	rightVec := rl.Vector3CrossProduct(forwardVec, camera.Up)
	rightVec = rl.Vector3Normalize(rightVec)

	// Move both position and target
	movement := rl.Vector3Add(
		rl.Vector3Scale(rightVec, right),
		rl.Vector3Scale(forwardVec, forward),
	)

	camera.Position = rl.Vector3Add(camera.Position, movement)
	camera.Target = rl.Vector3Add(camera.Target, movement)
}
