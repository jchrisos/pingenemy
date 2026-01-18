package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page struct {
	text string
}

func NewPage(text string) Page {
	return Page{text: text}
}

func (p Page) Init() tea.Cmd {
	return nil
}

func (p Page) View() string {
	return ""
}

func (p Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}
