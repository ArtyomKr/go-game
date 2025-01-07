package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/controls"
)

type GameCamera struct {
	Camera         *rl.Camera3D
	CameraControls *controls.CameraControls
}

func InitCamera() *rl.Camera3D {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0.0, 20.0, -20.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	return &camera
}

func New() *GameCamera {
	return &GameCamera{
		Camera:         InitCamera(),
		CameraControls: controls.NewCameraControls(),
	}
}
