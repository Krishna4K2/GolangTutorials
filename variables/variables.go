package main

import "fmt"

func main() {
	// ========== LESSON 1: VARIABLES AND DATA TYPES ==========

	// Method 1: Explicit variable declaration
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.8
	var isStudent bool = true

	fmt.Println("=== Method 1: Explicit Declaration ===")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	fmt.Println()

	// Method 2: Let Go figure out the type (type inference)
	name2 := "Bob"     // Go knows this is a string
	age2 := 30         // Go knows this is an int
	salary := 50000.50 // Go knows this is a float64

	fmt.Println("=== Method 2: Type Inference (Shorter Way) ===")
	fmt.Println("Name:", name2)
	fmt.Println("Age:", age2)
	fmt.Println("Salary:", salary)
	fmt.Println()

	// Method 3: Declare first, assign later
	var city string
	var population int

	city = "New York"
	population = 8000000

	fmt.Println("=== Method 3: Declare Then Assign ===")
	fmt.Println("City:", city)
	fmt.Println("Population:", population)
	fmt.Println()

	// ========== COMMON DATA TYPES ==========

	// Strings (text)
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName // Combining strings

	// Numbers
	var wholeNumber int = 42
	var decimal float64 = 3.14159
	var negativeNumber int = -10

	// Boolean (true/false)
	var isRaining bool = false
	var isSunny bool = true

	fmt.Println("=== Different Data Types ===")
	fmt.Println("Full Name:", fullName)
	fmt.Println("Whole Number:", wholeNumber)
	fmt.Println("Decimal:", decimal)
	fmt.Println("Negative:", negativeNumber)
	fmt.Println("Is Raining?", isRaining)
	fmt.Println("Is Sunny?", isSunny)
	fmt.Println()

	// ========== BASIC OPERATIONS ==========

	// Math operations
	a := 10
	b := 3

	fmt.Println("=== Math Operations ===")
	fmt.Println("a =", a, ", b =", b)
	fmt.Println("Addition (a + b):", a+b)
	fmt.Println("Subtraction (a - b):", a-b)
	fmt.Println("Multiplication (a * b):", a*b)
	fmt.Println("Division (a / b):", a/b)
	fmt.Println("Remainder (a % b):", a%b)
	fmt.Println()

	// String operations
	greeting := "Hello"
	world := "World"
	message := greeting + " " + world + "!"

	fmt.Println("=== String Operations ===")
	fmt.Println("Combined message:", message)
	fmt.Println()

	// ========== CHANGING VARIABLE VALUES ==========

	score := 0
	fmt.Println("Initial score:", score)

	score = 10 // Changing the value
	fmt.Println("After scoring:", score)

	score = score + 5 // Adding to existing value
	fmt.Println("After bonus:", score)

	score += 3 // Shorter way to add
	fmt.Println("After another bonus:", score)
}
