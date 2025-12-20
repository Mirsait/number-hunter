package main

import (
	"os"

	lg "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func getWidth() (int, int) {
	var fd = os.Stdout.Fd()
	var physicalWidth, physicalHeight, _ = term.GetSize(fd)
	return physicalWidth, physicalHeight
}

var w, h = getWidth()

const background = lg.Color("#000e17")
const hintColor = lg.Color("#444")
const mainTextColor = lg.Color("#FFFEB3")
const inputColor = lg.Color("#10B6B0")
const accentColor = lg.Color("#FF71CC")
const lowColor = lg.Color("#3943FF")
const highColor = lg.Color("#ff4339")
const motoColor = lg.Color("#B94FFF")
const exitColor = lg.Color("#b50303")
const defeatBgColor = lg.Color("#000e17")
const defeatFgColor = lg.Color("#ff4339")

// start view
var startStyle = lg.NewStyle().
	Width(w).Height(h).
	Background(background).
	PaddingTop(10).
	PaddingLeft(4).
	PaddingRight(4).
	PaddingBottom(10).
	Bold(true)

var welcomeStyle = lg.NewStyle().
	Width(w).
	Align(lg.Center).
	Foreground(accentColor)

var textStyle = lg.NewStyle().
	Width(w).
	Align(lg.Center).
	Foreground(mainTextColor)

var motoStyle = lg.NewStyle().
	Width(w).
	Align(lg.Center).
	Bold(true).
	Foreground(motoColor)

var keyHintStype = lg.NewStyle().
	Width(h).
	Align(lg.Center).
	Foreground(hintColor)

// game view
var gameStyle = lg.NewStyle().
	Width(w).Height(h).
	Background(background).
	PaddingTop(10).
	PaddingLeft(4).
	PaddingRight(4).
	PaddingBottom(10).
	Bold(true)

var lowStyle = lg.NewStyle().
	Width(h).
	Align(lg.Center).
	Foreground(lowColor)

var highStyle = lg.NewStyle().
	Width(h).
	Align(lg.Center).
	Foreground(highColor)

var inputStyle = lg.NewStyle().
	Width(h).
	Align(lg.Center).
	Foreground(inputColor)

// win view
var winStyle = lg.NewStyle().
	Width(w).Height(h).
	Background(background).
	PaddingTop(10).
	PaddingLeft(4).
	PaddingRight(4).
	PaddingBottom(10).
	Bold(true)

var countStyle = lg.NewStyle().
	Align(lg.Center).
	Foreground(accentColor)

// exit view
var exitStyle = lg.NewStyle().
	Width(w).Height(h).
	Bold(true).
	Foreground(exitColor).
	Background(background).
	PaddingTop(10).
	PaddingLeft(4).
	PaddingRight(4).
	PaddingBottom(10).
	Align(lg.Center)

var defeatStyle = lg.NewStyle().
	Width(w).Height(h).
	Bold(true).
	Background(defeatBgColor).
	PaddingTop(10).
	PaddingLeft(4).
	PaddingRight(4).
	PaddingBottom(10).
	Align(lg.Center)

var defeatTextStyle = lg.NewStyle().
	Width(h).
	Align(lg.Center).
	Foreground(defeatFgColor).
	Background(defeatBgColor)
