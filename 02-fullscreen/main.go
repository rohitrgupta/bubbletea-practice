package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type HelloBubble struct {
	text string
}

func (m HelloBubble) Init() tea.Cmd {
	return nil
}

func (m HelloBubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m HelloBubble) View() string {
	return m.text
}

func main() {
	p := tea.NewProgram(HelloBubble{text: "Hello Bulbbles\n"}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
