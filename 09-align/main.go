package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	var b1 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(24).
		Align(lipgloss.Left)

	fmt.Println(b1.Render("Align Left"))

	var b2 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(24).
		Align(lipgloss.Right)

	fmt.Println(b2.Render("Align Right"))

	var b3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Width(24).
		Align(lipgloss.Center)

	fmt.Println(b3.Render("Align Center"))

}
