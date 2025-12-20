package main

import (
	"math/rand/v2"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

func getRandomInt() int {
	return rand.IntN(100) + 1
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEsc:
			if m.state == Exit {
				return m, tea.Quit
			}
			m.state = Exit
			return m, tea.ClearScreen
		case tea.KeyEnter:
			switch m.state {
			case Start:
				m.state = Game
				m.target_number = getRandomInt()
				m.guess_numbers = []int{}
				m.hot = 0
				return m, tea.ClearScreen
			case Game:
				num, err := strconv.Atoi(m.textinput.Value())
				if err != nil {
					return m, tea.ClearScreen
				}
				m.guess_number = num
				m.guess_numbers = append(m.guess_numbers, num)
				if m.guess_number == m.target_number {
					m.state = Win
				} else {
					m.hot = m.guess_number - m.target_number
				}
				m.textinput.Reset()
				return m, tea.ClearScreen
			case Win:
				m.state = Start
				m.guess_numbers = []int{}
				return m, tea.ClearScreen
			case Exit:
				return m, tea.Quit
			}
			return m, tea.ClearScreen
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}
