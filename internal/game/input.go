package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
	"go-game/internal/state"
)

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
	ray := rl.GetScreenToWorldRay(g.State.MousePos, *g.Camera.Camera)

	floorCol, _ := g.World.Floor.CheckCollision(ray)

	entityCol, _ := g.World.GetEntityAtRay(ray)

	fmt.Println(entityCol)

	// Handle building mode inputs
	if floorCol.Hit {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			g.World.AddEntity(&entity.Building{Position: rl.NewVector3(floorCol.Point.X, 10, floorCol.Point.Z), Size: rl.NewVector3(10, 10, 10), Color: rl.Red})
		}
	}

}
