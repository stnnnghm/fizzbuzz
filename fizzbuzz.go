package main

import "fmt"

func main() {
	fmt.Println("Enter an integer to play fizzbuzz:")
	var input int

	// Gather user input
	fmt.Scanf("%d", &input)

	// Fizzbuzz using user input
	for i := 0; i <= input; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(v int) {
	fizz := "fizz"
	buzz := "buzz"

	// Zero doesn't count
	if v == 0 {
		fmt.Println(v)
		return
	}

	// Magic calculations
	if v%3 == 0 && v%5 == 0 {
		fmt.Println(v, fizz+buzz)
	} else if v%3 == 0 {
		fmt.Println(v, fizz)
	} else if v%5 == 0 {
		fmt.Println(v, buzz)
	} else {
		fmt.Println(v)
	}
}
