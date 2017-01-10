package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {

	car1 := Car{"Corola", 2017, "White"}
	fmt.Println(car1.info())
}
