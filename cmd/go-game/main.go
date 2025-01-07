package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	cm "go-game/internal/camera"
	"go-game/internal/controls"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth, screenHeight, "Strategy Game")
	defer rl.CloseWindow()

	camera := cm.InitCamera()
	cameraControls := cm.NewCameraControls()

	gameControls := controls.New()

	cubePosition := rl.NewVector3(0, 0, 0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update camera based on controls
		cm.UpdateCamera(&camera, &cameraControls)
		controls.Update(&gameControls)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(cubePosition, 5.0, 5.0, 5.0, rl.Red)
		rl.DrawCubeWires(cubePosition, 5.0, 5.0, 5.0, rl.Maroon)

		rl.DrawGrid(1000, 100.0)

		rl.EndMode3D()

		rl.DrawText("WASD/Arrows: Move Camera\nQ/E: Rotate Camera\nR/F: Adjust Height\nMiddle Mouse: Rotate\nRight Mouse: Pan",
			10, 40, 20, rl.DarkGray)

		rl.DrawText(fmt.Sprintf("Camera position: X:%v Y:%v Z: %v", camera.Position.X, camera.Position.Y, camera.Position.Z), 10, 200, 20, rl.DarkGray)
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
