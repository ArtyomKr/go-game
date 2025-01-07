package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/camera"
	"go-game/internal/state"
	"go-game/internal/ui"
	"go-game/internal/world"
)

type Game struct {
	World  *world.World
	UI     *ui.UI
	Camera *camera.GameCamera
	State  *state.State
}

func New() *Game {
	return &Game{
		World:  world.New(),
		UI:     ui.New(),
		Camera: camera.New(),
		State:  state.NewState(),
	}
}

func (g *Game) Update() {
	g.Camera.Update()
	g.handleInput()
	g.World.Update()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	// 3D World
	rl.BeginMode3D(*g.Camera.Camera)
	g.World.Draw()
	rl.EndMode3D()

	// UI Layer
	g.UI.Draw(g.State)

	rl.EndDrawing()
}

func (g *Game) handleInput() {
	g.State.MousePos = rl.GetMousePosition()

	// Handle global inputs like state changes
	if rl.IsKeyPressed(rl.KeyB) {
		if g.State.Mode == state.Normal {
			g.State.Mode = state.Building
		} else {
			g.State.Mode = state.Normal
		}
	}

	// Handle state-specific inputs
	switch g.State.Mode {
	case state.Normal:
		g.handleNormalInput()
	case state.Building:
		g.handleBuildingInput()
	}
}

func (g *Game) handleNormalInput() {
	// Handle normal gameplay inputs
}

func (g *Game) handleBuildingInput() {
	// Handle building mode inputs
}
