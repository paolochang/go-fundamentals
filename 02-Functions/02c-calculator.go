package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

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

    c, err := adder(a, b)
    if err != nil {
        fmt.Println("adding failed:", err)
    } else {
        fmt.Println(c)
    }

    d, err := subtractor(a, b)
    if err != nil {
        fmt.Println("subtracting failed:", err)
    } else {
        fmt.Println(d)
    }

    e, err := multiplier(a, b)
    if err != nil {
        fmt.Println("multiplying failed:", err)
    } else {
        fmt.Println(e)
    }
		
    f, err := divider(a, b)
    if err != nil {
        fmt.Println("dividing failed:", err)
    } else {
        fmt.Println(f)
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