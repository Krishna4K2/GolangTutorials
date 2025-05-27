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

#### When to Use float32 vs float64
```go
// Use float32 when:
// - Memory is critical (large arrays)
// - Working with graphics APIs
// - Interfacing with C libraries that use float

var vertices []float32 = []float32{
    -0.5, -0.5, 0.0,  // Vertex 1
     0.5, -0.5, 0.0,  // Vertex 2
     0.0,  0.5, 0.0,  // Vertex 3
}

// Use float64 for:
// - Mathematical calculations requiring precision
// - Financial calculations
// - Scientific computing
// - Most general-purpose floating-point math

var account struct {
    balance     float64  // Monetary values need precision
    interestRate float64
}
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

#### String Iteration
```go
text := "Hello, 世界"

// Iterate over bytes
for i := 0; i < len(text); i++ {
    fmt.Printf("Byte %d: %c\n", i, text[i])
}

// Iterate over runes (Unicode code points)
for i, r := range text {
    fmt.Printf("Rune at %d: %c (U+%04X)\n", i, r, r)
}

// Convert to rune slice for character manipulation
runes := []rune(text)
fmt.Println(len(runes))  // Actual character count
```

#### String Immutability
```go
name := "Alice"
// name[0] = 'B'  // ERROR: cannot assign to string index

// To modify, convert to rune slice or byte slice
runes := []rune(name)
runes[0] = 'B'
newName := string(runes)  // "Blice"

// Or use string building
import "strings"

var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
result := builder.String()  // "Hello World"
```

### Boolean Type

#### Boolean Basics
```go
var isActive bool = true
var isComplete bool = false
var defaultBool bool  // false (zero value)

// Boolean from expressions
age := 25
isAdult := age >= 18        // true
isTeenager := age >= 13 && age < 20  // false
isNotChild := age > 12      // true
```

#### Boolean Operations
```go
a := true
b := false

// Logical AND
fmt.Println(a && b)   // false
fmt.Println(a && true)  // true

// Logical OR
fmt.Println(a || b)   // true
fmt.Println(false || false)  // false

// Logical NOT
fmt.Println(!a)  // false
fmt.Println(!b)  // true

// Short-circuit evaluation
func expensiveCheck() bool {
    fmt.Println("Expensive check called")
    return true
}

result := false && expensiveCheck()  // expensiveCheck() NOT called
result = true || expensiveCheck()    // expensiveCheck() NOT called
```

#### Boolean in Conditionals
```go
// Direct boolean usage
isReady := true
if isReady {
    fmt.Println("Ready to proceed")
}

// Boolean expressions
score := 85
if score >= 80 && score < 90 {
    fmt.Println("B grade")
}

// Boolean variables from conditions
hasPermission := user.Role == "admin" || user.ID == resource.OwnerID
canDelete := hasPermission && resource.Status != "locked"

if canDelete {
    // Perform deletion
}
```

### Complex Types (Advanced)
```go
// Complex numbers (rarely used in typical applications)
var c64 complex64 = 1 + 2i
var c128 complex128 = complex(3.0, 4.0)  // 3 + 4i

// Extract real and imaginary parts
import "math/cmplx"

fmt.Println(real(c128))      // 3
fmt.Println(imag(c128))      // 4
fmt.Println(cmplx.Abs(c128)) // Magnitude: 5
```

---

## Zero Values

### What are Zero Values?
Zero values are the default values that variables receive when declared without explicit initialization. This is a key safety feature in Go - **no variable is ever uninitialized**.

### Zero Values by Type
```go
// Numeric types
var i int           // 0
var i8 int8         // 0
var ui uint         // 0
var f float64       // 0.0
var c complex128    // (0+0i)

// Boolean
var b bool          // false

// String
var s string        // "" (empty string)

// Pointers, slices, maps, channels, functions, interfaces
var ptr *int        // nil
var slice []int     // nil
var m map[string]int // nil
var ch chan int     // nil
var fn func()       // nil
var iface interface{} // nil

// Arrays get zero values for all elements
var arr [3]int      // [0, 0, 0]

// Structs get zero values for all fields
type Person struct {
    Name string
    Age  int
    Active bool
}
var p Person        // {Name: "", Age: 0, Active: false}
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

// Error handling pattern
func divide(a, b float64) (float64, error) {
    var result float64  // 0.0
    var err error      // nil
    
    if b == 0 {
        err = errors.New("division by zero")
        return result, err  // Returns 0.0, error
    }
    
    result = a / b
    return result, err  // Returns result, nil
}
```

### Testing for Zero Values
```go
// Numeric zero values
var num int
if num == 0 {
    fmt.Println("num is zero")
}

// String zero value
var str string
if str == "" {
    fmt.Println("str is empty")
}

// Boolean zero value
var flag bool
if !flag {  // or flag == false
    fmt.Println("flag is false")
}

// Pointer/slice/map zero values
var ptr *int
var slice []int
var m map[string]int

if ptr == nil {
    fmt.Println("ptr is nil")
}
if slice == nil {
    fmt.Println("slice is nil")
}
if m == nil {
    fmt.Println("map is nil")
}

// Checking slice/map length is also safe with nil
fmt.Println(len(slice))  // 0 (safe with nil slice)
fmt.Println(len(m))      // 0 (safe with nil map)
```

---

## Type Inference

### How Go Infers Types
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

### Numeric Literal Type Inference
```go
// Integer literals
var small = 10      // int
var large = 1000000 // int

// Float literals
var pi = 3.14       // float64
var precise = 3.141592653589793238462643383279  // float64

// Scientific notation
var avogadro = 6.022e23  // float64
var tiny = 1e-10         // float64

// Hexadecimal, octal, binary
var hex = 0xFF           // int (255 in decimal)
var oct = 0o755          // int (493 in decimal)
var bin = 0b1010         // int (10 in decimal)
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

// Array literals
var arr = [3]int{1, 2, 3}              // [3]int
var autoSize = [...]string{"a", "b"}    // [2]string
```

### Type Inference with Operations
```go
var a = 10      // int
var b = 20      // int
var sum = a + b // int (result of int + int)

var x = 3.14    // float64
var y = 2.0     // float64
var product = x * y  // float64

// Mixed operations require explicit conversion
var intVal = 10
var floatVal = 3.14
// var result = intVal + floatVal  // ERROR: type mismatch
var result = float64(intVal) + floatVal  // OK: float64
```

### When Type Inference Might Surprise You
```go
// Default types might not be what you expect
var smallInt = 10    // int (not int8, even if 10 fits in int8)
var precision = 3.14 // float64 (not float32)

// Constants vs variables
const exactPi = 3.14159265358979323846
var approxPi = 3.14159265358979323846  // Still float64, but may lose precision

// Interface{} inference
var anything = 42        // int
var boxed interface{} = 42  // interface{} containing int

// Function type inference
var fn = func(x int) int { return x * 2 }  // func(int) int
```

---

## Type Conversion

### Explicit Type Conversion
Go requires **explicit conversion** between different types, even compatible ones.

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

// Safe conversion with checking
func safeIntToInt8(i int) (int8, error) {
    if i < -128 || i > 127 {
        return 0, fmt.Errorf("value %d out of range for int8", i)
    }
    return int8(i), nil
}
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

// String to int with base
hexStr := "FF"
num, err = strconv.ParseInt(hexStr, 16, 64)  // base 16, 64-bit result
fmt.Printf("Hex FF = %d\n", num)  // 255

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
str = strconv.FormatInt(int64(num), 2)   // "101010" (binary)

// Float to string
f := 3.14159
str = strconv.FormatFloat(f, 'f', 2, 64)  // "3.14" (2 decimal places)
str = strconv.FormatFloat(f, 'e', -1, 64) // Scientific notation

// Bool to string
b := true
str = strconv.FormatBool(b)  // "true"
```

#### String and Rune/Byte Conversions
```go
// String to []byte
str := "Hello, 世界"
bytes := []byte(str)
fmt.Printf("Bytes: %v\n", bytes)

// []byte to string
newStr := string(bytes)
fmt.Printf("String: %s\n", newStr)

// String to []rune (for Unicode handling)
runes := []rune(str)
fmt.Printf("Runes: %v\n", runes)
fmt.Printf("Character count: %d\n", len(runes))

// []rune to string
backToStr := string(runes)
fmt.Printf("Back to string: %s\n", backToStr)

// Single rune to string
var r rune = '世'
runeStr := string(r)  // "世"

// Single byte to string (careful with Unicode!)
var by byte = 65
byteStr := string(by)  // "A"
```

### Interface Conversions

#### Type Assertions
```go
var i interface{} = 42

// Type assertion (panics if wrong type)
num := i.(int)
fmt.Printf("Number: %d\n", num)

// Safe type assertion
num, ok := i.(int)
if ok {
    fmt.Printf("Number: %d\n", num)
} else {
    fmt.Println("Not an int")
}

// Type switch
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
case bool:
    fmt.Printf("Boolean: %t\n", v)
default:
    fmt.Printf("Unknown type: %T\n", v)
}
```

### Conversion Gotchas and Best Practices
```go
// Overflow example
var bigNum int64 = 300
var smallNum int8 = int8(bigNum)
fmt.Printf("300 as int8: %d\n", smallNum)  // 44 (300 - 256)

// Float precision loss
var precise float64 = 1.23456789012345
var lessShort float32 = float32(precise)
fmt.Printf("Original: %.15f\n", precise)
fmt.Printf("Float32:  %.15f\n", float64(lessShort))

// String conversion gotcha
var num int = 65
var wrongStr string = string(num)  // "A" (ASCII), not "65"
var correctStr string = strconv.Itoa(num)  // "65"

// Safe conversion function example
func convertStringToInt(s string) (int, error) {
    num, err := strconv.Atoi(s)
    if err != nil {
        return 0, fmt.Errorf("cannot convert '%s' to integer: %w", s, err)
    }
    return num, nil
}
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

### Loop Scope
```go
func loopScopeExample() {
    // Loop variable scope
    for i := 0; i < 3; i++ {
        // i exists only within the loop
        fmt.Printf("Loop iteration: %d\n", i)
        
        // Variables declared in loop body
        temp := i * 2
        fmt.Printf("Temp value: %d\n", temp)
    }
    // i and temp are not accessible here
    
    // Range loop scope
    numbers := []int{10, 20, 30}
    for index, value := range numbers {
        // index and value exist only within the loop
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    // index and value are not accessible here
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

#### Loop Variable Capture Problem
```go
// WRONG - Common mistake with closures
func wrongClosureExample() {
    var funcs []func()
    
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() {
            fmt.Printf("Wrong: %d\n", i)  // All will print 3!
        })
    }
    
    for _, fn := range funcs {
        fn()  // Prints "Wrong: 3" three times
    }
}

// CORRECT - Capture loop variable properly
func correctClosureExample() {
    var funcs []func()
    
    for i := 0; i < 3; i++ {
        i := i  // Create new variable in loop scope
        funcs = append(funcs, func() {
            fmt.Printf("Correct: %d\n", i)
        })
    }
    
    for _, fn := range funcs {
        fn()  // Prints "Correct: 0", "Correct: 1", "Correct: 2"
    }
}
```

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
        result, err = someFunction()