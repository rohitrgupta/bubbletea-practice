package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	for i := range 256 {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color(fmt.Sprintf("%d", i))).
			Width(5)
		fmt.Print(style.Render(fmt.Sprintf("%d", i)))
		if i%36 == 35 {
			fmt.Println()
		}
	}
	fmt.Println("")
}
