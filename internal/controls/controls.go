package controls

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Controls struct {
	BuildMenuButton rl.MouseButton
}

func New() Controls {
	return Controls{
		BuildMenuButton: rl.MouseLeftButton,
	}
}
