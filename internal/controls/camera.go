package controls

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

func NewCameraControls() *CameraControls {
	return &CameraControls{
		EdgeScrollSpeed:  0.5,
		MovementSpeed:    1,
		RotationSpeed:    2.0,
		HeightSpeed:      0.5,
		ZoomSpeed:        1.0,
		EdgeScrollMargin: 20,
		MinHeight:        5.0,
		MaxHeight:        50.0,
		MinZoom:          20.0,
		MaxZoom:          100.0,
	}
}
