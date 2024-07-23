package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func main() {
	rows := [][]string{
		{"Chinese", "您好", "你好"},
		{"Japanese", "こんにちは", "やあ"},
		{"Arabic", "أهلين", "أهلا"},
		{"Russian", "Здравствуйте", "Привет"},
		{"Spanish", "Hola", "¿Qué tal?"},
	}
	t := table.New().
		Border(lipgloss.DoubleBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("LANGUAGE", "FORMAL", "INFORMAL").
		Rows(rows...)

	t.Row("English", "You look absolutely fabulous.", "How's it going?")
	fmt.Println(t)
}
