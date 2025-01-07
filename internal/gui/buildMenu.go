package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawBuildMenu(mousePos *rl.Vector2) {
	rl.DrawRectangle(int32(mousePos.X), int32(mousePos.Y), 100, 100, rl.Fade(rl.SkyBlue, 0.5))
}
