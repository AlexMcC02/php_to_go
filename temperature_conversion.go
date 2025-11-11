package main

import "fmt"

func celsiusToFarenheit(temp int) int {
	return (temp * (9 / 5)) + 32
}

func farenheitToCelsius(temp int) int {
	return (temp - 32) * 5 / 9
}

func choicer(running *bool) {
	var choice int
	fmt.Print("\n1. Convert to Celsius.\n2. Convert to Farenheit.\n3. Quit.\n\n> ")
	fmt.Scan(&choice);

	switch choice {
	case 1:
		var temperature int
		fmt.Print("\nEnter the temperature in Farenheit:\n\n> ")
		fmt.Scan(&temperature)
		fmt.Printf("\n%d Degrees Farenheit is equivalent to %d Degrees Celsius\n", temperature, farenheitToCelsius(temperature))
	case 2:
		var temperature int
		fmt.Print("\nEnter the temperature in Celsius:\n\n> ")
		fmt.Scan(&temperature)
		fmt.Printf("\n%d Degrees Celsius is equivalent to %d Degrees Farenheit\n", temperature, celsiusToFarenheit(temperature))
	case 3:
		*running = false
	}
}

func temperature_conversion() {
	var running bool = true

	for running {
		choicer(&running)
	}
}