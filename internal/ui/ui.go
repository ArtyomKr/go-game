package ui

import "go-game/internal/state"

type UI struct {
	BuildMenu    *BuildMenu
	DebugOverlay *DebugOverlay
}

func New() *UI {
	return &UI{
		BuildMenu:    NewBuildMenu(),
		DebugOverlay: NewDebugOverlay(),
	}
}

func (ui *UI) Draw(gameState *state.State) {
	ui.DebugOverlay.Draw()

	if gameState.Mode == state.Building {
		ui.BuildMenu.Draw(&gameState.MousePos)
	}
}
