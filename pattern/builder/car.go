package main

import "fmt"

type Car struct {
	Brand        string
	Engine       int
	Brakes       string
	Transmission string
}

func (c *Car) Print() {
	fmt.Printf("%s Engine:[%d] Brakes:[%d] Transmission:[%d]", c.Brand, c.Engine, c.Brakes, c.Transmission)
}
