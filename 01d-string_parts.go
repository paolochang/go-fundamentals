package main

import "fmt"

func main() {
    baseString := "The quick brown fox jumped over the lazy dog"
    fmt.Println(baseString[0], baseString[len(baseString)-10])
    fmt.Println(baseString[16:26])
    
		sunString := "It's a 🌞 morning!" // this looks like it's length should be 17...
    fmt.Println(len(sunString))
    fmt.Println(sunString[7])
    fmt.Println(sunString[:8])
    fmt.Println(sunString[:11])

		stringBytes := []byte(sunString)
    fmt.Println(len(stringBytes), stringBytes)
    roundTrip := string(stringBytes)
    fmt.Println(roundTrip)
    b := sunString[0]
    fmt.Println(b, string(b))
    
    stringRunes := []rune(sunString)
    fmt.Println(len(stringRunes), stringRunes)
    roundTripRunes := string(stringRunes)
    fmt.Println(roundTripRunes)
    r := 'X'
    fmt.Println(r, string(r))
}
