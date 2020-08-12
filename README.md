# Gokativ

Czech Vocative noun case in Go

## API

* `func Vokativ(name string) string` Vrátí oslovení pro dané jméno
* `func IsWoman(name string) bool` Test jestli je jméno ženské

* `func Vokativ(name string) string` Returns vocative for given name (noun)
* `func IsWoman(name string) bool` Tests whether name (noun) is feminine

## Usage

Viz [test/main.go](test/main.go)

## Test

```
cd test
go run main.go
```

## Credits

This is almost a verbatim copy of https://github.com/Mimino666/vokativ
