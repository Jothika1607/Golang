package main

import (
	"fmt"
)

func main() {
	if 20 > 18 { // if condition
		fmt.Println("20 is greater than 18")
	}

	time := 20 //if else
	if time < 18 {
		fmt.Println("Good day.")
	} else { //if else condition
		fmt.Println("Good evening.")
	}

	day := 4
	switch day { // single switch  case
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	}

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
