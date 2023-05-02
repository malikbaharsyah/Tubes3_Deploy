package algorithm

import (
	"fmt"
	"regexp"
	"time"
)

func Calendar(input string) {
	// Read the input from the user
	// fmt.Scanln(&input)
	// fmt.Println("Input:", input)

	// Define the regular expression pattern to match the input
	pattern := "((0[1-9]|[1-9]|[12]\\d|3[01])/(0[1-9]|1[0-2]|[1-9])/[12]\\d{3})"
	match, _ := regexp.MatchString(pattern, input)

	if !match {
		fmt.Println("Invalid input")
		return
	}

	// Parse the date from the input
	var layout string
	switch len(input) {
	case 10:
		layout = "02/01/2006"
	case 9:
		if input[1] == '/' {
			layout = "2/01/2006"
		} else {
			layout = "02/1/2006"
		}
	case 8:
		layout = "2/1/2006"
	default:
		fmt.Println("Invalid input")
		return
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	date, err := time.ParseInLocation(layout, input, loc)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	fmt.Println("Date:", date.Format("Monday, January 2, 2006"))
}
