package main

import "fmt"

// Function with explicit types for parameters and return
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Multiple return values (common in Go)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	message := greet("World")
	fmt.Println(message)

	result, err := divide(0, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
