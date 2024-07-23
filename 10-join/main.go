package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var b1 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(15).
		Align(lipgloss.Left)

	row := lipgloss.JoinHorizontal(lipgloss.Top,
		b1.Render("JoinHorizontal Layout"),
		b1.Render("JoinHorizontal "))
	fmt.Println(row)
}
