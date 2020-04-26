package main

import (
	"github.com/01-edu/z01"

	correct "./correct"
	student "./student"
)

func main() {
	table := z01.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		z01.RandStr(z01.RandIntBetween(1, 15), z01.RandAlnum()),
	)
	for _, arg := range table {
		z01.Challenge("LastRune", student.LastRune, correct.LastRune, arg)
	}
}
