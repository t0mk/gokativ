# Gokativ

Czech Vocative noun case in Go

Generování oslovení ze jména (oslovení = 5. pád = vokativ)

## Usage

```
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
```

## Credits

This is almost a verbatim copy of https://github.com/Mimino666/vokativ
