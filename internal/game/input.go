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

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		entityCol, collusion := g.World.GetEntityAtRay(ray)

		switch v := entityCol.(type) {
		case *entity.Building:
			fmt.Println(entityCol, v)
		case *entity.Floor:
			fmt.Println(collusion, v)
			building := entity.NewBuilding(entity.BuildingType(1), rl.NewVector2(collusion.Point.X, collusion.Point.Z), rl.NewVector3(10, 10, 10), rl.Red)
			g.State.AddBuilding(building)
		}
	}
}
