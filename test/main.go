package main

import (
	"fmt"

	"github.com/t0mk/gokativ"
)

func main() {
	ns := []string{
		"Karásek", "Bláha", "Matějů", "Petr", "Jana", "Tomáš",
		"Tříska", "Kolek",
		"Kovářová", "Kaplan", "Herůdek", "Profesor", "Doktorka",
	}
	for _, n := range ns {
		fmt.Println(n, gokativ.Vokativ(n), gokativ.IsWoman(n))
	}
}
