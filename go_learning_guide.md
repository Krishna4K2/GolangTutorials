# Go Programming Learning Guide - Getting Started

## What is Go?
Go (also called Golang) is a programming language created by Google in 2009. It's designed to be:
- **Simple**: Easy to read and write
- **Fast**: Compiles quickly and runs efficiently  
- **Reliable**: Helps you write bug-free code
- **Scalable**: Great for building large applications

## Setting Up Go
1. **Install Go**: Download from https://golang.org/dl/
2. **Verify installation**: Open terminal/command prompt and type `go version`
3. **Set up workspace**: Create a folder for your Go projects

## Your First Go Program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Let's break this down:**
- `package main`: Every Go program starts with a package declaration
- `import "fmt"`: We're importing the "fmt" package to use print functions
- `func main()`: This is the main function - where your program starts running
- `fmt.Println()`: This prints text to the screen

## Basic Concepts You'll Learn

### 1. Variables and Data Types
```go
// Declaring variables
var name string = "Alice"
var age int = 25
var height float64 = 5.6
var isStudent bool = true

// Short declaration (Go figures out the type)
name := "Alice"
age := 25
```

### 2. Basic Data Types
- `string`: Text ("Hello")
- `int`: Whole numbers (42)
- `float64`: Decimal numbers (3.14)
- `bool`: True or false

### 3. Functions
```go
func greet(name string) string {
    return "Hello, " + name + "!"
}
```

### 4. Control Structures
```go
// If statements
if age >= 18 {
    fmt.Println("You're an adult")
} else {
    fmt.Println("You're a minor")
}

// For loops
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Learning Path
1. **Week 1**: Variables, data types, basic input/output
2. **Week 2**: Functions, control structures (if/else, loops)
3. **Week 3**: Arrays, slices, maps (data collections)
4. **Week 4**: Structs and methods (organizing code)
5. **Week 5**: Error handling and packages
6. **Week 6**: Concurrency basics (goroutines)
7. **Week 7-8**: Building real projects

## Practice Exercise
Try to create and run your first "Hello, World!" program:

1. Create a new file called `main.go`
2. Copy the first program example above
3. Run it with: `go run main.go`

## Next Steps
Once you get your first program running, we'll dive deeper into:
- Understanding variables and how to store different types of data
- Getting user input
- Making decisions with if/else statements
- Writing your own functions
- Building small projects to practice

Ready to start? Let me know when you've set up Go and run your first program, or if you have any questions about getting started!