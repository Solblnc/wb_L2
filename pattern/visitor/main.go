package main

import "fmt"

func main() {
	square := square{side: 2}
	circle := circle{radius: 6}
	rectangle := rectangle{
		l: 3,
		b: 4,
	}

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
	fmt.Println()
}
