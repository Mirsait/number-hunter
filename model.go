package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type GameState uint

const (
	Start GameState = iota
	Game
	Win
	Exit
)

type Model struct {
	state         GameState
	guess_number  int
	guess_numbers []int
	target_number int
	textinput     textinput.Model
	err           error
	hot           int
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "0"
	ti.Focus()
	ti.CharLimit = 2
	ti.Width = 20
	return Model{
		textinput:     ti,
		err:           nil,
		guess_number:  0,
		target_number: 0,
		hot:           0,
		state:         Start,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
