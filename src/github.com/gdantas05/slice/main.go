package main

import "fmt"

func main() {

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
		"Better 2",
	}

	fmt.Println(sliceString)
	fmt.Println(len(sliceString))
	fmt.Println(cap(sliceString))
	fmt.Println(append(sliceString, "Better 3"))
	fmt.Println(len(sliceString))
	fmt.Println(cap(sliceString))
	fmt.Println(sliceString[0])   // HELLO
	fmt.Println(sliceString[:2])  //HELLO WORLD
	fmt.Println(sliceString[1:2]) //WORLD
	fmt.Println(sliceString[2:4]) //MUCH BETTER
	fmt.Println(sliceString[2:])  //MUCH BETTER BETTER 2
}
