package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if (len(os.Args) < 4) {
		panic("Correct use: go run <file> <x> <operator> <y>.")
	}

	x, _ := strconv.Atoi(os.Args[1])
	operator := os.Args[2]
	y, _ := strconv.Atoi(os.Args[3])
	result := 0

	switch operator {
	case "x":
		result = x * y
	case "/":
		result = x / y
	case "+":
		result = x + y
	case "i":
		result = x - y
	default:
		fmt.Println("Unidentified operator.")
	}

	fmt.Printf("%d %s %d = %d\n", x, operator, y, result)
}