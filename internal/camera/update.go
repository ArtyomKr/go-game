package camera

import rl "github.com/gen2brain/raylib-go/raylib"

func (c *Camera) Update() {
	controls := c.CameraControls
	camera := c.Camera
	mousePos := rl.GetMousePosition()
	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	// Edge scrolling
	if mousePos.X < float32(controls.EdgeScrollMargin) {
		moveCamera(camera, -controls.EdgeScrollSpeed, 0)
	}
	if mousePos.X > screenWidth-float32(controls.EdgeScrollMargin) {
		moveCamera(camera, controls.EdgeScrollSpeed, 0)
	}
	if mousePos.Y < float32(controls.EdgeScrollMargin) {
		moveCamera(camera, 0, controls.EdgeScrollSpeed)
	}
	if mousePos.Y > screenHeight-float32(controls.EdgeScrollMargin) {
		moveCamera(camera, 0, -controls.EdgeScrollSpeed)
	}

	// Keyboard movement
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		moveCamera(camera, 0, controls.MovementSpeed)
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		moveCamera(camera, 0, -controls.MovementSpeed)
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		moveCamera(camera, -controls.MovementSpeed, 0)
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		moveCamera(camera, controls.MovementSpeed, 0)
	}

	// Height adjustment
	if rl.IsKeyDown(rl.KeyR) {
		camera.Position.Y -= controls.HeightSpeed
		camera.Target.Y -= controls.HeightSpeed
		if camera.Position.Y < controls.MinHeight {
			camera.Position.Y = controls.MinHeight
			camera.Target.Y = camera.Position.Y - 20.0
		}
	}
	if rl.IsKeyDown(rl.KeyF) {
		camera.Position.Y += controls.HeightSpeed
		camera.Target.Y += controls.HeightSpeed
		if camera.Position.Y > controls.MaxHeight {
			camera.Position.Y = controls.MaxHeight
			camera.Target.Y = camera.Position.Y - 20.0
		}
	}

	// Camera rotation
	if rl.IsKeyDown(rl.KeyQ) {
		rotateCamera(camera, -controls.RotationSpeed)
	}
	if rl.IsKeyDown(rl.KeyE) {
		rotateCamera(camera, controls.RotationSpeed)
	}

	// Middle mouse rotation
	if rl.IsMouseButtonDown(rl.MouseMiddleButton) {
		mouseDelta := rl.GetMouseDelta()
		rotateCamera(camera, mouseDelta.X*0.1)
	}

	// Right mouse panning
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		mouseDelta := rl.GetMouseDelta()
		moveCamera(camera, -mouseDelta.X*0.1, mouseDelta.Y*0.1)
	}

	mouseWheel := rl.GetMouseWheelMove()
	if mouseWheel != 0 {
		zoomCamera(camera, mouseWheel*controls.ZoomSpeed, controls.MinZoom, controls.MaxZoom)
	}
}
