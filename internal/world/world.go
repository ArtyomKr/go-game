package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go-game/internal/entity"
)

type World struct {
	Entities []entity.Drawable
	Floor    *entity.Floor
	//Grid      Grid
}

func New() *World {
	return &World{
		Entities: make([]entity.Drawable, 0),
		Floor:    entity.NewFloor(rl.NewVector3(0, 0, 0), rl.NewVector2(100, 100)),
	}
}

func (w *World) AddEntity(e entity.Drawable) {
	w.Entities = append(w.Entities, e)
}

// GetEntityAtRay checks for collisions between a ray and collidable entities in the world.
// It returns the entity that was hit (if any) and the collision information.
func (w *World) GetEntityAtRay(ray rl.Ray) (entity.Drawable, rl.RayCollision) {
	var closestEntity entity.Drawable
	var closestCollision rl.RayCollision
	closestCollision.Distance = -1 // Initialize with an invalid distance

	// Check for collisions with all entities
	for _, e := range w.Entities {
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
	for _, e := range w.Entities {
		if updatable, ok := e.(entity.Updatable); ok {
			updatable.Update()
		}
	}
}

func (w *World) Draw() {
	w.Floor.Draw()

	for _, e := range w.Entities {
		e.Draw()
	}

	rl.DrawGrid(1000, 100.0)
}

func (w *World) Unload() {
	// Unload all unloadable entities
	for _, e := range w.Entities {
		if unloadable, ok := e.(entity.Unloadable); ok {
			unloadable.Unload()
		}
	}
}
