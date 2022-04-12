# Go Fundamentals

## Using Types and Declarations in Go

### Variable Declarations & Type Conversions

#### Declarations and Print Output

```go
func main() {
    var x int = 10 // full variable declaration format
    var y = 20     // type isn't needed, since the default type for the numeric literal 20 is int
    var z int      // type is specified, but value isn't, initializing z to the zero value for int (0)

    fmt.Println(x, y, z)
}
```

Use `Println` function in the `fmt` package to output the values of our variables. If we don't read a variable value at least once, the Go compiler refuses to compile our code. Let's see that by commenting out the `import` statement and the `fmt.Println` statement.

Unused variables and packages will return the error message.

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
