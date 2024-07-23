package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/list"
)

func main() {
	l := list.New("A", "B", "C")
	fmt.Println(l)
	fmt.Println("")

	l2 := list.New(
		"A", list.New("Artichoke"),
		"B", list.New("Baking Flour", "Bananas", "Barley", "Bean Sprouts").Enumerator(list.Roman),
		"C", list.New("Cashew Apple", "Cashews", "Coconut Milk", "Curry Paste", "Currywurst").Enumerator(list.Arabic),
		"D", list.New("Dill", "Dragonfruit", "Dried Shrimp").Enumerator(list.Alphabet),
		"E", list.New("Eggs"),
		"F", list.New("Fish Cake", "Furikake").Enumerator(list.Bullet),
		"J", list.New("Jicama"),
		"K", list.New("Kohlrabi"),
		"L", list.New("Leeks", "Lentils", "Licorice Root").Enumerator(list.Tree),
	)
	fmt.Println(l2)

}
