# Complete Go Variables Reference Guide

## Table of Contents
1. [Introduction to Variables](#introduction-to-variables)
2. [Variable Declaration Methods](#variable-declaration-methods)
3. [Data Types in Detail](#data-types-in-detail)
4. [Zero Values](#zero-values)
5. [Type Inference](#type-inference)
6. [Type Conversion](#type-conversion)
7. [Variable Scope](#variable-scope)
8. [Constants](#constants)
9. [Memory Management](#memory-management)
10. [Best Practices](#best-practices)
11. [Common Patterns](#common-patterns)
12. [Examples and Use Cases](#examples-and-use-cases)

---

## Introduction to Variables

### What are Variables?
Variables are **named storage locations** in computer memory that hold data values. Think of them as labeled boxes where you can store different types of information.

### Key Characteristics in Go
- **Statically typed**: Variable type is determined at compile time
- **Type safe**: Go prevents operations between incompatible types
- **Zero-valued**: Variables have default values when declared
- **Strongly typed**: No implicit type conversions

### Variable Lifecycle
1. **Declaration**: Variable is announced with a name and type
2. **Initialization**: Variable is given an initial value
3. **Usage**: Variable is read from or written to
4. **Scope**: Variable exists within its defined scope
5. **Garbage Collection**: Memory is automatically freed when no longer needed

---

## Variable Declaration Methods

### Method 1: Explicit Declaration with var
```go
// Basic syntax: var name type = value
var name string = "Alice"
var age int = 25
var height float64 = 5.8
var isActive bool = true

// Without initial value (gets zero value)
var username string     // ""
var userAge int        // 0
var balance float64    // 0.0
var isLoggedIn bool    // false

// Multiple variables of same type
var firstName, lastName string = "John", "Doe"
var x, y, z int = 1, 2, 3

// Group declaration
var (
    server   string = "localhost"
    port     int    = 8080
    timeout  int    = 30
    useHTTPS bool   = true
)
```

### Method 2: Type Inference with var
```go
// Go automatically determines the type
var name = "Alice"        // string
var age = 25             // int
var height = 5.8         // float64
var isActive = true      // bool
var data = []byte{1, 2}  // []byte

// Multiple variables with type inference
var firstName, lastName = "John", "Doe"
var x, y = 10, 20.5  // x is int, y is float64

// Group declaration with type inference
var (
    server   = "localhost"    // string
    port     = 8080          // int
    timeout  = time.Second   // time.Duration
    config   = map[string]string{"env": "dev"}  // map[string]string
)
```

### Method 3: Short Declaration (:=)
```go
// Only available inside functions
func main() {
    name := "Alice"        // string
    age := 25             // int
    height := 5.8         // float64
    isActive := true      // bool
    
    // Multiple variables
    firstName, lastName := "John", "Doe"
    x, y := 10, 20.5
    
    // Mixed declaration and assignment
    var existing int = 5
    existing, newVar := 10, "hello"  // existing reassigned, newVar declared
}

// NOT allowed at package level
// name := "Alice"  // This would cause an error outside functions
```

### Method 4: Declaration then Assignment
```go
// Declare first
var name string
var age int
var config map[string]string

// Assign later
name = "Alice"
age = 25
config = make(map[string]string)
config["environment"] = "production"

// Function scope example
func processUser() {
    var user User  // Declared with zero value
    
    if someCondition {
        user = User{Name: "Alice", Age: 25}
    } else {
        user = User{Name: "Bob", Age: 30}
    }
    
    // Use user...
}
```

### When to Use Each Method

| Method | Use Case | Example |
|--------|----------|---------|
| `var name type = value` | When type clarity is important | `var userID int64 = 12345` |
| `var name = value` | Package-level variables | `var defaultConfig = Config{...}` |
| `name := value` | Most common in functions | `result := calculate()` |
| `var name type` | When initial value comes later | `var err error` |

---

## Data Types in Detail

### Integer Types

#### Signed Integers
```go
var i8 int8 = 127           // -128 to 127 (1 byte)
var i16 int16 = 32767       // -32,768 to 32,767 (2 bytes)
var i32 int32 = 2147483647  // -2^31 to 2^31-1 (4 bytes)
var i64 int64 = 9223372036854775807  // -2^63 to 2^63-1 (8 bytes)

// Platform-dependent size (32 or 64 bit)
var i int = 42

// Common use cases
var (
    statusCode    int8  = 1      // Small range values
    port          int16 = 8080   // Port numbers
    timestamp     int32 = 1640995200  // Unix timestamps (until 2038)
    largeNumber   int64 = 1234567890123456789  // Large numbers
    generalPurpose int  = 42     // Most common
)
```

#### Unsigned Integers
```go
var ui8 uint8 = 255            // 0 to 255 (1 byte)
var ui16 uint16 = 65535        // 0 to 65,535 (2 bytes)
var ui32 uint32 = 4294967295   // 0 to 2^32-1 (4 bytes)
var ui64 uint64 = 18446744073709551615  // 0 to 2^64-1 (8 bytes)

// Platform-dependent size
var ui uint = 42

// Special aliases
var b byte = 255      // Alias for uint8
var r rune = 'A'      // Alias for int32 (Unicode code point)

// Common use cases
var (
    pixelValue byte   = 128      // Image processing
    character  rune   = '🌟'     // Unicode characters
    fileSize   uint64 = 1048576  // File sizes
    index      uint   = 0        // Array indices (when negative not possible)
)
```

#### Integer Operations and Gotchas
```go
// Overflow behavior
var maxInt8 int8 = 127
maxInt8++  // Becomes -128 (overflow wraps around)

// Mixed type operations require conversion
var a int32 = 100
var b int64 = 200
// var c = a + b  // ERROR: cannot mix types
var c = int64(a) + b  // OK: explicit conversion

// Division behavior
var x int = 7
var y int = 3
fmt.Println(x / y)  // 2 (integer division, no decimals)

// Modulo operation
fmt.Println(x % y)  // 1 (remainder)
```

### Floating-Point Types

#### Basic Float Types
```go
var f32 float32 = 3.14159    // ~7 decimal digits precision
var f64 float64 = 3.14159265358979323846  // ~15 decimal digits precision

// Default float type is float64
var pi = 3.14159  // float64

// Scientific notation
var avogadro = 6.022e23      // 6.022 × 10^23
var planck = 6.626e-34       // 6.626 × 10^-34
var large float32 = 1.23e10  // 1.23 × 10^10
```

#### Float Operations and Precision
```go
// Precision limitations
var f1 float32 = 0.1
var f2 float32 = 0.2
fmt.Printf("%.10f\n", f1+f2)  // Not exactly 0.3!

// Special float values
import "math"

var positiveInf = math.Inf(1)   // +Inf
var negativeInf = math.Inf(-1)  // -Inf
var notANumber = math.NaN()     // NaN

// Checking special values
fmt.Println(math.IsInf(positiveInf, 1))  // true
fmt.Println(math.IsNaN(notANumber))      // true

// Common float operations
var radius = 5.0
area := math.Pi * radius * radius
circumference := 2 * math.Pi * radius
```

### String Type

#### String Basics
```go
// String literals
var name string = "Alice"
var greeting = "Hello, World!"
var empty string  // "" (empty string)

// Raw string literals (backticks)
var multiline = `This is a
multi-line
string that preserves
    whitespace and formatting`

var path = `C:\Users\Alice\Documents\file.txt`  // No need to escape backslashes

// String with escape sequences
var escaped = "Line 1\nLine 2\tTabbed\r\nWindows line ending"
var quoted = "She said, \"Hello!\""
var unicode = "Unicode: \u2603 \U0001F600"  // ☃ 😀
```

#### String Operations
```go
// String concatenation
firstName := "John"
lastName := "Doe"
fullName := firstName + " " + lastName

// String length (in bytes, not characters!)
name := "Alice"
fmt.Println(len(name))  // 5

// Unicode string length
unicodeName := "Åse"  // Contains non-ASCII character
fmt.Println(len(unicodeName))        // 4 (bytes)
fmt.Println(len([]rune(unicodeName))) // 3 (characters)

// String indexing (returns bytes)
fmt.Println(name[0])    // 65 (ASCII value of 'A')
fmt.Println(string(name[0]))  // "A"

// String slicing
fmt.Println(name[1:4])  // "lic"
fmt.Println(name[:3])   // "Ali"
fmt.Println(name[2:])   // "ice"
```

### Boolean Type
```go
var isActive bool = true
var isComplete bool = false
var defaultBool bool  // false (zero value)

// Boolean from expressions
age := 25
isAdult := age >= 18        // true
isTeenager := age >= 13 && age < 20  // false
isNotChild := age > 12      // true

// Boolean operations
a := true
b := false

// Logical AND
fmt.Println(a && b)   // false
// Logical OR
fmt.Println(a || b)   // true
// Logical NOT
fmt.Println(!a)  // false
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

---

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

### Practical Use of Zero Values
```go
// Zero values make code safer
func processNumbers(numbers []int) int {
    var sum int  // Starts at 0, safe to use immediately
    for _, num := range numbers {
        sum += num
    }
    return sum
}

// Counter example
type Counter struct {
    value int  // Starts at 0
}

func (c *Counter) Increment() {
    c.value++  // Safe even without initialization
}
```

---

## Type Inference

Go uses the **type of the initializing expression** to determine the variable's type.

### Basic Type Inference
```go
// Literals determine type
var a = 42          // int (default for integer literals)
var b = 42.0        // float64 (default for float literals)
var c = "hello"     // string
var d = true        // bool
var e = 'A'         // rune (int32)

// Function return types determine variable type
func getString() string { return "hello" }
func getInt() int { return 42 }

var s = getString()  // string
var i = getInt()     // int
```

### Complex Type Inference
```go
// Slice literals
var numbers = []int{1, 2, 3}           // []int
var mixed = []interface{}{1, "hello"}   // []interface{}

// Map literals
var ages = map[string]int{             // map[string]int
    "Alice": 30,
    "Bob":   25,
}

// Struct literals
type Person struct {
    Name string
    Age  int
}
var person = Person{Name: "Alice", Age: 30}  // Person
```

---

## Type Conversion

### Explicit Type Conversion
Go requires **explicit conversion** between different types.

### Numeric Conversions
```go
// Basic numeric conversions
var i int = 42
var f float64 = float64(i)  // int to float64
var u uint = uint(i)        // int to uint
var i32 int32 = int32(i)    // int to int32

// Precision loss warnings
var bigFloat float64 = 3.14159
var smallFloat float32 = float32(bigFloat)  // May lose precision

var largeInt int64 = 1000000
var smallInt int8 = int8(largeInt)  // Will overflow! Results in unexpected value
```

### String Conversions

#### String to Numbers
```go
import "strconv"

// String to int
str := "42"
num, err := strconv.Atoi(str)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Number: %d\n", num)  // 42
}

// String to float
floatStr := "3.14159"
f, err := strconv.ParseFloat(floatStr, 64)  // 64-bit precision
if err == nil {
    fmt.Printf("Float: %f\n", f)
}

// String to bool
boolStr := "true"
b, err := strconv.ParseBool(boolStr)  // Accepts: true, false, 1, 0, etc.
if err == nil {
    fmt.Printf("Bool: %t\n", b)
}
```

#### Numbers to String
```go
import "strconv"

// Int to string
num := 42
str := strconv.Itoa(num)  // "42"

// Int to string with base
str = strconv.FormatInt(int64(num), 16)  // "2a" (hexadecimal)

// Float to string
f := 3.14159
str = strconv.FormatFloat(f, 'f', 2, 64)  // "3.14" (2 decimal places)

// Bool to string
b := true
str = strconv.FormatBool(b)  // "true"
```

---

## Variable Scope

### Package-Level Scope
```go
package main

import "fmt"

// Package-level variables (global scope)
var globalCounter int = 0
var appName string = "MyApp"
var isDebugMode bool = true

// These are accessible from any function in the same package
func incrementCounter() {
    globalCounter++  // Can access and modify
}

func printAppInfo() {
    fmt.Printf("App: %s, Debug: %t, Counter: %d\n", 
        appName, isDebugMode, globalCounter)
}
```

### Function Scope
```go
func processOrder() {
    // Function-level variables
    orderID := "ORD-12345"
    totalAmount := 99.99
    isValid := true
    
    // These variables exist only within this function
    fmt.Printf("Processing order %s for $%.2f\n", orderID, totalAmount)
    
    if isValid {
        // Can access function-level variables from inner blocks
        fmt.Printf("Order %s is valid\n", orderID)
    }
}  // orderID, totalAmount, isValid are destroyed here
```

### Block Scope
```go
func demonstrateBlockScope() {
    outerVar := "I'm in the outer scope"
    
    if true {
        // Block scope - can access outer variables
        fmt.Println(outerVar)  // OK
        
        // Variables declared in block
        innerVar := "I'm in the inner scope"
        
        if true {
            // Nested block scope
            fmt.Println(outerVar)  // OK
            fmt.Println(innerVar)  // OK
            
            nestedVar := "I'm in the nested scope"
            fmt.Println(nestedVar)  // OK
        }
        // nestedVar is not accessible here
        
        fmt.Println(innerVar)  // OK
    }
    // innerVar is not accessible here
    
    fmt.Println(outerVar)  // OK
}
```

### Variable Shadowing
```go
var globalVar string = "global"

func shadowingExample() {
    fmt.Println(globalVar)  // "global"
    
    // Shadow the global variable
    globalVar := "function level"  // This creates a new variable!
    fmt.Println(globalVar)  // "function level"
    
    if true {
        // Shadow the function-level variable
        globalVar := "block level"
        fmt.Println(globalVar)  // "block level"
    }
    
    fmt.Println(globalVar)  // "function level" (block shadow gone)
}

func afterShadowing() {
    fmt.Println(globalVar)  // "global" (original unchanged)
}
```

### Common Scope Issues

#### Short Declaration Gotcha
```go
func shortDeclGotcha() {
    var err error
    
    if someCondition {
        // This creates a NEW variable, doesn't assign to outer err!
        result, err := someFunction()  // err is shadowed
        fmt.Println(result)
    }
    
    // The outer err is still nil!
    if err != nil {
        // This condition might not work as expected
    }
}

// Correct approach
func shortDeclCorrect() {
    var err error
    var result string
    
    if someCondition {
        result, err = someFunction()  // Now assigns to outer err
    }
    
    if err != nil {
        // This works as expected
    }
}
```

---

## Constants

### Basic Constants
```go
const pi = 3.14159
const greeting = "Hello"
const isDebug = false

// Multiple constants
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
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

### Advanced Iota Usage
```go
// Complex iota expressions
const (
    ReadPerm = 1 << iota    // 1 (binary: 001)
    WritePerm               // 2 (binary: 010)
    ExecPerm                // 4 (binary: 100)
)

// Reset iota in new const block
const (
    Apple = iota    // 0
    Orange          // 1
    Banana          // 2
)

const (
    Red = iota      // 0 (iota resets)
    Green           // 1
    Blue            // 2
)

// String constants with iota
const (
    StatusActive = "active"
    StatusInactive = "inactive"
    StatusPending = "pending"
)
```

---

## Memory Management

### Stack vs Heap Allocation
```go
// Stack allocation (local variables)
func stackExample() {
    var x int = 42      // Allocated on stack
    var arr [100]int    // Allocated on stack
    
    // These are automatically cleaned up when function returns
}

// Heap allocation (dynamic memory)
func heapExample() *int {
    var x int = 42
    return &x           // x escapes to heap, returned as pointer
}

// Slice/map allocation
func dynamicExample() {
    slice := make([]int, 1000)      // Allocated on heap
    m := make(map[string]int)       // Allocated on heap
    
    // Go's garbage collector will clean these up
}
```

### Escape Analysis
```go
// Variable stays on stack
func staysOnStack() {
    var x int = 42
    fmt.Println(x)  // x doesn't escape
}

// Variable escapes to heap
func escapesToHeap() *int {
    var x int = 42
    return &x       // x escapes because we return its address
}

// Large structures may go to heap
func largeStruct() {
    var large [10000]int  // May be allocated on heap due to size
    _ = large
}
```

### Memory Efficiency Tips
```go
// Prefer stack allocation when possible
func efficientCode() {
    // Good: small, local variables
    var counter int
    var name string
    
    // Be careful with large arrays
    var huge [1000000]int  // Consider using slice instead
    
    // Use slices with known capacity
    numbers := make([]int, 0, 100)  // Avoids reallocations
    
    // Reuse slices when possible
    numbers = numbers[:0]  // Reset length, keep capacity
}
```

---

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

// Boolean variables should be questions
var isReady bool
var hasPermission bool
var canDelete bool
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

// Group related variables
var (
    serverHost = "localhost"
    serverPort = 8080
    serverTimeout = 30 * time.Second
)
```

### Error Handling
```go
// Check errors immediately
value, err := someFunction()
if err != nil {
    return fmt.Errorf("failed to get value: %w", err)
}

// Use blank identifier for unused values
_, err := someFunction()
if err != nil {
    return err
}

// Don't ignore errors
result := dangerousFunction()  // BAD: ignoring potential error

result, err := dangerousFunction()  // GOOD: handling error
if err != nil {
    log.Printf("Warning: %v", err)
}
```

### Performance Considerations
```go
// Avoid unnecessary allocations
func inefficient() {
    var result string