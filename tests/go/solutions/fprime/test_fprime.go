package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := []string{}
	for i := 0; i < 10; i++ {
		table = append(table, strconv.Itoa(z01.RandIntBetween(1, 100)))
	}

	table = append(table, " ")
	table = append(table, "1")
	table = append(table, "1 1")
	table = append(table, "hello")
	table = append(table, "p 1")
	table = append(table, "804577")
	table = append(table, "225225")
	table = append(table, "8333325")
	table = append(table, "42")
	table = append(table, "9539")
	table = append(table, "1000002")
	table = append(table, "1000003")

	for _, s := range table {
		z01.ChallengeMain("fprime", strings.Fields(s)...)
	}
}
