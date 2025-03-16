package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
)

type ID string

type Entity interface {
	GetID() ID
	Drawable
	Updatable
	Collidable
	Unloadable
}

func NewID() ID {
	return ID(uuid.New().String())
}

type BaseEntity struct {
	ID       ID
	Position rl.Vector3
}

func NewBaseEntity(position rl.Vector3) BaseEntity {
	return BaseEntity{
		ID:       NewID(),
		Position: position,
	}
}

func (b *BaseEntity) GetID() ID {
	return b.ID
}
func (b *BaseEntity) Update()         {}
func (b *BaseEntity) Draw()           {}
func (b *BaseEntity) Unload()         {}
func (b *BaseEntity) CheckCollision() {}

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
