package noSmoke

import (
	"fmt"
)

func Demo1() {
	pocketPrice := 45
	var day int
	fmt.Println("How many days has it been? ")
	fmt.Scan(&day)
	var total int

	total = (pocketPrice * day)
	fmt.Println("Saved money:", total)

	healtPoint := 0.33
	var total2 float64
	var maxPoint int = 100
	total2 = (healtPoint * float64(day))

	if total2 >= 100 {
		total2 = float64(maxPoint)
	}

	fmt.Println("Healt point is rise up %", total2)
}
