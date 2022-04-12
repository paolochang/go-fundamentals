package main

import "fmt"

func main() {
    m := map[string]int{}
    m2 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    m3 := make(map[string]int)

    fmt.Println(m, m2, m3)
    m["hello"] = 20
    fmt.Println(m["hello"])
    fmt.Println(len(m), len(m2), len(m3))
    
		/** 
		 * Never use an uninitialized map variable.
		 * Unlike a slice, you cannot add a value to a nil map.
		 */
    // var badMap map[string]int
    // badMap["oops"] = 1
    // fmt.Println(badMap)
    
    m["present"] = 0
    fmt.Println(m["present"], m["not present"])
    
    val, ok := m["present"]
    fmt.Println(val, ok)

    val, ok = m["not present"]
    fmt.Println(val, ok)
		
		delete(m, "present")
    val, ok = m["present"]
    fmt.Println(val, ok)

    delete(m, "not present")
    val, ok = m["not present"]
    fmt.Println(val, ok)
}