package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
}

func NewTui() Tui {
	return Tui{}
}

func (t Tui) Start() {
	p := tea.NewProgram(NewPage("Hello"))

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
