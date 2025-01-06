package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type CameraControls struct {
	EdgeScrollSpeed  float32
	MovementSpeed    float32
	RotationSpeed    float32
	HeightSpeed      float32
	ZoomSpeed        float32
	EdgeScrollMargin int32
	MinHeight        float32
	MaxHeight        float32
	MinZoom          float32
	MaxZoom          float32
}

func NewCameraControls() CameraControls {
	return CameraControls{
		EdgeScrollSpeed:  0.5,
		MovementSpeed:    1,
		RotationSpeed:    2.0,
		HeightSpeed:      0.5,
		ZoomSpeed:        1.0,
		EdgeScrollMargin: 20,
		MinHeight:        5.0,
		MaxHeight:        50.0,
		MinZoom:          5.0,
		MaxZoom:          50.0,
	}
}

func InitCamera() rl.Camera3D {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0.0, 20.0, -20.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	return camera
}

func UpdateCamera(camera *rl.Camera3D, controls *CameraControls) {
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
		panCamera(camera, mouseDelta.X*0.1, mouseDelta.Y*0.1)
	}

	mouseWheel := rl.GetMouseWheelMove()
	if mouseWheel != 0 {
		zoomCamera(camera, mouseWheel*controls.ZoomSpeed, controls.MinZoom, controls.MaxZoom)
	}
}

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

func panCamera(camera *rl.Camera3D, dx, dy float32) {
	// Pan camera based on screen space movement
	right := rl.Vector3CrossProduct(
		rl.Vector3Subtract(camera.Target, camera.Position),
		camera.Up,
	)
	right = rl.Vector3Normalize(right)

	movement := rl.Vector3Add(
		rl.Vector3Scale(right, dx),
		rl.Vector3Scale(camera.Up, -dy),
	)

	camera.Position = rl.Vector3Add(camera.Position, movement)
	camera.Target = rl.Vector3Add(camera.Target, movement)
}

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
