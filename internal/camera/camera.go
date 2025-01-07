package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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
