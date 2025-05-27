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
    for i := 0; i < 1000; i++ {
        result += strconv.Itoa(i)  // Creates new string each time
    }
    return result
}

func efficient() string {
    var builder strings.Builder
    for i := 0; i < 1000; i++ {
        builder.WriteString(strconv.Itoa(i))  // More efficient
    }
    return builder.String()
}

// Pre-allocate slices when size is known
func inefficientSlice() []int {
    var numbers []int
    for i := 0; i < 1000; i++ {
        numbers = append(numbers, i)  // May cause multiple reallocations
    }
    return numbers
}

func efficientSlice() []int {
    numbers := make([]int, 0, 1000)  // Pre-allocate capacity
    for i := 0; i < 1000; i++ {
        numbers = append(numbers, i)  // No reallocations needed
    }
    return numbers
}

// Reuse variables when possible
func processItems(items []string) {
    var result string  // Declare once outside loop
    for _, item := range items {
        result = strings.ToUpper(item)  // Reuse variable
        fmt.Println(result)
    }
}
```

---

## Common Patterns

### Initialization Patterns
```go
// Factory pattern
func NewUser(name string, age int) *User {
    return &User{
        Name:      name,
        Age:       age,
        CreatedAt: time.Now(),
        IsActive:  true,  // Default value
    }
}

// Builder pattern
type ConfigBuilder struct {
    config Config
}

func NewConfigBuilder() *ConfigBuilder {
    return &ConfigBuilder{
        config: Config{
            Timeout: 30 * time.Second,  // Default
            MaxRetries: 3,              // Default
        },
    }
}

func (cb *ConfigBuilder) WithTimeout(timeout time.Duration) *ConfigBuilder {
    cb.config.Timeout = timeout
    return cb
}

func (cb *ConfigBuilder) Build() Config {
    return cb.config
}

// Usage
config := NewConfigBuilder().
    WithTimeout(10 * time.Second).
    Build()
```

### Option Pattern
```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(options ...ServerOption) *Server {
    server := &Server{
        host:    "localhost",  // Default
        port:    8080,         // Default
        timeout: 30 * time.Second,  // Default
    }
    
    for _, option := range options {
        option(server)
    }
    
    return server
}

// Usage
server := NewServer(
    WithHost("example.com"),
    WithPort(9000),
)
```

### Error Handling Patterns
```go
// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Result and error pattern
func processFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    data, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    
    return data, nil
}

// Early return pattern
func validateUser(user User) error {
    if user.Name == "" {
        return errors.New("name is required")
    }
    
    if user.Age < 0 {
        return errors.New("age must be positive")
    }
    
    if user.Email == "" {
        return errors.New("email is required")
    }
    
    return nil  // All validations passed
}
```

### Concurrency Patterns
```go
// Channel communication
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        result := expensiveOperation(job)
        results <- result
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

// Mutex for shared state
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

---

## Examples and Use Cases

### Configuration Management
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

// Configuration structure
type Config struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Logging  LoggingConfig  `json:"logging"`
}

type ServerConfig struct {
    Host         string        `json:"host"`
    Port         int           `json:"port"`
    ReadTimeout  time.Duration `json:"read_timeout"`
    WriteTimeout time.Duration `json:"write_timeout"`
}

type DatabaseConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Database string `json:"database"`
}

type LoggingConfig struct {
    Level  string `json:"level"`
    Format string `json:"format"`
    Output string `json:"output"`
}

// Default configuration values
var defaultConfig = Config{
    Server: ServerConfig{
        Host:         "localhost",
        Port:         8080,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,
    },
    Database: DatabaseConfig{
        Host:     "localhost",
        Port:     5432,
        Username: "app_user",
        Database: "app_db",
    },
    Logging: LoggingConfig{
        Level:  "info",
        Format: "json",
        Output: "stdout",
    },
}

// Load configuration from file or use defaults
func LoadConfig(filename string) (*Config, error) {
    config := defaultConfig  // Start with defaults
    
    // Try to read config file
    data, err := os.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Printf("Config file %s not found, using defaults\n", filename)
            return &config, nil
        }
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    // Parse JSON config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    // Override with environment variables
    if host := os.Getenv("SERVER_HOST"); host != "" {
        config.Server.Host = host
    }
    if port := os.Getenv("SERVER_PORT"); port != "" {
        var p int
        if _, err := fmt.Sscanf(port, "%d", &p); err == nil {
            config.Server.Port = p
        }
    }
    
    return &config, nil
}

// Validate configuration
func (c *Config) Validate() error {
    if c.Server.Port <= 0 || c.Server.Port > 65535 {
        return fmt.Errorf("invalid server port: %d", c.Server.Port)
    }
    
    if c.Database.Username == "" {
        return fmt.Errorf("database username is required")
    }
    
    validLevels := map[string]bool{
        "debug": true, "info": true, "warn": true, "error": true,
    }
    if !validLevels[c.Logging.Level] {
        return fmt.Errorf("invalid log level: %s", c.Logging.Level)
    }
    
    return nil
}

func main() {
    config, err := LoadConfig("app.json")
    if err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        return
    }
    
    if err := config.Validate(); err != nil {
        fmt.Printf("Invalid config: %v\n", err)
        return
    }
    
    fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("Database: %s@%s:%d/%s\n", 
        config.Database.Username, config.Database.Host, 
        config.Database.Port, config.Database.Database)
}
```

### Data Processing Pipeline
```go
package main

import (
    "fmt"
    "strings"
    "time"
)

// Record represents a data record to process
type Record struct {
    ID        int
    Name      string
    Email     string
    Status    string
    CreatedAt time.Time
}

// Processing statistics
type ProcessingStats struct {
    TotalRecords   int
    ValidRecords   int
    InvalidRecords int
    ProcessedAt    time.Time
    Duration       time.Duration
}

// Data processor with configurable options
type DataProcessor struct {
    emailDomain     string
    requiredFields  []string
    statusWhitelist []string
    batchSize       int
}

// Create new processor with defaults
func NewDataProcessor() *DataProcessor {
    return &DataProcessor{
        emailDomain:     "@company.com",
        requiredFields:  []string{"Name", "Email"},
        statusWhitelist: []string{"active", "pending", "inactive"},
        batchSize:       100,
    }
}

// Process records with validation and transformation
func (dp *DataProcessor) ProcessRecords(records []Record) (*ProcessingStats, []Record, error) {
    startTime := time.Now()
    
    var validRecords []Record
    var invalidCount int
    
    // Process in batches
    for i := 0; i < len(records); i += dp.batchSize {
        end := i + dp.batchSize
        if end > len(records) {
            end = len(records)
        }
        
        batch := records[i:end]
        fmt.Printf("Processing batch %d-%d...\n", i+1, end)
        
        for _, record := range batch {
            if dp.validateRecord(record) {
                processedRecord := dp.transformRecord(record)
                validRecords = append(validRecords, processedRecord)
            } else {
                invalidCount++
            }
        }
    }
    
    stats := &ProcessingStats{
        TotalRecords:   len(records),
        ValidRecords:   len(validRecords),
        InvalidRecords: invalidCount,
        ProcessedAt:    startTime,
        Duration:       time.Since(startTime),
    }
    
    return stats, validRecords, nil
}

// Validate record according to business rules
func (dp *DataProcessor) validateRecord(record Record) bool {
    // Check required fields
    if strings.TrimSpace(record.Name) == "" {
        return false
    }
    
    if strings.TrimSpace(record.Email) == "" {
        return false
    }
    
    // Validate email format
    if !strings.Contains(record.Email, "@") {
        return false
    }
    
    // Check status whitelist
    statusValid := false
    for _, validStatus := range dp.statusWhitelist {
        if record.Status == validStatus {
            statusValid = true
            break
        }
    }
    
    if !statusValid {
        return false
    }
    
    return true
}

// Transform record (normalize data)
func (dp *DataProcessor) transformRecord(record Record) Record {
    // Normalize name (title case)
    record.Name = strings.Title(strings.ToLower(strings.TrimSpace(record.Name)))
    
    // Normalize email (lowercase)
    record.Email = strings.ToLower(strings.TrimSpace(record.Email))
    
    // Ensure company domain
    if !strings.HasSuffix(record.Email, dp.emailDomain) {
        parts := strings.Split(record.Email, "@")
        if len(parts) == 2 {
            record.Email = parts[0] + dp.emailDomain
        }
    }
    
    // Normalize status
    record.Status = strings.ToLower(record.Status)
    
    return record
}

func main() {
    // Sample data
    rawRecords := []Record{
        {ID: 1, Name: "john doe", Email: "JOHN@EXAMPLE.COM", Status: "Active", CreatedAt: time.Now()},
        {ID: 2, Name: "", Email: "invalid", Status: "pending", CreatedAt: time.Now()},
        {ID: 3, Name: "jane smith", Email: "jane@test.com", Status: "INACTIVE", CreatedAt: time.Now()},
        {ID: 4, Name: "bob wilson", Email: "bob@company.com", Status: "active", CreatedAt: time.Now()},
    }
    
    processor := NewDataProcessor()
    stats, validRecords, err := processor.ProcessRecords(rawRecords)
    
    if err != nil {
        fmt.Printf("Error processing records: %v\n", err)
        return
    }
    
    // Print statistics
    fmt.Printf("\nProcessing completed in %v\n", stats.Duration)
    fmt.Printf("Total records: %d\n", stats.TotalRecords)
    fmt.Printf("Valid records: %d\n", stats.ValidRecords)
    fmt.Printf("Invalid records: %d\n", stats.InvalidRecords)
    fmt.Printf("Success rate: %.2f%%\n", 
        float64(stats.ValidRecords)/float64(stats.TotalRecords)*100)
    
    // Print valid records
    fmt.Println("\nValid records after processing:")
    for _, record := range validRecords {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Status: %s\n",
            record.ID, record.Name, record.Email, record.Status)
    }
}
```

### HTTP API Server with Variable Management
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
    "time"
)

// User represents a user in our system
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Server holds the application state
type Server struct {
    users    map[int]*User
    nextID   int
    mutex    sync.RWMutex
    host     string
    port     int
    timeout  time.Duration
    started  time.Time
}

// ServerStats represents server statistics
type ServerStats struct {
    TotalUsers    int           `json:"total_users"`
    ActiveUsers   int           `json:"active_users"`
    InactiveUsers int           `json:"inactive_users"`
    Uptime        time.Duration `json:"uptime"`
    StartedAt     time.Time     `json:"started_at"`
}

// Create new server instance
func NewServer(host string, port int) *Server {
    return &Server{
        users:   make(map[int]*User),
        nextID:  1,
        host:    host,
        port:    port,
        timeout: 30 * time.Second,
        started: time.Now(),
    }
}

// Add sample data
func (s *Server) seedData() {
    sampleUsers := []*User{
        {Name: "Alice Johnson", Email: "alice@company.com", Status: "active"},
        {Name: "Bob Smith", Email: "bob@company.com", Status: "active"},
        {Name: "Carol Davis", Email: "carol@company.com", Status: "inactive"},
    }
    
    for _, user := range sampleUsers {
        s.createUser(user)
    }
}

// Create user (thread-safe)
func (s *Server) createUser(user *User) *User {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    user.ID = s.nextID
    s.nextID++
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    if user.Status == "" {
        user.Status = "active"
    }
    
    s.users[user.ID] = user
    return user
}

// Get user by ID (thread-safe)
func (s *Server) getUser(id int) (*User, bool) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    user, exists := s.users[id]
    return user, exists
}

// Get all users (thread-safe)
func (s *Server) getAllUsers() []*User {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    users := make([]*User, 0, len(s.users))
    for _, user := range s.users {
        users = append(users, user)
    }
    return users
}

// Update user (thread-safe)
func (s *Server) updateUser(id int, updates *User) (*User, bool) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    
    user, exists := s.users[id]
    if !exists {
        return nil, false
    }
    
    // Update fields if provided
    if updates.Name != "" {
        user.Name = updates.Name
    }
    if updates.Email != "" {
        user.Email = updates.Email
    }
    if updates.Status != "" {
        user.Status = updates.Status
    }
    
    user.UpdatedAt = time.Now()
    return user, true
}

// Get server statistics
func (s *Server) getStats() *ServerStats {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    
    var activeCount, inactiveCount int
    for _, user := range s.users {
        if user.Status == "active" {
            activeCount++
        } else {
            inactiveCount++
        }
    }
    
    return &ServerStats{
        TotalUsers:    len(s.users),
        ActiveUsers:   activeCount,
        InactiveUsers: inactiveCount,
        Uptime:        time.Since(s.started),
        StartedAt:     s.started,
    }
}

// HTTP Handlers

func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        s.handleGetUsers(w, r)
    case http.MethodPost:
        s.handleCreateUser(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) {
    users := s.getAllUsers()
    
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Validate required fields
    if user.Name == "" || user.Email == "" {
        http.Error(w, "Name and email are required", http.StatusBadRequest)
        return
    }
    
    createdUser := s.createUser(&user)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdUser)
}

func (s *Server) handleUserByID(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL path
    idStr := r.URL.Path[len("/users/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    switch r.Method {
    case http.MethodGet:
        s.handleGetUser(w, r, id)
    case http.MethodPut:
        s.handleUpdateUser(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request, id int) {
    user, exists := s.getUser(id)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request, id int) {
    var updates User
    if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    user, exists := s.updateUser(id, &updates)
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
    stats := s.getStats()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

// Start the server
func (s *Server) Start() error {
    // Setup routes
    http.HandleFunc("/users", s.handleUsers)
    http.HandleFunc("/users/", s.handleUserByID)
    http.HandleFunc("/stats", s.handleStats)
    
    // Add sample data
    s.seedData()
    
    addr := fmt.Sprintf("%s:%d", s.host, s.port)
    fmt.Printf("Starting server on %s\n", addr)
    fmt.Printf("Endpoints:\n")
    fmt.Printf("  GET    /users     - List all users\n")
    fmt.Printf("  POST   /users     - Create new user\n")
    fmt.Printf("  GET    /users/:id - Get user by ID\n")
    fmt.Printf("  PUT    /users/:id - Update user\n")
    fmt.Printf("  GET    /stats     - Get server statistics\n")
    
    server := &http.Server{
        Addr:         addr,
        ReadTimeout:  s.timeout,
        WriteTimeout: s.timeout,
    }
    
    return server.ListenAndServe()
}

func main() {
    // Server configuration
    var host string = "localhost"
    var port int = 8080
    
    // Override from environment if available
    if envHost := os.Getenv("HOST"); envHost != "" {
        host = envHost
    }
    if envPort := os.Getenv("PORT"); envPort != "" {
        if p, err := strconv.Atoi(envPort); err == nil {
            port = p
        }
    }
    
    server := NewServer(host, port)
    
    if err := server.Start(); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
```

### Financial Calculator with Precision Handling
```go
package main

import (
    "fmt"
    "math"
)

// Investment represents an investment calculation
type Investment struct {
    Principal       float64 // Initial amount
    AnnualRate      float64 // Annual interest rate (as decimal)
    CompoundPeriods int     // Times compounded per year
    Years           int     // Investment period in years
}

// LoanPayment represents a loan payment calculation
type LoanPayment struct {
    Principal   float64 // Loan amount
    AnnualRate  float64 // Annual interest rate (as decimal)
    Years       int     // Loan term in years
    Payments    int     // Payments per year
}

// FinancialCalculator provides various financial calculations
type FinancialCalculator struct {
    precision int // Decimal places for rounding
}

// Create new financial calculator
func NewFinancialCalculator(precision int) *FinancialCalculator {
    if precision < 0 {
        precision = 2 // Default to 2 decimal places
    }
    return &FinancialCalculator{precision: precision}
}

// Round to specified precision
func (fc *FinancialCalculator) round(value float64) float64 {
    multiplier := math.Pow10(fc.precision)
    return math.Round(value*multiplier) / multiplier
}

// Calculate compound interest
func (fc *FinancialCalculator) CompoundInterest(inv Investment) (float64, error) {
    if inv.Principal <= 0 {
        return 0, fmt.Errorf("principal must be positive")
    }
    if inv.AnnualRate < 0 {
        return 0, fmt.Errorf("interest rate cannot be negative")
    }
    if inv.CompoundPeriods <= 0 {
        return 0, fmt.Errorf("compound periods must be positive")
    }
    if inv.Years < 0 {
        return 0, fmt.Errorf("years cannot be negative")
    }
    
    // A = P(1 + r/n)^(nt)
    rate := inv.AnnualRate / float64(inv.CompoundPeriods)
    exponent := float64(inv.CompoundPeriods * inv.Years)
    amount := inv.Principal * math.Pow(1+rate, exponent)
    
    return fc.round(amount), nil
}

// Calculate monthly loan payment
func (fc *FinancialCalculator) MonthlyPayment(loan LoanPayment) (float64, error) {
    if loan.Principal <= 0 {
        return 0, fmt.Errorf("loan amount must be positive")
    }
    if loan.AnnualRate < 0 {
        return 0, fmt.Errorf("interest rate cannot be negative")
    }
    if loan.Years <= 0 {
        return 0, fmt.Errorf("loan term must be positive")
    }
    if loan.Payments <= 0 {
        return 0, fmt.Errorf("payments per year must be positive")
    }
    
    // Handle zero interest rate
    if loan.AnnualRate == 0 {
        return fc.round(loan.Principal / float64(loan.Years*loan.Payments)), nil
    }
    
    // M = P * [r(1+r)^n] / [(1+r)^n - 1]
    monthlyRate := loan.AnnualRate / float64(loan.Payments)
    totalPayments := float64(loan.Years * loan.Payments)
    
    numerator := loan.Principal * monthlyRate * math.Pow(1+monthlyRate, totalPayments)
    denominator := math.Pow(1+monthlyRate, totalPayments) - 1
    
    payment := numerator / denominator
    return fc.round(payment), nil
}

// Calculate total interest paid on loan
func (fc *FinancialCalculator) TotalInterest(loan LoanPayment) (float64, error) {
    monthlyPayment, err := fc.MonthlyPayment(loan)
    if err != nil {
        return 0, err
    }
    
    totalPayments := float64(loan.Years * loan.Payments)
    totalPaid := monthlyPayment * totalPayments
    totalInterest := totalPaid - loan.Principal
    
    return fc.round(totalInterest), nil
}

// Calculate future value of annuity
func (fc *FinancialCalculator) AnnuityFutureValue(payment float64, annualRate float64, years int, paymentsPerYear int) (float64, error) {
    if payment <= 0 {
        return 0, fmt.Errorf("payment must be positive")
    }
    if annualRate < 0 {
        return 0, fmt.Errorf("interest rate cannot be negative")
    }
    if years <= 0 {
        return 0, fmt.Errorf("years must be positive")
    }
    if paymentsPerYear <= 0 {
        return 0, fmt.Errorf("payments per year must be positive")
    }
    
    // Handle zero interest rate
    if annualRate == 0 {
        return fc.round(payment * float64(years*paymentsPerYear)), nil
    }
    
    // FV = PMT * [((1+r)^n - 1) / r]
    periodRate := annualRate / float64(paymentsPerYear)
    totalPeriods := float64(years * paymentsPerYear)
    
    futureValue := payment * (math.Pow(1+periodRate, totalPeriods) - 1) / periodRate
    return fc.round(futureValue), nil
}

// Calculate break-even point for investment
func (fc *FinancialCalculator) BreakEvenTime(principal, targetAmount, annualRate float64, compoundPeriods int) (float64, error) {
    if principal <= 0 {
        return 0, fmt.Errorf("principal must be positive")
    }
    if targetAmount <= principal {
        return 0, fmt.Errorf("target amount must be greater than principal")
    }
    if annualRate <= 0 {
        return 0, fmt.Errorf("interest rate must be positive")
    }
    if compoundPeriods <= 0 {
        return 0, fmt.Errorf("compound periods must be positive")
    }
    
    // t = ln(A/P) / (n * ln(1 + r/n))
    rate := annualRate / float64(compoundPeriods)
    years := math.Log(targetAmount/principal) / (float64(compoundPeriods) * math.Log(1+rate))
    
    return fc.round(years), nil
}

// Investment comparison result
type ComparisonResult struct {
    InvestmentA InvestmentResult `json:"investment_a"`
    InvestmentB InvestmentResult `json:"investment_b"`
    Difference  float64          `json:"difference"`
    Better      string           `json:"better"`
}

type InvestmentResult struct {
    FinalAmount   float64 `json:"final_amount"`
    TotalInterest float64 `json:"total_interest"`
    EffectiveRate float64 `json:"effective_rate"`
}

// Compare two investments
func (fc *FinancialCalculator) CompareInvestments(invA, invB Investment) (*ComparisonResult, error) {
    amountA, err := fc.CompoundInterest(invA)
    if err != nil {
        return nil, fmt.Errorf("investment A: %w", err)
    }
    
    amountB, err := fc.CompoundInterest(invB)
    if err != nil {
        return nil, fmt.Errorf("investment B: %w", err)
    }
    
    resultA := InvestmentResult{
        FinalAmount:   amountA,
        TotalInterest: fc.round(amountA - invA.Principal),
        EffectiveRate: fc.round((amountA/invA.Principal - 1) * 100),
    }
    
    resultB := InvestmentResult{
        FinalAmount:   amountB,
        TotalInterest: fc.round(amountB - invB.Principal),
        EffectiveRate: fc.round((amountB/invB.Principal - 1) * 100),
    }
    
    difference := fc.round(math.Abs(amountA - amountB))
    better := "A"
    if amountB > amountA {
        better = "B"
    } else if amountA == amountB {
        better = "Equal"
    }
    
    return &ComparisonResult{
        InvestmentA: resultA,
        InvestmentB: resultB,
        Difference:  difference,
        Better:      better,
    }, nil
}

func main() {
    calculator := NewFinancialCalculator(2)
    
    fmt.Println("=== Financial Calculator Demo ===\n")
    
    // Compound Interest Example
    fmt.Println("1. Compound Interest Calculation")
    investment := Investment{
        Principal:       10000.0, // $10,000
        AnnualRate:      0.06,    // 6% annual rate
        CompoundPeriods: 12,      // Monthly compounding
        Years:           10,      // 10 years
    }
    
    finalAmount, err := calculator.CompoundInterest(investment)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        interestEarned := finalAmount - investment.Principal
        fmt.Printf("Initial Investment: $%.2f\n", investment.Principal)
        fmt.Printf("Final Amount: $%.2f\n", finalAmount)
        fmt.Printf("Interest Earned: $%.2f\n", interestEarned)
        fmt.Printf("Effective Rate: %.2f%%\n", (finalAmount/investment.Principal-1)*100)
    }
    
    fmt.Println()
    
    // Loan Payment Example
    fmt.Println("2. Loan Payment Calculation")
    loan := LoanPayment{
        Principal:  250000.0, // $250,000 mortgage
        AnnualRate: 0.045,    // 4.5% annual rate
        Years:      30,       // 30-year term
        Payments:   12,       // Monthly payments
    }
    
    monthlyPayment, err := calculator.MonthlyPayment(loan)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        totalInterest, _ := calculator.TotalInterest(loan)
        totalPaid := monthlyPayment * float64(loan.Years*loan.Payments)
        
        fmt.Printf("Loan Amount: $%.2f\n", loan.Principal)
        fmt.Printf("Monthly Payment: $%.2f\n", monthlyPayment)
        fmt.Printf("Total Interest: $%.2f\n", totalInterest)
        fmt.Printf("Total Paid: $%.2f\n", totalPaid)
    }
    
    fmt.Println()
    
    // Investment Comparison
    fmt.Println("3. Investment Comparison")
    invA := Investment{Principal: 5000, AnnualRate: 0.05, CompoundPeriods: 1, Years: 20}   // 5% annually
    invB := Investment{Principal: 5000, AnnualRate: 0.048, CompoundPeriods: 12, Years: 20} // 4.8% monthly
    
    comparison, err := calculator.CompareInvestments(invA, invB)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Investment A (5%% annual): $%.2f (%.2f%% effective)\n", 
            comparison.InvestmentA.FinalAmount, comparison.InvestmentA.EffectiveRate)
        fmt.Printf("Investment B (4.8%% monthly): $%.2f (%.2f%% effective)\n", 
            comparison.InvestmentB.FinalAmount, comparison.InvestmentB.EffectiveRate)
        fmt.Printf("Difference: $%.