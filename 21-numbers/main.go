package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var b1 = lipgloss.NewStyle().
		Width(3).
		Margin(1, 1, 1, 1).
		Align(lipgloss.Left)

	row := lipgloss.JoinHorizontal(lipgloss.Top,
		b1.Render(BigNumber(0)),
		b1.Render(BigNumber(1)),
		b1.Render(BigNumber(2)),
		b1.Render(BigNumber(3)),
		b1.Render(BigNumber(4)),
		b1.Render(BigNumber(5)),
		b1.Render(BigNumber(6)),
		b1.Render(BigNumber(7)),
		b1.Render(BigNumber(8)),
		b1.Render(BigNumber(9)))
	fmt.Println(row)
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

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	return strings.Join(digits[n], "\n")
}
