package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
)

type Mode int

const (
	Normal Mode = iota
	Building
)

type World struct {
	Buildings map[entity.ID]*entity.Building
	Floor     *entity.Floor
}

type State struct {
	Mode       Mode
	MousePos   rl.Vector2
	SelectedID entity.ID
	World
}

func New() *State {
	return &State{
		Mode:     Normal,
		MousePos: rl.Vector2{},
		World: World{
			Buildings: make(map[entity.ID]*entity.Building),
			Floor:     entity.NewFloor(rl.NewVector3(0, 0, 0), rl.NewVector2(100, 100))},
	}
}

// GetAllWorldEntities returns a slice of all world entities in the state.
func (s *State) GetAllWorldEntities() []entity.Entity {
	// TODO: rewrite this to avoid creating new array every time. Use listener pattern to listen to State.World changes in world package?
	entities := make([]entity.Entity, 0)

	// Add the floor
	if s.World.Floor != nil {
		entities = append(entities, s.World.Floor)
	}

	// Add all buildings
	for _, building := range s.World.Buildings {
		entities = append(entities, building)
	}

	return entities
}
