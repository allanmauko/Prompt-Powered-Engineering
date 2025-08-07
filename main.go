// Package declaration - every Go file must start with this
// "main" is a special package name that tells Go this is an executable program
// (not a library that other programs can import)
package main

// Import statement - brings in external packages we need
// "fmt" is a standard library package for formatted I/O operations
// It provides functions like Print, Println, Printf, Scanf, etc.
import "fmt"

// The main function - this is the entry point of our program
// Go automatically runs this function when the program starts
// Every executable Go program must have exactly one main() function
func main() {
	// fmt.Println() is a function from the "fmt" package
	// Println = "Print line" - it prints text and adds a newline at the end
	// The text inside quotes is called a "string literal"
	// This line outputs the message to the console/terminal
	fmt.Println("Hello, World! Welcome to my GO journey")

	// After this line executes, the program ends
	// Since main() has no return statement, it implicitly returns with exit code 0
	// (which means "success" in most operating systems)
}
