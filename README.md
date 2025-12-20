# Number Hunter

**Number Hunter** is a console-based implementation of the classic 
*Guess the Number* game, built with **Golang** and powered by the **Bubble Tea**
and **Lipgloss** libraries for a stylish terminal UI.

## Gameplay

- The computer secretly chooses a random number between **1 and 100**.  
- The player must guess the number within a **limited number of attempts**.  
- After each guess, hints are provided **Too low** or **Too high**
- If the player guesses correctly, the hunt ends in **victory**.  
- If attempts run out, the hunt ends in **defeat**.  

Players can continue the adventure by starting **multiple rounds** until they 
decide to exit.

## Features

- Random number generation between 1–100  
- Limited attempts per round  
- Victory and defeat messages with immersive hunter-themed narrative  
- Option to replay multiple rounds  
- Stylish **TUI** built with:
  - [Bubble Tea](https://github.com/charmbracelet/bubbletea) – functional and 
  reactive terminal UI framework  
  - [Lipgloss](https://github.com/charmbracelet/lipgloss) – beautiful styling 
  for terminal applications  

## Getting Started

### Prerequisites
- [Go](https://golang.org/) 1.25+ installed on your system

### Installation
```bash
git clone https://github.com/Mirsait/number-hunter.git
cd number-hunter
go build -o build/
cd build
./number-hunter
```

## License
[MIT License](LICENSE) — feel free to use, modify, and distribute.
