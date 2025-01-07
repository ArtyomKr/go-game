package state

import rl "github.com/gen2brain/raylib-go/raylib"

type Mode int

const (
	Normal Mode = iota
	Building
)

type State struct {
	Mode     Mode
	MousePos rl.Vector2
}

func NewState() *State {
	return &State{
		Mode:     Normal,
		MousePos: rl.Vector2{},
	}
}
