package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read the input from the user
	var input string
	fmt.Scanln(&input)
	fmt.Println("Input:", input)

	// Define the regular expression pattern to match the input
	pattern := "^[-+]?[0-9]+(\\.[0-9]+)?(\\s*[-+*/]\\s*[0-9]+(\\.[0-9]+)?)*$"
	match, _ := regexp.MatchString(pattern, input)

	if !match {
		fmt.Println("Invalid input")
		return
	}

	// Split the input into numbers and operators
	numbers := strings.FieldsFunc(input, func(c rune) bool {
		return c == '+' || c == '-' || c == '*' || c == '/'
	})
	operators := strings.FieldsFunc(input, func(c rune) bool {
		return c >= '0' && c <= '9' || c == '.'
	})

	// Convert the numbers from strings to float64
	var operands []float64
	for _, n := range numbers {
		f, _ := strconv.ParseFloat(n, 64)
		operands = append(operands, f)
	}

	// Perform the calculations
	var result float64 = operands[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += operands[i+1]
		case "-":
			result -= operands[i+1]
		case "*":
			result *= operands[i+1]
		case "/":
			result /= operands[i+1]
		}
	}

	fmt.Println("Result:", result)
}