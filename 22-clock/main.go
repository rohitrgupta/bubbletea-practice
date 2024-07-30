package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type clock struct {
	quitting bool
	t        int
}

func (m clock) Init() tea.Cmd {
	return tickCmd()
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
		m.t = int(time.Now().Unix())
		return m, tickCmd()
	}
	return m, nil
}

func (m clock) View() string {
	if m.t == 0 {
		return ""
	}
	var b1 = lipgloss.NewStyle().
		Width(3).
		Margin(1, 1, 1, 1).
		Align(lipgloss.Left)
	sec := (m.t + 5.5*3600) % 86400
	min := sec / 60
	sec %= 60
	hour := min / 60
	min %= 60

	s := lipgloss.JoinHorizontal(lipgloss.Top,
		b1.Render(BigNumber(hour/10)),
		b1.Render(BigNumber(hour%10)),
		b1.Render(BigNumber(10)),
		b1.Render(BigNumber(min/10)),
		b1.Render(BigNumber(min%10)),
		b1.Render(BigNumber(10)),
		b1.Render(BigNumber(sec/10)),
		b1.Render(BigNumber(sec%10)),
	)

	if m.quitting {
		s += "\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(clock{t: int(time.Now().Unix())}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type placeholder []string

func BigNumber(n int) string {
	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	dot := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine, dot,
	}
	return strings.Join(digits[n], "\n")
}
