package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type clock struct {
	quitting bool
	t        int
}

func (m clock) Init() tea.Cmd {
	return tickEvery()
}

func (m clock) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}
	case tickMsg:
		m.t++
		time.Sleep(time.Microsecond * 100)
		return m, tickEvery()
	}
	return m, nil
}

func (m clock) View() string {
	s := fmt.Sprintf("%d %s\n", m.t, time.Now().Format("15:04:05.000"))
	if m.quitting {
		s += "\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(clock{})

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

func tickEvery() tea.Cmd {
	return tea.Every(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
