package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/game"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Strategy Game")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	game := game.New()

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}
}
