package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type BuildingType int

const (
	House BuildingType = iota
	Workshop
	Market
)

type Building struct {
	Type     BuildingType
	Position rl.Vector3
	Size     rl.Vector3
	Color    rl.Color
}

func (b *Building) Draw() {
	rl.DrawCube(b.Position, b.Size.X, b.Size.Y, b.Size.Z, b.Color)
	rl.DrawCubeWires(b.Position, b.Size.X, b.Size.Y, b.Size.Z, rl.Maroon)
}

func (b *Building) CheckCollision(ray rl.Ray) (rl.RayCollision, bool) {
	min := rl.Vector3Subtract(b.Position, rl.Vector3Scale(b.Size, 0.5))
	max := rl.Vector3Add(b.Position, rl.Vector3Scale(b.Size, 0.5))
	boundingBox := rl.NewBoundingBox(min, max)

	collision := rl.GetRayCollisionBox(ray, boundingBox)
	return collision, collision.Hit
}
