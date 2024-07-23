package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	var b1 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		PaddingTop(2)
	fmt.Println(b1.Render("Padding Top"))

	var b2 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		PaddingBottom(2)
	fmt.Println(b2.Render("Padding Bottom"))

	var b3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		PaddingLeft(2)
	fmt.Println(b3.Render("Padding Left"))

	var b4 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		PaddingRight(2)
	fmt.Println(b4.Render("Padding Right"))

	var b5 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Padding(2, 4)
	fmt.Println(b5.Render("Padding Set"))

	var b6 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Padding(2, 4, 3, 1)
	fmt.Println(b6.Render("Padding All"))

}
