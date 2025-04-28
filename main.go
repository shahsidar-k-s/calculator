package main

import (
	"fmt"
	"os"
)

// readInput handles user input from the terminal
func readInput() (float64, string, float64) {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator (+, -, *, /, %): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	return a, operator, b
}

// calculate performs the calculation based on operator and operands
func calculate(a float64, operator string, b float64) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("cannot modulo by zero")
		}
		return float64(int(a) % int(b)), nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func main() {
	fmt.Println(" CLI Calculator")

	a, operator, b := readInput()

	result, err := calculate(a, operator, b)
	if err != nil {
		fmt.Println(" Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", a, operator, b, result)
}
