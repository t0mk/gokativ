# Gokativ

Czech Vocative noun case in Go

Generování oslovení ze jména (oslovení = 5. pád = vokativ)

Pouzivane na https://www.oslovovani.cz

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
		fmt.Println(n, gokativ.Vokativ(n), gokativ.IsFem(n))
	}
}
```

.. vypise:

```
tomk@xps ~/gokativ/test ±master » go run main.go                           1 ↵
Karásek Karásku false
Bláha Bláho false
Matějů Matějů false
Petr Petře false
Jana Jano true
Tomáš Tomáši false
Tříska Třísko false
Kolek Kolku false
Kovářová Kovářová true
Kaplan Kaplane false
Herůdek Herůdku false
Profesor Profesore false
Doktorka Doktorko false           <-- chyba IsFem
```

## Credits

This is almost a verbatim copy of https://github.com/Mimino666/vokativ
