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
	Defeat
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
	min_guesses   int
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "0"
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 5
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
