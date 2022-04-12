# Go Fundamentals

## Index

[Variable Declarations & Type Conversions](###Variable-Declarations-&-Type-Conversions)
[Strings and Slices](###Strings-and-Slices)

## Using Types and Declarations in Go

### Variable Declarations & Type Conversions

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

### Strings and Slices

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

```
[quick brown]
[the quick]
[brown bear]
```

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
