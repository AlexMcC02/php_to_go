package main

import "fmt"

// Fizzbuzz.
// Loop from 1 to 100 (inclusive).
// Echo out fizz if divisible by 3.
// Echo out buzz if divisible by 5.
// Echo out fizzbuzz if divisible by 3 or 5 (or 15).

func main() {
	for i := 1; i <= 100; i++ {
		if (i % 15 == 0) {
			fmt.Printf("%d FizzBuzz!\n", i)
		} else if (i % 5 == 0) {
			fmt.Printf("%d Buzz!\n", i)
		} else if (i % 3 == 0) {
			fmt.Printf("%d Fizz!\n", i)
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
