package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Two integer parameters and an operator expected")
        os.Exit(1)
    }

    a, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("invalid first argument:", err)
        os.Exit(1)
    }

    op := os.Args[2]
    
    b, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("invalid second argument:", err)
        os.Exit(1)
    }

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

    v, err := action(a, b)
    if err != nil {
        fmt.Println(actionName, "failed:", err)
    } else {
        fmt.Println(v)
    }
}


func adder(a, b int) (int, error) {
    x := a + b
    if (x > a) != (b > 0) {
        return x, errors.New("addition out of bounds")
    }
    return x, nil
}

func subtractor(a, b int) (int, error) {
    x := a - b
    if (x < a) != (b > 0) {
        return x, errors.New("subtraction out of bounds")
    }
    return x, nil
}

func multiplier(a, b int) (int, error) {
    if a == 0 || b == 0 {
        return 0, nil
    }
    x := a * b
    if x/b != a {
        return x, errors.New("multiplication out of bounds")
    }
    return x, nil
}

func divider(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}