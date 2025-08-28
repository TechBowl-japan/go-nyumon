package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start FizzBuzz....")

	FizzBuzz() // この関数を実装する

	fmt.Println("FizzBuzz completed!")
}

func FizzBuzz() {
	for x := 1; x <= 100; x++ {
		switch {
		case x%3 == 0 && x%5 == 0:
			fmt.Println("FizzBuzz")
		case x%3 == 0:
			fmt.Println("Fizz")
		case x%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(x)
		}
	}
}
