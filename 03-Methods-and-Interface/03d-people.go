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

type Manager struct {
    Employee
    Reports []int
}


type SimpleEmployeeData struct {
    employees map[int]Employee
    managers  map[int]Manager
    nextID    int
}

func NewSimpleEmployeeData() *SimpleEmployeeData {
    return &SimpleEmployeeData{
        employees: map[int]Employee{},
        managers:  map[int]Manager{},
        nextID:    0,
    }
}

func (ed *SimpleEmployeeData) AddEmployee(firstName, lastName string, dateHired time.Time) int {
    ed.nextID++
    ed.employees[ed.nextID] = Employee {
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

func (ed *SimpleEmployeeData) AddManager(firstName, lastName string, dateHired time.Time, reports []int) int {
    ed.nextID++
    ed.managers[ed.nextID] = Manager {
        Employee: Employee {
            ID:        ed.nextID,
            FirstName: firstName,
            LastName:  lastName,
            DateHired: dateHired,
        },
        Reports: reports,
    }
    return ed.nextID
}

func (ed SimpleEmployeeData) GetManager(id int) (Manager, bool) {
    m, ok := ed.managers[id]
    return m, ok
}

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
    
	m1ID := ed.AddManager("Boss", "BossPerson", DMYToTime(17, time.June, 1982), []int{e1ID, e2ID})
    m1, exists4 := ed.GetManager(m1ID)
    fmt.Println(m1.FirstName, m1.Reports)

    //* cannot use m1 (variable of type Manager) as type Employee in variable declaration
    // var e4 Employee = m1;
    // fmt.Println(e4.LastName)
    
	e4 := m1.Employee
    fmt.Println(m1, exists4)
    fmt.Println(e4)
}
