package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println("Normal")

	var bold = lipgloss.NewStyle().Bold(true)
	fmt.Println(bold.Render("Bold"))

	var itelic = lipgloss.NewStyle().Italic(true)
	fmt.Println(itelic.Render("Itelic"))

	var faint = lipgloss.NewStyle().Faint(true)
	fmt.Println(faint.Render("Faint"))

	var blink = lipgloss.NewStyle().Blink(true)
	fmt.Println(blink.Render("Blink"))

	var strike = lipgloss.NewStyle().Strikethrough(true)
	fmt.Println(strike.Render("Strike"))

	var under = lipgloss.NewStyle().Underline(true)
	fmt.Println(under.Render("Under"))

	var reverse = lipgloss.NewStyle().Reverse(true)
	fmt.Println(reverse.Render("Reverse"))

}
