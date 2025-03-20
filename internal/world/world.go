package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
	"go-game/internal/state"
)

type World struct {
	State *state.State
}

func New(state *state.State) *World {
	return &World{
		State: state,
	}
}

// GetEntityAtRay checks for collisions between a ray and collidable entities in the world.
// It returns the entity that was hit (if any) and the collision information.
func (w *World) GetEntityAtRay(ray rl.Ray) (entity.Drawable, rl.RayCollision) {
	var closestEntity entity.Drawable
	var closestCollision rl.RayCollision
	closestCollision.Distance = -1 // Initialize with an invalid distance

	// Check for collisions with all entities
	for _, e := range w.State.GetAllWorldEntities() {
		if collidable, ok := e.(entity.Collidable); ok {
			collision, hit := collidable.CheckCollision(ray)
			if hit && (closestCollision.Distance < 0 || collision.Distance < closestCollision.Distance) {
				closestCollision = collision
				closestEntity = e
			}
		}
	}

	return closestEntity, closestCollision
}

// *** life cycle methods ***

func (w *World) Update() {
	// Update all updatable entities
	for _, e := range w.State.GetAllWorldEntities() {
		if updatable, ok := e.(entity.Updatable); ok {
			updatable.Update()
		}
	}
}

func (w *World) Draw() {
	for _, e := range w.State.GetAllWorldEntities() {
		if drawable, ok := e.(entity.Drawable); ok {
			drawable.Draw()
		}
	}

	if w.State.Mode == state.Building {
		w.DrawGrid(10)
	}

	w.drawAxisHelpers(100)
	rl.DrawGrid(1000, 100.0)
}

func (w *World) Unload() {
	// Unload all unloadable entities
	for _, e := range w.State.GetAllWorldEntities() {
		if unloadable, ok := e.(entity.Unloadable); ok {
			unloadable.Unload()
		}
	}
}
