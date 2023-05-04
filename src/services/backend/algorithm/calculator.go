package algorithm

import (
	"fmt"
	"regexp"
	"strconv"
)

func Calculator(input string) {
	fmt.Println("Input:", input)

	pattern := `^\s*-?\s*(\(\s*)*\s*-?\s*\d+(\.\d+)?(\s*\))*\s*(\s*[-+*/]\s*(\(\s*)*\s*-?\s*(\(\s*)*\d+(\.\d+)?(\s*\))*\s*)*$`
	match, _ := regexp.MatchString(pattern, input)

	if !match {
		fmt.Println("Invalid input")
		return
	}

	result, _ := evaluateExpression(input)
	fmt.Println("Result:", result)
}

func evaluateExpression(expression string) (float64, error) {
	numbers := make([]float64, 0)
	operators := make([]string, 0)
	i := 0
	for i < len(expression) {
		if expression[i] == ' ' {
			i++
			continue
		}
		switch expression[i] {
		case '(':
			subExpression := extractSubExpression(expression[i+1:])
			if subExpression == "" {
				return 0, fmt.Errorf("invalid input: unmatched parentheses")
			}
			result, err := evaluateExpression(subExpression)
			if err != nil {
				return 0, err
			}
			numbers = append(numbers, result)
			i += len(subExpression) + 2 // +2 to skip ')' as well
		case '+', '-', '*', '/':
			if len(numbers) == 0 {
				// Handle negative numbers
				if expression[i] == '-' {
					numbers = append(numbers, 0)
					operators = append(operators, string(expression[i]))
					i++
					continue
				} else {
					return 0, fmt.Errorf("invalid input")
				}
			}
			operators = append(operators, string(expression[i]))
			i++
		default:
			start := i
			for i < len(expression) && isDigit(expression[i]) {
				i++
			}
			number, _ := strconv.ParseFloat(expression[start:i], 64)
			numbers = append(numbers, number)
		}
	}

	for i := len(operators) - 1; i >= 0; i-- {
		switch operators[i] {
		case "*":
			numbers[i] *= numbers[i+1]
			numbers = numbers[:len(numbers)-1]
		case "/":
			numbers[i] /= numbers[i+1]
			numbers = numbers[:len(numbers)-1]
		}
	}

	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += numbers[i+1]
		case "-":
			result -= numbers[i+1]
		}
	}

	return result, nil
}

func extractSubExpression(expression string) string {
	count := 1
	i := 0
	for i = 0; i < len(expression); i++ {
		if expression[i] == '(' {
			count++
		} else if expression[i] == ')' {
			count--
		}
		if count == 0 {
			break
		}
	}
	if count != 0 {
		fmt.Println("Invalid input: unmatched parentheses")
		return ""
	}
	return expression[:i]
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9' || char == '.'
}
