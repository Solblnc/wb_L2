package main

import (
	"fmt"
	"math"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("\nCalculating area for square")
	a.area = s.side * s.side
	fmt.Printf("\nArea of square %d", a.area)
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("\nCalculating area for circle")
	a.area = int(math.Pow(float64((c.radius)), 2*math.Pi))
	fmt.Printf("\nArea of circle %d", a.area)

}

func (a *areaCalculator) visitForRectangle(t *rectangle) {
	fmt.Println("\nCalculating area for rectangle")
	a.area = t.l * t.b
	fmt.Printf("\nArea of rectangle %f\n", a.area)
}
