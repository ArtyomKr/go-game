package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type BuildMenu struct {
	title string
}

func NewBuildMenu() *BuildMenu {
	return &BuildMenu{
		title: "Select building",
	}
}

func (bm *BuildMenu) Draw(mousePos *rl.Vector2) {
	rl.DrawRectangle(int32(mousePos.X), int32(mousePos.Y), 100, 100, rl.Fade(rl.SkyBlue, 0.5))
}
