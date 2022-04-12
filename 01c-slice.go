package main

import "fmt"

func main() {
    var intSlice []int
    stringSlice := []string{"the", "quick", "brown", "fox"}
    zeroIntSlice := make([]int, 10)

    fmt.Println(intSlice, stringSlice, zeroIntSlice)
    fmt.Println(len(intSlice), len(stringSlice), len(zeroIntSlice))

    fmt.Println(stringSlice[0], stringSlice[len(stringSlice)-1])
    stringSlice[3] = "bear"
    fmt.Println(stringSlice[len(stringSlice)-1])

    intSlice = append(intSlice, 20)
    fmt.Println(intSlice)
    zeroIntSlice = append(zeroIntSlice, 100)
    fmt.Println(zeroIntSlice)
    fmt.Println(len(intSlice), len(zeroIntSlice))
    
    midSlice := stringSlice[1:3]
    fmt.Println(midSlice)
    startSlice := stringSlice[:2]
    fmt.Println(startSlice)
    endSlice := stringSlice[2:]
    fmt.Println(endSlice)

}
