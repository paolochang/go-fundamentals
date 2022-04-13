package main

import "fmt"

func main() {
    var x int = 10 // full variable declaration format
    var y = 20     // type isn't needed, since the default type for the numeric literal 20 is int
    var z int      // type is specified, but value isn't, initializing z to the zero value for int (0)

    fmt.Println(x, y, z)
    
    x = y
    y = 50
    z = 30 * x

    fmt.Println(x, y, z)

    var b byte = 30 // must convert the byte type with type conversion: int(b)
    x = int(b)
    y = y + int(b) 

    fmt.Println(x, y, z)

    f := 3.2
    i := 10

    fmt.Println(f, i)

    i, j := 30, 40
    
    fmt.Println(f, i, j)
}
