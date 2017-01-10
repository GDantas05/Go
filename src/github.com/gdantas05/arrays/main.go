package main

import "fmt"

func main() {
	var x [10]int

	for i := 0; i < 5; i++ {

		x[i] = i * i
		fmt.Println(x)
	}
}
