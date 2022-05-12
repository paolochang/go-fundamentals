package main

import (
	"fmt"
	"time"
)

type Employee struct {
    ID        int
    FirstName string
    LastName  string
    DateHired time.Time
}

//more1

type SimpleEmployeeData struct {
    employees map[int]Employee
    //more2
    nextID    int
}

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

//more4

func DMYToTime(day int, month time.Month, year int) time.Time {
    return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func main() {
    ed := NewSimpleEmployeeData()
    manageEmployees(ed)
}

func manageEmployees(ed *SimpleEmployeeData) {
    e1ID := ed.AddEmployee("Bob", "Bobson", DMYToTime(10, time.January, 2020))
    e2ID := ed.AddEmployee("Mary", "Maryson", DMYToTime(30, time.March, 2007))
    e1, exists1 := ed.GetEmployee(e1ID)
    e2, exists2 := ed.GetEmployee(e2ID)
    fmt.Println(e1, exists1)
    fmt.Println(e2, exists2)
    e3, exists3 := ed.GetEmployee(2000)
    fmt.Println(e3, exists3)
    //more5
}

//more6
