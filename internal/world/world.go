package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
)

type World struct {
	Buildings []entity.Building
	//Grid      Grid
}

func New() *World {
	return &World{
		Buildings: make([]entity.Building, 0),
		//Grid:      NewGrid(100, 100), // example size
	}
}

func (w *World) Update() {
	// Update world state
}

func (w *World) Draw() {
	//w.Grid.Draw()
	for _, building := range w.Buildings {
		building.Draw()
	}
	rl.DrawGrid(1000, 100.0)
}

func (w *World) AddBuilding(building entity.Building) {
	w.Buildings = append(w.Buildings, building)
}
