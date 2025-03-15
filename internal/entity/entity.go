package entity

import rl "github.com/gen2brain/raylib-go/raylib"

// Drawable defines an interface for entities that can be displayed.
type Drawable interface {
	Draw()
}

// Updatable defines an interface for entities that can be updated.
type Updatable interface {
	Update()
}

// Unloadable defines an interface for entities that need cleanup.
type Unloadable interface {
	Unload()
}

// Collidable defines an interface for entities that can collide with ray.
type Collidable interface {
	CheckCollision(ray rl.Ray) (rl.RayCollision, bool)
}
