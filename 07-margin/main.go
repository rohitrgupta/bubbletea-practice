package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	var b1 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		MarginTop(2)
	fmt.Println(b1.Render("Margin Top"))

	var b2 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		MarginBottom(2)
	fmt.Println(b2.Render("Margin Bottom"))

	var b3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		MarginLeft(2)
	fmt.Println(b3.Render("Margin Left"))

	var b4 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		MarginRight(2)
	fmt.Println(b4.Render("Margin Right"))

	var b5 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Margin(2, 4)
	fmt.Println(b5.Render("Margin Set"))

	var b6 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Margin(2, 4, 3, 1)
	fmt.Println(b6.Render("Margin All"))

}
