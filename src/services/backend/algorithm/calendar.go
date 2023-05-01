package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	// Read the input from the user
	var input string
	fmt.Scanln(&input)
	fmt.Println("Input:", input)

	// Define the regular expression pattern to match the input
	pattern := "((0[1-9]|[1-9]|[12]\\d|3[01])/(0[1-9]|1[0-2]|[1-9])/[12]\\d{3})"
	match, _ := regexp.MatchString(pattern, input)

	if !match {
		fmt.Println("Invalid input")
		return
	}

	// Parse the date from the input
	layout := "02/01/2006"
	date, _ := time.Parse(layout, input)

	fmt.Println("Date:", date.Format("Monday, January 2, 2006"))
}
