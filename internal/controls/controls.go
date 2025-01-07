package controls

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/gui"
)

type Controls struct {
	BuildMenuButton rl.MouseButton
}

func New() Controls {
	return Controls{
		BuildMenuButton: rl.MouseLeftButton,
	}
}

func Update(controls *Controls) {
	if rl.IsMouseButtonDown(controls.BuildMenuButton) {
		mousePos := rl.GetMousePosition()
		gui.DrawBuildMenu(&mousePos)
	}
}
