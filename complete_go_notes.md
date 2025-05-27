# Complete Go Programming Notes

## Table of Contents
1. [Introduction to Go](#introduction-to-go)
2. [Getting Started](#getting-started)
3. [Variables and Data Types](#variables-and-data-types)
4. [Constants](#constants)
5. [Input and Output](#input-and-output)
6. [Control Structures](#control-structures)
7. [Functions](#functions)
8. [Arrays and Slices](#arrays-and-slices)
9. [Maps](#maps)
10. [Structs](#structs)
11. [Pointers](#pointers)
12. [Error Handling](#error-handling)
13. [Packages and Modules](#packages-and-modules)
14. [Concurrency](#concurrency)
15. [File I/O](#file-io)
16. [Testing](#testing)
17. [Best Practices](#best-practices)
18. [Common Patterns](#common-patterns)

---

## Introduction to Go

### What is Go?
- **Created by**: Google (2009)
- **Creators**: Robert Griesemer, Rob Pike, Ken Thompson
- **Type**: Compiled, statically typed language
- **Philosophy**: Simple, readable, efficient

### Why Go?
- **Fast compilation** and execution
- **Built-in concurrency** support
- **Garbage collection** (automatic memory management)
- **Strong standard library**
- **Cross-platform** compilation
- **Simple syntax** - easy to learn and maintain

### Key Features
- Static typing with type inference
- Compiled language (produces single binary)
- Built-in testing framework
- Package management system
- Excellent concurrency primitives (goroutines, channels)

---

## Getting Started

### Installation
1. Download from [golang.org](https://golang.org/dl/)
2. Follow platform-specific installation instructions
3. Verify: `go version`

### Basic Go Program Structure
```go
package main          // Package declaration (main = executable)

import "fmt"          // Import packages

func main() {         // Main function (entry point)
    fmt.Println("Hello, World!")
}
```

### Running Go Programs
```bash
# Run directly
go run main.go

# Build executable
go build main.go
./main  # (or main.exe on Windows)

# Build and install
go install
```

---

## Variables and Data Types

### Variable Declaration Methods

#### Method 1: Explicit Declaration
```go
var name string = "Alice"
var age int = 25
var height float64 = 5.8
var isActive bool = true
```

#### Method 2: Type Inference
```go
var name = "Alice"      // Go infers string
var age = 25           // Go infers int
var height = 5.8       // Go infers float64
var isActive = true    // Go infers bool
```

#### Method 3: Short Declaration (inside functions only)
```go
name := "Alice"
age := 25
height := 5.8
isActive := true
```

#### Method 4: Declare then Assign
```go
var name string
var age int

name = "Alice"
age = 25
```

### Basic Data Types

#### Numeric Types
```go
// Integers
var i int = 42
var i8 int8 = 127        // -128 to 127
var i16 int16 = 32767    // -32768 to 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// Unsigned integers
var ui uint = 42
var ui8 uint8 = 255      // 0 to 255
var ui16 uint16 = 65535  // 0 to 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

#### String Type
```go
var str string = "Hello, World!"
var multiLine string = `This is a
multi-line
string`

// String operations
str1 := "Hello"
str2 := "World"
combined := str1 + " " + str2
length := len(str1)
```

#### Boolean Type
```go
var isTrue bool = true
var isFalse bool = false

// Boolean operations
result := true && false  // AND
result = true || false   // OR
result = !true          // NOT
```

#### Zero Values
```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // ""
```

### Type Conversion
```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var u uint = uint(f)        // Convert float64 to uint
var s string = string(i)    // Convert int to string (ASCII)

// String conversions (using strconv package)
import "strconv"

str := strconv.Itoa(42)           // int to string
num, err := strconv.Atoi("42")    // string to int
f, err := strconv.ParseFloat("3.14", 64)  // string to float
```

---

## Constants

### Declaration
```go
const Pi = 3.14159
const Language = "Go"
const Version = 1.18

// Grouped constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)

// Typed constants
const Pi float64 = 3.14159
const MaxUsers int = 1000
```

### iota (Automatic Incrementing)
```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)

const (
    _ = iota         // Skip 0
    KB = 1 << (10 * iota)  // 1024
    MB                     // 1048576
    GB                     // 1073741824
)
```

---

## Input and Output

### Basic Output
```go
import "fmt"

fmt.Print("Hello")           // Print without newline
fmt.Println("Hello")         // Print with newline
fmt.Printf("Age: %d\n", 25)  // Formatted print

// Common format specifiers
fmt.Printf("%s\n", "string")     // %s for strings
fmt.Printf("%d\n", 42)           // %d for integers
fmt.Printf("%f\n", 3.14)         // %f for floats
fmt.Printf("%.2f\n", 3.14159)    // %.2f for 2 decimal places
fmt.Printf("%t\n", true)         // %t for booleans
fmt.Printf("%v\n", anything)     // %v for any value
```

### Basic Input
```go
import "fmt"

var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name)  // Read one word

// Read entire line
import "bufio"
import "os"

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
```

---

## Control Structures

### If Statements
```go
age := 18

// Basic if
if age >= 18 {
    fmt.Println("You're an adult")
}

// If-else
if age >= 18 {
    fmt.Println("You're an adult")
} else {
    fmt.Println("You're a minor")
}

// If-else if-else
if age < 13 {
    fmt.Println("Child")
} else if age < 18 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Adult")
}

// If with short statement
if num := 10; num > 5 {
    fmt.Println("Number is greater than 5")
}
```

### Switch Statements
```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of work week")
case "Friday":
    fmt.Println("TGIF!")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Regular day")
}

// Switch without expression (like if-else chain)
score := 85

switch {
case score >= 90:
    fmt.Println("A grade")
case score >= 80:
    fmt.Println("B grade")
case score >= 70:
    fmt.Println("C grade")
default:
    fmt.Println("Need improvement")
}

// Type switch
var x interface{} = 42

switch v := x.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### Loops

#### For Loop (Traditional)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Multiple variables
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Printf("i=%d, j=%d\n", i, j)
}
```

#### For Loop (While-like)
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

#### Infinite Loop
```go
for {
    fmt.Println("This runs forever")
    // Use break to exit
    if condition {
        break
    }
}
```

#### Range Loop
```go
// Array/Slice
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// String (iterates over runes)
for index, char := range "Hello" {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}

// Map
ages := map[string]int{"Alice": 25, "Bob": 30}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

#### Loop Control
```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue  // Skip this iteration
    }
    if i == 7 {
        break     // Exit loop
    }
    fmt.Println(i)
}
```

---

## Functions

### Basic Function Syntax
```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

### Function Examples
```go
// Simple function
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a int, b int) int {
    return a + b
}

// Multiple parameters of same type
func multiply(a, b, c int) int {
    return a * b * c
}

// Multiple return values
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return  // naked return
}

// Variadic function (variable number of arguments)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result := sum(1, 2, 3, 4, 5)
```

### Anonymous Functions and Closures
```go
// Anonymous function
add := func(a, b int) int {
    return a + b
}

// Closure
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
```

### Function as First-Class Citizens
```go
// Function as parameter
func apply(fn func(int, int) int, a, b int) int {
    return fn(a, b)
}

add := func(a, b int) int { return a + b }
result := apply(add, 5, 3)  // 8
```

---

## Arrays and Slices

### Arrays
```go
// Array declaration
var arr [5]int                    // Array of 5 integers, zero-valued
var arr2 = [5]int{1, 2, 3, 4, 5} // Array with initial values
arr3 := [...]int{1, 2, 3}        // Size inferred from elements

// Accessing elements
arr[0] = 10
value := arr[0]

// Array properties
length := len(arr)

// Iterate through array
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}
```

### Slices (Dynamic Arrays)
```go
// Slice declaration
var slice []int                    // nil slice
slice = make([]int, 5)            // slice of length 5
slice2 := make([]int, 5, 10)      // length 5, capacity 10
slice3 := []int{1, 2, 3, 4, 5}    // slice literal

// Slice operations
slice = append(slice, 6)           // Add element
slice = append(slice, 7, 8, 9)     // Add multiple elements

// Slicing
arr := [5]int{1, 2, 3, 4, 5}
slice4 := arr[1:4]                 // Elements 1, 2, 3 (index 1 to 3)
slice5 := arr[:3]                  // First 3 elements
slice6 := arr[2:]                  // From index 2 to end

// Slice properties
length := len(slice)
capacity := cap(slice)

// Copy slices
source := []int{1, 2, 3}
dest := make([]int, len(source))
copy(dest, source)

// Multi-dimensional slices
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

---

## Maps

### Map Declaration and Usage
```go
// Map declaration
var m map[string]int               // nil map
m = make(map[string]int)           // Initialize map
m2 := make(map[string]int)         // Short form
m3 := map[string]int{              // Map literal
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}

// Map operations
m["Alice"] = 25                    // Set value
age := m["Alice"]                  // Get value
age, exists := m["Alice"]          // Get with existence check

delete(m, "Alice")                 // Delete key

// Check if key exists
if age, ok := m["Alice"]; ok {
    fmt.Printf("Alice is %d years old\n", age)
} else {
    fmt.Println("Alice not found")
}

// Iterate through map
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Map properties
length := len(m)
```

### Map with Different Types
```go
// String to slice map
groups := map[string][]string{
    "fruits":     {"apple", "banana", "orange"},
    "vegetables": {"carrot", "broccoli", "spinach"},
}

// Nested maps
users := map[string]map[string]interface{}{
    "alice": {
        "age":    25,
        "email":  "alice@example.com",
        "active": true,
    },
}
```

---

## Structs

### Basic Struct Definition
```go
// Struct definition
type Person struct {
    Name    string
    Age     int
    Email   string
    IsActive bool
}

// Creating struct instances
var p1 Person                           // Zero-valued struct
p2 := Person{"Alice", 25, "alice@example.com", true}
p3 := Person{
    Name:     "Bob",
    Age:      30,
    Email:    "bob@example.com",
    IsActive: true,
}

// Accessing struct fields
p1.Name = "Charlie"
fmt.Println(p1.Name)
```

### Struct Methods
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver (can modify struct)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()
rect.Scale(2)
```

### Embedded Structs (Composition)
```go
type Address struct {
    Street string
    City   string
    State  string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedded struct
}

// Usage
p := Person{
    Name: "Alice",
    Age:  25,
    Address: Address{
        Street: "123 Main St",
        City:   "Anytown",
        State:  "CA",
    },
}

// Can access embedded fields directly
fmt.Println(p.Street)  // Same as p.Address.Street
```

### Struct Tags
```go
import "encoding/json"

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"`
    Password string `json:"-"`  // Ignored in JSON
}

// JSON marshaling/unmarshaling
user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
jsonData, _ := json.Marshal(user)
fmt.Println(string(jsonData))
```

---

## Pointers

### Pointer Basics
```go
// Pointer declaration
var p *int      // Pointer to int
var x int = 42
p = &x          // p points to x

// Dereferencing
fmt.Println(*p) // Print value at pointer (42)
*p = 100        // Change value through pointer
fmt.Println(x)  // x is now 100

// new() function creates pointer to zero value
p2 := new(int)  // Creates pointer to zero int
*p2 = 50
```

### Pointers with Functions
```go
// Function that modifies value through pointer
func increment(x *int) {
    *x++
}

num := 10
increment(&num)
fmt.Println(num)  // 11

// Pointer to struct
type Person struct {
    Name string
    Age  int
}

func updateAge(p *Person, newAge int) {
    p.Age = newAge  // Go automatically dereferences
}

person := Person{Name: "Alice", Age: 25}
updateAge(&person, 26)
```

### Pointer vs Value Receivers
```go
type Counter struct {
    count int
}

// Value receiver - creates copy
func (c Counter) GetCount() int {
    return c.count
}

// Pointer receiver - works with original
func (c *Counter) Increment() {
    c.count++
}

counter := Counter{count: 0}
counter.Increment()  // count becomes 1
```

---

## Error Handling

### Basic Error Handling
```go
import (
    "errors"
    "fmt"
)

// Function that returns error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Printf("Result: %f\n", result)
```

### Custom Error Types
```go
// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}

// Function using custom error
func validateAge(age int) error {
    if age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "cannot be negative",
        }
    }
    if age > 150 {
        return ValidationError{
            Field:   "age",
            Message: "cannot exceed 150",
        }
    }
    return nil
}
```

### Error Wrapping (Go 1.13+)
```go
import "fmt"

func processFile(filename string) error {
    err := readFile(filename)
    if err != nil {
        return fmt.Errorf("failed to process %s: %w", filename, err)
    }
    return nil
}

// Unwrapping errors
import "errors"

var originalErr error
if errors.Is(err, originalErr) {
    // Handle specific error
}

var customErr *ValidationError
if errors.As(err, &customErr) {
    // Handle custom error type
}
```

### Panic and Recover
```go
// Panic - stops normal execution
func riskyFunction() {
    panic("something went wrong!")
}

// Recover - catches panics
func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    riskyFunction()  // This will panic
    fmt.Println("This won't execute")
}
```

---

## Packages and Modules

### Package Basics
```go
// Every Go file belongs to a package
package main  // Executable package

package mypackage  // Library package

// Importing packages
import "fmt"
import "strings"
import "net/http"

// Multiple imports
import (
    "fmt"
    "strings"
    "net/http"
)

// Import with alias
import (
    f "fmt"
    str "strings"
)

// Import for side effects only
import _ "image/png"
```

### Creating Your Own Package
```go
// File: math/operations.go
package math

// Exported function (starts with capital letter)
func Add(a, b int) int {
    return a + b
}

// Unexported function (starts with lowercase)
func multiply(a, b int) int {
    return a * b
}

// Exported variable
var Pi = 3.14159

// unexported variable
var version = "1.0"
```

### Go Modules
```bash
# Initialize module
go mod init github.com/username/projectname

# Add dependency
go get github.com/gorilla/mux

# Update dependencies
go mod tidy

# Vendor dependencies
go mod vendor
```

### go.mod File Example
```
module github.com/username/myproject

go 1.19

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.7
)

require (
    github.com/gorilla/context v1.1.1 // indirect
)
```

---

## Concurrency

### Goroutines
```go
import (
    "fmt"
    "time"
)

// Regular function
func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    // Run in goroutine
    go say("world")
    say("hello")
    
    // Anonymous goroutine
    go func() {
        fmt.Println("Anonymous goroutine")
    }()
    
    time.Sleep(time.Second)  // Wait for goroutines
}
```

### Channels
```go
// Channel declaration
ch := make(chan int)        // Unbuffered channel
ch2 := make(chan int, 5)    // Buffered channel

// Sending and receiving
go func() {
    ch <- 42  // Send value
}()

value := <-ch  // Receive value

// Channel directions in function parameters
func sender(ch chan<- int) {    // Send-only channel
    ch <- 42
}

func receiver(ch <-chan int) {  // Receive-only channel
    value := <-ch
    fmt.Println(value)
}
```

### Channel Patterns

#### Worker Pool
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 9; r++ {
        <-results
    }
}
```

#### Select Statement
```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(1 * time.Second)
    ch1 <- "from ch1"
}()

go func() {
    time.Sleep(2 * time.Second)
    ch2 <- "from ch2"
}()

for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout")
    default:
        fmt.Println("no channels ready")
        time.Sleep(500 * time.Millisecond)
    }
}
```

### Sync Package
```go
import "sync"

// WaitGroup
var wg sync.WaitGroup

for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Printf("Goroutine %d\n", id)
    }(i)
}

wg.Wait()  // Wait for all goroutines

// Mutex
var mu sync.Mutex
var counter int

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}

// Once
var once sync.Once

func setup() {
    fmt.Println("Setup called")
}

func doSomething() {
    once.Do(setup)  // Will only call setup once
}
```

---

## File I/O

### Reading Files
```go
import (
    "io/ioutil"
    "os"
    "bufio"
)

// Read entire file
data, err := ioutil.ReadFile("file.txt")
if err != nil {
    panic(err)
}
fmt.Println(string(data))

// Read file line by line
file, err := os.Open("file.txt")
if err != nil {
    panic(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

### Writing Files
```go
// Write entire file
data := []byte("Hello, World!")
err := ioutil.WriteFile("output.txt", data, 0644)
if err != nil {
    panic(err)
}

// Write with more control
file, err := os.Create("output.txt")
if err != nil {
    panic(err)
}
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Hello, World!\n")
writer.Flush()
```

### Working with JSON
```go
import "encoding/json"

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Marshal (struct to JSON)
person := Person{Name: "Alice", Age: 25}
jsonData, err := json.Marshal(person)
if err != nil {
    panic(err)
}
fmt.Println(string(jsonData))

// Unmarshal (JSON to struct)
jsonString := `{"name":"Bob","age":30}`
var p Person
err = json.Unmarshal([]byte(jsonString), &p)
if err != nil {
    panic(err)
}
fmt.Printf("%+v\n", p)
```

---

## Testing

### Basic Testing
```go
// File: math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-1, 1)
    expected := 0
    
    if result != expected {
        t.Errorf("Add(-1, 1) = %d; want %d", result, expected)
    }
}
```

### Table-Driven Tests
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
        {-2, -3, -5},
    }
    
    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", 
                test.a, test.b, result, test.expected)
        }
    }
}
```

### Benchmarking
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

### Running Tests
```bash
go test                    # Run tests in current package
go test ./...             # Run tests in all packages
go test -v                # Verbose output
go test -cover            # Show coverage
go test -bench=.          # Run benchmarks
```

---

## Best Practices

### Code Organization
- One package per directory
- Package names should be short and clear
- Use meaningful variable and function names
- Keep functions small and focused
- Group related functionality

### Error Handling
- Always handle errors explicitly
- Don't ignore errors with `_`
- Use custom error types for specific cases
- Wrap errors with context when appropriate

### Concurrency
- Don't communicate by sharing memory; share memory by communicating
- Use channels for coordination between goroutines
- Always close channels when done sending
- Use `sync.WaitGroup` to wait for multiple goroutines

### Performance
- Use pointers for large structs
- Prefer slices over arrays for function parameters
- Use `make()` with capacity when size is known
- Profile before optimizing

### Security
- Validate all inputs
- Use `crypto