package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var number = [6]int{5, 2, 3, 4, 5, 6}
	// or
	var numberDynamic = []int{5, 2, 3, 4}

	fmt.Println(number)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of integers:", len(number))
	fmt.Println("Number of integers (dynamic):", len(numberDynamic))
}
