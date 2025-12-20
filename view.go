package main

import (
	"fmt"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	switch m.state {
	case Start:
		return viewStart()
	case Game:
		return viewGame(m)
	case Win:
		return viewWin(m)
	case Exit:
		return viewExit()
	}
	return ""
}

func viewStart() string {
	wP := welcomeStyle.Render("Welcome to NumberHunter!\n\n")
	var sb strings.Builder
	sb.WriteString("You’ve entered the hunt for the secret number.\n")
	sb.WriteString("Test your intuition and logic — \n")
	sb.WriteString("can you uncover the mystery faster than anyone else?\n\n")
	aP := textStyle.Render(sb.String())
	bP := motoStyle.Render("NumberHunter — the hunt begins, the number awaits!\n\n")
	cP := keyHintStype.Render("(enter to continue, esc to quit)")
	abc := lg.JoinVertical(lg.Center, wP, aP, bP, cP)
	return startStyle.Render(abc)
}
func viewGame(m Model) string {
	var sb strings.Builder
	if m.hot > 0 {
		hP := highStyle.Render("Your strike soared too far... Lower your aim!\n")
		sb.WriteString(hP)
	} else if m.hot < 0 {
		lP := lowStyle.Render("Your arrow fell short... Aim higher, Hunter!\n")
		sb.WriteString(lP)
	}
	textP := textStyle.Render("Hunter, the shadows whisper...\nWhat number will you strike at?")
	inpP := inputStyle.Render(fmt.Sprintf("\n\n%s\n\n", m.textinput.View()))
	exitP := keyHintStype.Render("(enter to check, ecs to quit)")
	abc := lg.JoinVertical(lg.Center, sb.String(), textP, inpP, exitP)
	return gameStyle.Render(abc)
}

func viewWin(m Model) string {
	var sb strings.Builder
	sb.WriteString("Bullseye! You have slain the secret number.\n")
	sb.WriteString("The hunt is complete — glory is yours!\n\n")
	text := textStyle.Render(sb.String())

	logP := countStyle.Render(fmt.Sprintf("You loosed your arrows %d %s",
		len(m.guess_numbers),
		"times in pursuit of the secret number.\n"))
	exitP := keyHintStype.Render("(enter to continue, esc to quit)")
	abc := lg.JoinVertical(lg.Center, text, logP, exitP)
	return winStyle.Render(abc)
}

func viewExit() string {
	str := "The Hunter lays down his bow...\n"
	str += "The hunt ends, and the shadows remain unchallenged.\n"
	exitP := keyHintStype.Render("(enter to quit)")
	abc := lg.JoinVertical(lg.Center, str, exitP)
	return exitStyle.Render(abc)
}
