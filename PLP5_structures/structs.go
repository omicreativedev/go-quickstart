package main

import "fmt"

// Person object with attributes
type Person struct {
    name string
    age  int
}

// Person functions
func (p Person) Introduce() string {
    return "Hi, I'm " + p.name
}

func (p *Person) HaveBirthday() {
    p.age++
    fmt.Println(p.name, "is now", p.age, "years old!")
}

// Student inherits from Person
type Student struct {
    Person       // This means Student gets all Person stuff
    studentID string
    grade     string
}

// Student's own function
func (s Student) Study() string {
    return s.name + " is studying hard!"
}

func main() {
    // Make a Person
    person := Person{name: "John", age: 25}
    
    // Make a Student (this inherits from Person)
    student := Student{
        Person:   Person{name: "Sarah", age: 20},
        studentID: "S123",
        grade:     "A",
    }
    
    // Test Person
    fmt.Println("PERSON")
    fmt.Println(person.Introduce())
    person.HaveBirthday()
    
    // Test Student
    fmt.Println("\nSTUDENT")
    fmt.Println(student.Introduce())  // Uses inherited method
    fmt.Println(student.Study())      // Its own method
    fmt.Println("ID:", student.studentID)
    fmt.Println("Grade:", student.grade)
    
    // Modify student's age (inherited from Person)
    student.HaveBirthday()
}
