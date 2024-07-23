package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	var b1 = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder())
	fmt.Println(b1.Render("NormalBorder"))

	var b2 = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder())
	fmt.Println(b2.Render("RoundedBorder"))

	var b3 = lipgloss.NewStyle().
		BorderStyle(lipgloss.BlockBorder())
	fmt.Println(b3.Render("BlockBorder"))

	var b4 = lipgloss.NewStyle().
		BorderStyle(lipgloss.OuterHalfBlockBorder())
	fmt.Println(b4.Render("OuterHalfBlockBorder"))

	var b5 = lipgloss.NewStyle().
		BorderStyle(lipgloss.InnerHalfBlockBorder())
	fmt.Println(b5.Render("InnerHalfBlockBorder"))

	var b6 = lipgloss.NewStyle().
		BorderStyle(lipgloss.ThickBorder())
	fmt.Println(b6.Render("ThickBorder"))

	var b7 = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder())
	fmt.Println(b7.Render("DoubleBorder"))

	var b8 = lipgloss.NewStyle().
		BorderStyle(lipgloss.HiddenBorder())
	fmt.Println(b8.Render("HiddenBorder"))

}
