package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
)

// AddBuilding adds a building to the state.
func (s *State) AddBuilding(building *entity.Building) {
	s.Buildings[building.ID] = building
}

// RemoveBuilding removes a building from the state.
func (s *State) RemoveBuilding(id entity.ID) {
	delete(s.Buildings, id)
}

// UpdateBuilding updates a building's position or other properties.
func (s *State) UpdateBuilding(id entity.ID, position rl.Vector3) {
	if building, exists := s.Buildings[id]; exists {
		building.Position = position
	}
}
