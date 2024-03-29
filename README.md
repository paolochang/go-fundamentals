# Go Fundamentals

## Index

1. [Using Types and Declarations in Go](#declarations)
   1. [Variable Declarations & Type Conversions](#variable-declarations)
   2. [Strings and Slices](#strings-and-slices)
   3. [Map](#map)
2. [Calling Functions](#calling-functions)
   1. [Declaring and Calling Function](#declaring-and-calling-function)
   2. [Reading the Command Line Using Arguments](#reading-the-command-line)
   3. [Handling Errors](#handling-errors)
   4. [Assigning Functions to Variables](#assigning-functions-to-variables)
   5. [Looping over Functions](#looping-over-functions)
   6. [Run a Single Operation](#run-a-single-operation)
   7. [Lookup Functions via a map](#lookup-functions-via-a-map)
3. [Methods and Interfaces](#methods-and-interfaces)
   1. [Declaring Your Own Types](#declaring-your-own-type)
   2. [Adding Methods to Types](#adding-methods-to-types)
   3. [Embedded Types](#embedded-types)

## <a name="declarations"></a>1. Using Types and Declarations in Go

### <a name="variable-declarations"></a>1.1. Variable Declarations & Type Conversions

#### Declarations and Print Output

```go
package main

import "fmt"

func main() {
    var x int = 10 // full variable declaration format
    var y = 20     // type isn't needed, since the default type for the numeric literal 20 is int
    var z int      // type is specified, but value isn't, initializing z to the zero value for int (0)

    fmt.Println(x, y, z)
}
```

The `Println` function in the `fmt` package allows to output the values of our variables. If the variable value has not been used at least once, the Go compiler refuses to compile the code. As the same reason, unused packages will refuse to compile and return the error message.

```go
    f := 3.2
    i := 10
    fmt.Println(f, i)
    i, j := 30, 40
    fmt.Println(f, i, j)
```

The `:=` operator allows to declare and assign a value to a variable. The `:=` operator can be used with a mix of new and old variables, as long as least one variable on the left hand side is new. In the line `i, j := 30, 40`, `i` is an already-declared variable and `j` is newly declared.

#### Type Conversions

```go
    var b byte = 30
    x = int(b)
    y = y + int(b) // must need type conversion with int(b)

    fmt.Println(x, y, z)
```

### <a name="strings-and-slices"></a>1.2. Strings and Slices

#### Strings

```go
package main

import "fmt"

func main() {
    s1 := "This is an interpreted string literal.\n\tIt must be written on a single line.\n\tEscape sequences are evaluated."
    fmt.Println(s1)

    s2 := `This is a raw string literal.
    It can be written across multiple lines.
    Escape sequences like \n are not evaluated.`
    fmt.Println(s2)
}
```

#### Slice

A `slice` is a growable list of values.

```go
package main

import "fmt"

func main() {
    var intSlice []int
    stringSlice := []string{"the", "quick", "brown", "fox"}
    zeroIntSlice := make([]int, 10)

    fmt.Println(intSlice, stringSlice, zeroIntSlice)
    fmt.Println(len(intSlice), len(stringSlice), len(zeroIntSlice))
}
```

| function               | description                              |
| ---------------------- | ---------------------------------------- |
| `make(type, length)`   | create a slice of specific length        |
| `len(slice)`           | find the length of a lice                |
| `append(slice, value)` | add new elements at the end of the slice |

```go
    midSlice := stringSlice[1:3]
    fmt.Println(midSlice)
    startSlice := stringSlice[:2]
    fmt.Println(startSlice)
    endSlice := stringSlice[2:]
    fmt.Println(endSlice)
```

**Output**:

    [quick brown]
    [the quick]
    [brown bear]

The slice expression using inside of brackets, `:`, has `starting offset` and `ending offset`

```go
    stringRunes := []rune(sunString)
    fmt.Println(len(stringRunes), stringRunes)
    roundTripRunes := string(stringRunes)
    fmt.Println(roundTripRunes)
    r := 'X'
    fmt.Println(r, string(r))
```

A `rune` is an Unicode code point, whish is a 32-bit int and rune literal is a single character within single quotes.

### <a name="map"></a>1.3. Maps

A `map` is Go's version of an associative array. It holds `key-value` pairs and is similar to the `HashMap` in Java, the `dict` in Python, or the `Hash` in Ruby.

#### Map Declaration

**ex.** [01e-maps.go](./01-Declarations/01e-maps.go) (line 6-12):

```go
    m := map[string]int{}
    m2 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    m3 := make(map[string]int)
```

**Two ways of creating a map**

- Initialize map as `m`, `m2`
- Initialize map with `make` function like `m3`

Note: Never use an uninitialized map variable. Unlike a slice, you cannot add a value to a `nil` map.

#### Comma ok Idiom

**ex.** [01e-maps.go](./01-Declarations/01e-maps.go) (line 27-34):

```go
    m["present"] = 0
    fmt.Println(m["present"], m["not present"])

    val, ok := m["present"]
    fmt.Println(val, ok)

    val, ok = m["not present"]
    fmt.Println(val, ok)
```

**Output**:

    0 true
    0 false

#### Deleting From Maps

**ex.** [01e-maps.go](./01-Declarations/01e-maps.go) (line 36-38):

```go
    delete(m, "present")
    val, ok = m["present"]
    fmt.Println(val, ok)
```

Use the delete function to remove a key-value pair from a map. The first parameter passed to the delete function is the map, and the second parameter is the key to remove. This function returns nothing. If a key that's not in the map, is passed, the function call does nothing.

## <a name="calling-functions"></a>2. Calling Functions

### <a name="declaring-and-calling-function"></a>2.1 Declaring and Calling Function

#### Example of Calling Functions

**ex.** [02a-calculator.go](./02-Functions/02a-calculator.go) (line 7-34):

```go
func main() {
    a := 20
    b := 10
    c := adder(a, b)
    d := subtractor(a, b)
    e := multiplier(a, b)
    f := divider(a, b)
}

func adder(a, b int) int {
    return a + b
}

func subtractor(a, b int) int {
    return a - b
}

func multiplier(a, b int) int {
    return a * b
}

func divider(a, b int) int {
    return a / b
}
```

### <a name="reading-the-command-line"></a>2.2. Reading the Command Line Using Arguments

Go allows to accept arguments as other program languages.

**ex.** [02b-calculator.go](./02-Functions/02b-calculator.go) (line 3-11):

```go
import (
    (...)
	"os"
	"strconv"
)

func main() {
    a, _ := strconv.Atoi(os.Args[1])
    b, _ := strconv.Atoi(os.Args[2])
    c := adder(a, b)
    (...)
}
```

- `os.Args` from `os` packages allows to use the arguments
- `strconv.Atoi` from `strconv` package convert `string` to `int`

**Command Line**:

    $ go run <program_name> <arguments>

### <a name="handling-errors"></a>2.3. Handling Errors

#### Validating the Input

**ex.** [02c-calculator.go](./02-Functions/02c-calculator.go) (line 10-26):

```go
func main() {
    if len(os.Args) != 3 {
        fmt.Println("Two integer parameters expected")
        os.Exit(1)
    }

    a, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("invalid first argument:", err)
        os.Exit(1)
    }

    b, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("invalid second argument:", err)
        os.Exit(1)
    }
    c := adder(a, b)
    (...)
}
```

- `os.Exit` ends the program early with an error code

#### Fixing the Function Signatures and Handling Panics

`division by zero` will return the `panic: runtime error`. Before the division, check the divider is not equal to zero and return error message to prevent the panic error.

**ex.** [02c-calculator.go](./02-Functions/02c-calculator.go) (line 84-89):

```go
func divider(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

#### More Data Handling Improvements

While addition, subtraction, and multiplication can't cause panic, the calculation can still return the incorrect results when it uses too large numbers. It will result in `overflow`. By adding error handler to check the numbers that are beyond the range of an `int`, it can prevents the `overflow` and `underflow`.

**Command line**:

    $ go run calculator.go 9223372036854775807 9223372036854775807

**Output**:

    -2
    0
    1
    1

**ex.** [02c-calculator.go](./02-Functions/02c-calculator.go) (line 57-63):

```go
func adder(a, b int) (int, error) {
    x := a + b
    if (x > a) != (b > 0) {
        return x, errors.New("addition out of bounds")
    }
    return x, nil
}
```

The boolean condition `(x > a) != (b > 0)` is a very clever way to check if either of two things are true:

- If the sum of the two numbers is greater than one of them, the other number must be positive
- If the sum of the two numbers is less than one of them, the other one must be negative

### <a name="assigning-functions-to-variables"></a>2.4. Assigning Functions to Variables

Rather than call the functions directly, those functions can be assigned to the local variables. Then use the local variables as a functions on the lines.

**ex.** [02d-calculator.go](./02-Functions/02d-calculator.go) (line 26-31)::

```go
func main() {
    (...)
    add := adder
    sub := subtractor
    mul := multiplier
    div := divider

    c, err := add(a, b)
    (...)
    d, err := sub(a, b)
    (...)
    e, err := mul(a, b)
    (...)
    f, err := div(a, b)
    (...)
}
```

### <a name="looping-over-functions"></a>2.5. Looping over Functions

Slice of functions can make the program shorten.

**ex.** [02e-calculator.go](./02-Functions/02e-calculator.go) (line 28-43):

```go
func main() {
    (...)
    funcs := []func(int, int) (int, error){
        adder, subtractor, multiplier, divider,
    }

    ops := []string{
        "adding", "subtracting", "multiplying", "dividing",
    }

    for i, f := range funcs {
        v, err := f(a, b)
        if err != nil {
            fmt.Println(ops[i], "failed:", err)
        } else {
            fmt.Println(v)
        }
    }
}
```

Slice of strings called `ops` contains the name of the functions. This can be used for the error reporting.

### <a name="run-a-single-operation"></a>2.6. Run a Single Operation

By using `switch`, and assigning a function to a variable, the single operation can be implemented to perform a calculation. Now the arguments accept 3 parameters, two `numbers` and a `operator`.

**ex.** [02f-calculator.go](./02-Functions/02f-calculator.go) (line 30-49):

```go
func main() {
    (...)
    var action func(int, int) (int, error)
    var actionName string

    switch op {
        case "+":
            action = adder
            actionName = "adding"
        case "-":
            action = subtractor
            actionName = "subtracting"
        case "x":
            action = multiplier
            actionName = "multiplying"
        case "/":
            action = divider
            actionName = "dividing"
        default:
            fmt.Println("Unknown operator:", op)
            os.Exit(1)
    }
    (...)
}
```

### <a name="lookup-functions-via-a-map"></a>2.7. Lookup Functions via a map

**ex.** [02g-calculator.go](./02-Functions/02g-calculator.go) (line 10-51):

```go
var opMap = map[string]func(int, int) (int, error){
    "+": adder,
    "-": subtractor,
    "x": multiplier,
    "/": divider,
}

var opNameMap = map[string]string{
    "+": "adding",
    "-": "subtracting",
    "x": "multiplying",
    "/": "dividing",
}

func main() {
    (...)
    action, ok := opMap[op]
    if !ok {
        fmt.Println("Unknown operator:", op)
        os.Exit(1)
    }
    v, err := action(a, b)
    if err != nil {
        fmt.Println(opNameMap[op], "failed:", err)
    } else {
        fmt.Println(v)
    }
}

```

`opMap` associates the valid operators with the functions that implement their functionality. `opNameMap` associates the valid operators with the names of the operations. This is a good use of package-level state, because this information is immutable; The values in the maps never going to be changed.

While you can assign a function as the value for a map, it cannot be a key, because a function type (like a slice or a map) isn't comparable.

## <a name="methods-and-interfaces"></a>3. Methods and Interfaces

### <a name="declaring-your-own-type"></a>3.1. Declaring Your Own Types

Define an `Employee` for personnel management system

**ex.** [03a-people.go](./03-Methods-and-Interface/03a-people.go) (line 8-21):

```go
type Employee struct {
    ID        int
    FirstName string
    LastName  string
    DateHired time.Time
}

func main() {
    e1 := Employee{
        ID:        1,
        FirstName: "Bob",
        LastName:  "Bobson",
        DateHired: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
    }
    // Bob got married and changed his last name
    e1.LastName = "Bobson-Smith"
    fmt.Println(e1.LastName)
}
```

Use the user-defined type with the keyword `type`, followed by the name of the type (in example, `Employee`). Lists the fields after the `struct` keyword.

**ex.** [03b-people.go](./03-Methods-and-Interface/03b-people.go) (line 15-47):

```go
var Employees = map[int]Employee{}
var nextID = 0

func AddEmployee(firstName, lastName string, dateHired time.Time) int {
    nextID++
    Employees[nextID] = Employee{
        ID:        nextID,
        FirstName: firstName,
        LastName:  lastName,
        DateHired: dateHired,
    }
    return nextID
}

func GetEmployee(id int) (Employee, bool) {
    p, ok := Employees[id]
    return p, ok
}

func DMYToTime(day int, month time.Month, year int) time.Time {
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func main() {
    e1ID := AddEmployee("Bob", "Bobson", DMYToTime(10, time.January, 2020))
    e2ID := AddEmployee("Mary", "Maryson", DMYToTime(30, time.March, 2007))
    e1, exists1 := GetEmployee(e1ID)
    e2, exists2 := GetEmployee(e2ID)
    fmt.Println(e1, exists1)
    fmt.Println(e2, exists2)
    e3, exists3 := GetEmployee(2000)
    fmt.Println(e3, exists3)
}
```

Rather than access the `Employees` map directly, `AddEmployee`, and `GetEmployee` are added.

In `main`, `Employee` is not instantiated directly. Instead, call `AddEmployee` to create `Employee` instances that are stored in the `Employees` map and call `GetEmployee` to access them.

### <a name="adding-methods-to-types"></a>3.2. Adding Methods to Types

Go is not a pure object oriented programming language. Go does not provide **classes** but it does provide **structs**. **Methods** ca be added on structs.

**ex.** [03c-people.go](./03-Methods-and-Interface/03c-people.go) (line 23-45):

```go
func NewSimpleEmployeeData() *SimpleEmployeeData {
    return &SimpleEmployeeData{
        employees: map[int]Employee{},
        //more3
        nextID:    0,
    }
}

func (ed *SimpleEmployeeData) AddEmployee(firstName, lastName string, dateHired time.Time) int {
    ed.nextID++
    ed.employees[ed.nextID] = Employee{
        ID:        ed.nextID,
        FirstName: firstName,
        LastName:  lastName,
        DateHired: dateHired,
    }
    return ed.nextID
}

func (ed SimpleEmployeeData) GetEmployee(id int) (Employee, bool) {
    e, ok := ed.employees[id]
    return e, ok
}
```

Go does not support constructors, but `NewSimpleEmployeeData()` will act as a constructor.

The `AddEmployee` and `GetEmployee` methods are added to the `SimpleEmployeeData` struct.

The `GetEmployee` function declaration looks like:

```go
func GetEmployee(id int) (Employee, bool) {}
```

The `GetEmployee` method declaration looks like:

```go
func (ed SimpleEmployeeData) GetEmployee(id int) (Employee, bool) {}
```

The only difference is the **receiver** declaration between the `func` keyword and the method name. Within the body of the `GetEmployee` method, we access the fields on our `SimpleEmployeeData` instance using the `ed` receiver:

These methods are operates on the `SimpleEmployeeData` struct bundled together akin to a class.

### <a name="embedded-types"></a>3.2. Embedded Types

**ex.** [03d-people.go](./03-Methods-and-Interface/03d-people.go) (line 15-67, 78-93):

Add new type, `Manager` to `SimpleEmployeeData`.

```go
type Manager struct {
    Employee        // <- embedded field
    Reports []int
}


type SimpleEmployeeData struct {
    employees map[int]Employee
    managers  map[int]Manager
    nextID    int
}

func NewSimpleEmployeeData() *SimpleEmployeeData {
    return &SimpleEmployeeData{
        employees: map[int]Employee{},
        managers:  map[int]Manager{},
        nextID:    0,
    }
}

func (ed *SimpleEmployeeData) AddManager(firstName, lastName string, dateHired time.Time, reports []int) int {
    ed.nextID++
    ed.managers[ed.nextID] = Manager {
        Employee: Employee {
            ID:        ed.nextID,
            FirstName: firstName,
            LastName:  lastName,
            DateHired: dateHired,
        },
        Reports: reports,
    }
    return ed.nextID
}

func (ed SimpleEmployeeData) GetManager(id int) (Manager, bool) {
    m, ok := ed.managers[id]
    return m, ok
}

(...)

func manageEmployees(ed *SimpleEmployeeData) {
    e1ID := ed.AddEmployee("Bob", "Bobson", DMYToTime(10, time.January, 2020))
    e2ID := ed.AddEmployee("Mary", "Maryson", DMYToTime(30, time.March, 2007))
    e1, exists1 := ed.GetEmployee(e1ID)
    e2, exists2 := ed.GetEmployee(e2ID)
    fmt.Println(e1, exists1)
    fmt.Println(e2, exists2)
    e3, exists3 := ed.GetEmployee(2000)
    fmt.Println(e3, exists3)

    m1ID := ed.AddManager("Boss", "BossPerson", DMYToTime(17, time.June, 1982), []int{e1ID, e2ID})
    m1, exists4 := ed.GetManager(m1ID)
    fmt.Println(m1.FirstName, m1.Reports)

    // e4 := m1.Employee
    fmt.Println(m1, exists4)
}
```

The advantages of using embedded fields is accesing the fields from `Employee` within `Manager`, as though the fields were declared directly on `Manager`. That's why `fmt.Println(m1.FirstName, m1.Reports)` are working, even though `FirstName` is a field in `Employee`.

**Be aware that embedding is not inheritance.** You cannot assign a value of type `Manager` to a variable or field of type `Employee`.
