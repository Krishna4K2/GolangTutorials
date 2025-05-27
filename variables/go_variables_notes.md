# Go Variables - Complete Reference

## Table of Contents
1. [Variable Declaration](#variable-declaration)
2. [Variable Types](#variable-types)
3. [Zero Values](#zero-values)
4. [Variable Scope](#variable-scope)
5. [Constants](#constants)
6. [Type Conversion](#type-conversion)
7. [Best Practices](#best-practices)
8. [Examples](#examples)

## Variable Declaration

### 1. Using `var` keyword

```go
// Basic declaration
var name string
var age int
var isActive bool

// Declaration with initialization
var name string = "John"
var age int = 25
var isActive bool = true

// Multiple variables
var (
    name     string = "John"
    age      int    = 25
    isActive bool   = true
)

// Type inference
var name = "John"        // string inferred
var age = 25            // int inferred
var isActive = true     // bool inferred
```

### 2. Short Variable Declaration (`:=`)

```go
// Only inside functions
name := "John"
age := 25
isActive := true

// Multiple assignment
name, age := "John", 25
x, y, z := 1, 2, 3
```

### 3. Declaration vs Assignment

```go
var x int    // Declaration
x = 10       // Assignment

y := 20      // Declaration + Assignment (short form)
```

## Variable Types

### Basic Types

#### Numeric Types

```go
// Integers
var i8 int8 = 127           // -128 to 127
var i16 int16 = 32767       // -32768 to 32767
var i32 int32 = 2147483647  // -2^31 to 2^31-1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63-1

var ui8 uint8 = 255         // 0 to 255
var ui16 uint16 = 65535     // 0 to 65535
var ui32 uint32 = 4294967295 // 0 to 2^32-1
var ui64 uint64 = 18446744073709551615 // 0 to 2^64-1

// Platform dependent
var i int = 42              // int32 or int64
var ui uint = 42            // uint32 or uint64
var ptr uintptr             // integer type to hold pointer

// Aliases
var b byte = 255            // alias for uint8
var r rune = 'A'            // alias for int32 (Unicode code point)

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

#### String Type

```go
var str string = "Hello, World!"
var multiline string = `This is a
multi-line string
using backticks`

// String operations
var greeting = "Hello"
var name = "World"
var message = greeting + ", " + name + "!"
```

#### Boolean Type

```go
var isTrue bool = true
var isFalse bool = false
var result bool = (5 > 3) // true
```

### Composite Types

#### Arrays

```go
var arr1 [5]int                    // Array of 5 integers
var arr2 = [5]int{1, 2, 3, 4, 5}  // Initialized array
var arr3 = [...]int{1, 2, 3}      // Compiler determines size
```

#### Slices

```go
var slice1 []int               // nil slice
var slice2 = []int{1, 2, 3}    // Slice literal
var slice3 = make([]int, 5)    // Slice with length 5
var slice4 = make([]int, 5, 10) // Slice with length 5, capacity 10
```

#### Maps

```go
var map1 map[string]int                    // nil map
var map2 = map[string]int{"a": 1, "b": 2} // Map literal
var map3 = make(map[string]int)           // Empty map
```

#### Structs

```go
type Person struct {
    Name string
    Age  int
}

var person1 Person
var person2 = Person{"John", 25}
var person3 = Person{Name: "Jane", Age: 30}
```

#### Pointers

```go
var ptr *int        // Pointer to int
var x int = 42
ptr = &x           // Address of x
var value = *ptr   // Dereference pointer
```

#### Channels

```go
var ch1 chan int              // nil channel
var ch2 = make(chan int)      // Unbuffered channel
var ch3 = make(chan int, 10)  // Buffered channel
```

#### Interfaces

```go
var i interface{}        // Empty interface
var w io.Writer         // Interface type
```

#### Functions

```go
var fn func(int, int) int    // Function variable
fn = func(a, b int) int {
    return a + b
}
```

## Zero Values

Go initializes variables with zero values when no explicit value is provided:

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var ptr *int     // nil
var slice []int  // nil
var m map[string]int // nil
var ch chan int  // nil
var fn func()    // nil
```

## Variable Scope

### Package Level (Global)

```go
package main

var globalVar = "I'm global"

func main() {
    // Can access globalVar here
}
```

### Function Level

```go
func example() {
    var localVar = "I'm local"
    // localVar only accessible within this function
}
```

### Block Level

```go
func example() {
    if true {
        var blockVar = "I'm in a block"
        // blockVar only accessible within this if block
    }
    // blockVar not accessible here
}
```

### Shadowing

```go
var x = "global"

func example() {
    var x = "function"  // Shadows global x
    {
        var x = "block" // Shadows function x
        fmt.Println(x)  // Prints "block"
    }
    fmt.Println(x)      // Prints "function"
}
```

## Constants

### Basic Constants

```go
const pi = 3.14159
const greeting = "Hello"
const isDebug = false

// Multiple constants
const (
    StatusOK     = 200
    StatusNotFound = 404
    StatusError  = 500
)
```

### Typed Constants

```go
const typedInt int = 42
const typedString string = "Hello"
```

### Iota (Enumerated Constants)

```go
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)

// Custom iota patterns
const (
    _ = iota         // Skip 0
    KB = 1 << (10 * iota) // 1024
    MB               // 1048576
    GB               // 1073741824
)
```

## Type Conversion

### Explicit Conversion

```go
var i int = 42
var f float64 = float64(i)  // int to float64
var u uint = uint(f)        // float64 to uint

var str string = string(65) // int to string (ASCII)
var num int = int(3.14)     // float to int (truncates)
```

### String Conversions

```go
import "strconv"

// String to number
var str = "123"
var num, err = strconv.Atoi(str)      // string to int
var f, err = strconv.ParseFloat(str, 64) // string to float64

// Number to string
var i = 123
var str = strconv.Itoa(i)             // int to string
var str2 = strconv.FormatFloat(3.14, 'f', 2, 64) // float to string
```

## Best Practices

### Naming Conventions

```go
// Use camelCase for variables
var userName string
var maxRetryCount int

// Use PascalCase for exported variables
var DatabaseURL string

// Use short names for short-lived variables
for i := 0; i < 10; i++ {
    // i is fine for loop counter
}

// Use descriptive names for longer-lived variables
var userAuthenticationToken string
```

### Declaration Guidelines

```go
// Prefer := for local variables
func example() {
    name := "John"  // Good
    var name = "John" // Less preferred locally
}

// Use var for package-level variables
var GlobalConfig = Config{}

// Use var when zero value is desired
var counter int // 0

// Use var for complex initialization
var (
    config = loadConfig()
    logger = createLogger()
)
```

### Error Handling

```go
// Check errors immediately
value, err := someFunction()
if err != nil {
    return err
}

// Use blank identifier for unused values
_, err := someFunction()
if err != nil {
    return err
}
```

## Examples

### Complete Example Program

```go
package main

import (
    "fmt"
    "strconv"
)

// Package-level variables
var (
    appName    = "Go Variables Demo"
    version    = "1.0.0"
    debugMode  = false
)

// Constants
const (
    MaxUsers = 100
    DefaultTimeout = 30
)

type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    // Basic variable declarations
    var message string = "Welcome to Go!"
    count := 0
    
    // Array and slice
    var scores [3]int = [3]int{85, 90, 78}
    names := []string{"Alice", "Bob", "Charlie"}
    
    // Map
    userAges := map[string]int{
        "Alice":   25,
        "Bob":     30,
        "Charlie": 28,
    }
    
    // Struct
    user := User{
        ID:   1,
        Name: "John Doe",
        Age:  35,
    }
    
    // Pointer
    var ptr *int = &count
    
    // Type conversion
    ageStr := strconv.Itoa(user.Age)
    
    // Output
    fmt.Printf("App: %s v%s\n", appName, version)
    fmt.Printf("Message: %s\n", message)
    fmt.Printf("Count: %d (via pointer: %d)\n", count, *ptr)
    fmt.Printf("Scores: %v\n", scores)
    fmt.Printf("Names: %v\n", names)
    fmt.Printf("User Ages: %v\n", userAges)
    fmt.Printf("User: %+v\n", user)
    fmt.Printf("Age as string: %s\n", ageStr)
}
```

### Variable Lifecycle Example

```go
package main

import "fmt"

var globalVar = "I'm global"

func demonstrateScope() {
    var functionVar = "I'm in function"
    
    fmt.Println("Global:", globalVar)
    fmt.Println("Function:", functionVar)
    
    // Block scope
    {
        var blockVar = "I'm in block"
        functionVar = "Modified in block"
        
        fmt.Println("Block:", blockVar)
        fmt.Println("Function (modified):", functionVar)
    }
    
    // blockVar is not accessible here
    fmt.Println("Function after block:", functionVar)
}

func main() {
    demonstrateScope()
}
```

### Common Patterns

```go
// Multiple assignment
a, b := 1, 2
a, b = b, a  // Swap values

// Error handling pattern
if value, err := someFunction(); err != nil {
    return err
} else {
    // Use value
    fmt.Println(value)
}

// Map lookup pattern
if value, ok := myMap[key]; ok {
    // Key exists
    fmt.Println("Found:", value)
} else {
    // Key doesn't exist
    fmt.Println("Not found")
}

// Type assertion pattern
if str, ok := value.(string); ok {
    // value is a string
    fmt.Println("String value:", str)
}
```

---

## Quick Reference

| Declaration | Syntax | Scope | Notes |
|-------------|--------|-------|-------|
| `var x int` | Long form | Any | Zero value initialized |
| `var x = 5` | Type inference | Any | Type inferred from value |
| `x := 5` | Short form | Function only | Declaration + assignment |
| `const x = 5` | Constant | Any | Immutable value |

### Memory Layout

- **Stack**: Local variables, function parameters
- **Heap**: Dynamically allocated memory (pointers, slices, maps)
- **Data segment**: Global variables, constants

### Performance Tips

- Use appropriate types (don't use `int64` if `int32` suffices)
- Prefer stack allocation over heap when possible
- Initialize slices and maps with known capacity when possible
- Use constants for values that don't change

---

*This reference covers the fundamental aspects of Go variables. For advanced topics like reflection, unsafe operations, or memory management, refer to the official Go documentation.*