# Go Fundamentals

## Index

1. [Using Types and Declarations in Go](#declarations)
   1. [Variable Declarations & Type Conversions](#variable-declarations)
   2. [Strings and Slices](#strings-and-slices)
   3. [Map](#map)
2. [Calling Functions](#calling-functions)
   1. [Declaring and Calling Function](#declaring-and-calling-function)

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

`Output`:

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

ex. [01e-maps.go](./01-Declarations/01e-maps.go) (line 6-12):

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

ex. [01e-maps.go](./01-Declarations/01e-maps.go) (line 27-34):

```go
    m["present"] = 0
    fmt.Println(m["present"], m["not present"])

    val, ok := m["present"]
    fmt.Println(val, ok)

    val, ok = m["not present"]
    fmt.Println(val, ok)
```

`Output`:

    0 true
    0 false

#### Deleting From Maps

ex. [01e-maps.go](./01-Declarations/01e-maps.go) (line 36-38):

```go
    delete(m, "present")
    val, ok = m["present"]
    fmt.Println(val, ok)
```

Use the delete function to remove a key-value pair from a map. The first parameter passed to the delete function is the map, and the second parameter is the key to remove. This function returns nothing. If a key that's not in the map, is passed, the function call does nothing.

## <a name="calling-functions"></a>2. Calling Functions

### <a name="declaring-and-calling-function"></a>2.1 Declaring and Calling Function

#### Example of Calling Functions

ex. [02a-calculator.go](./02-Functions/02a-calculator.go)

```go
package main

import (
    "fmt"
)

func main() {
    a := 20
    b := 10
    c := adder(a, b)
    fmt.Println(c)
    d := subtractor(a, b)
    fmt.Println(d)
    e := multiplier(a, b)
    fmt.Println(e)
    f := divider(a, b)
    fmt.Println(f)
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

#### Reading the Command Line

Go allows to accept arguments as other program languages.

```
$ go run <program_name> <arguments>
```

- `os.Args` from `os` packages allows to use the arguments
- `strconv.Atoi` from `strconv` package convert `string` to `int`
