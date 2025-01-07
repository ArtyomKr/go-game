package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugOverlay struct {
	showFPS       bool
	showCameraPos bool
	showControls  bool
}

func NewDebugOverlay() *DebugOverlay {
	return &DebugOverlay{
		showFPS:       true,
		showCameraPos: true,
		showControls:  true,
	}
}

func (do *DebugOverlay) Draw() {
	if do.showFPS {
		rl.DrawFPS(10, 10)
	}

	if do.showControls {
		rl.DrawText("WASD/Arrows: Move Camera\nQ/E: Rotate Camera\nR/F: Adjust Height\nMiddle Mouse: Rotate\nRight Mouse: Pan",
			10, 40, 20, rl.DarkGray)
	}

	//rl.DrawText(fmt.Sprintf("Camera position: X:%v Y:%v Z: %v", camera.Position.X, camera.Position.Y, camera.Position.Z), 10, 200, 20, rl.DarkGray)
}
