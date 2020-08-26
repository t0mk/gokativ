package gokativ

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	s := []string{
		"Věra",
	}
	for _, i := range s {
		fmt.Println(Vokativ(i), IsFem(i))
	}
}
