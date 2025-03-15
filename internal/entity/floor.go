package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Floor struct {
	Position  rl.Vector3
	Size      rl.Vector2
	Material  rl.Material
	Mesh      rl.Mesh
	Transform rl.Matrix
}

func NewFloor(position rl.Vector3, size rl.Vector2) *Floor {
	mesh := rl.GenMeshPlane(size.X, size.Y, 10, 10)
	transform := rl.MatrixTranslate(position.X, position.Y, position.Z)
	material := rl.LoadMaterialDefault()

	return &Floor{
		Position:  position,
		Size:      size,
		Material:  material,
		Mesh:      mesh,
		Transform: transform,
	}
}

func (f *Floor) Draw() {
	rl.DrawMesh(f.Mesh, f.Material, f.Transform)
}

func (f *Floor) Unload() {
	rl.UnloadMesh(&f.Mesh)
	rl.UnloadMaterial(f.Material)
}

func (f *Floor) CheckCollision(ray rl.Ray) (rl.RayCollision, bool) {
	collision := rl.GetRayCollisionMesh(ray, f.Mesh, f.Transform)
	return collision, collision.Hit
}
