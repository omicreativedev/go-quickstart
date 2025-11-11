package main

import (
	"fmt"
	"strings"
)

// 1. Create a simple Person object (using a struct)
type Person struct {
	// 2. Class attributes (instance variables)
	firstName string
	lastName  string
	age       int
	email     string
}

// 3. Constructor-like function for Person (Go idiom)
func NewPerson(firstName, lastName string, age int, email string) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		email:     email,
	}
}

// 4. Functions (methods) for Person object

// Value receiver method - doesn't modify the original
func (p Person) FullName() string {
	return p.firstName + " " + p.lastName
}

// Value receiver method
func (p Person) Introduce() string {
	return fmt.Sprintf("Hello, I'm %s, I'm %d years old. Contact me at %s", 
		p.FullName(), p.age, p.email)
}

// Pointer receiver method - can modify the object
func (p *Person) HaveBirthday() {
	p.age++
	fmt.Printf("Happy Birthday %s! You are now %d years old.\n", p.firstName, p.age)
}

// Pointer receiver method
func (p *Person) UpdateEmail(newEmail string) {
	oldEmail := p.email
	p.email = newEmail
	fmt.Printf("%s's email updated from %s to %s\n", p.firstName, oldEmail, newEmail)
}

// Pointer receiver method
func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

// Getter methods (value receivers)
func (p Person) GetAge() int {
	return p.age
}

func (p Person) GetEmail() string {
	return p.email
}

// 5. Implement the Stringer interface (similar to toString() in Java)
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}", 
		p.FullName(), p.age, p.email)
}

// 6. Create a Student object that "inherits" from Person using composition
type Student struct {
	Person  // Embedded struct - this is Go's way of achieving inheritance
	studentID string
	major     string
	gpa       float64
	courses   []string
}

// Constructor for Student
func NewStudent(firstName, lastName string, age int, email, studentID, major string) *Student {
	return &Student{
		Person: Person{
			firstName: firstName,
			lastName:  lastName,
			age:       age,
			email:     email,
		},
		studentID: studentID,
		major:     major,
		gpa:       0.0,
		courses:   []string{},
	}
}

// 7. Student-specific methods

// Method overriding - Student has its own version of Introduce
func (s Student) Introduce() string {
	return fmt.Sprintf("Hello, I'm %s, a %s major with ID %s. My GPA is %.2f", 
		s.FullName(), s.major, s.studentID, s.gpa)
}

// Student-specific methods
func (s *Student) AddCourse(course string) {
	s.courses = append(s.courses, course)
	fmt.Printf("Added course '%s' to %s's schedule\n", course, s.firstName)
}

func (s *Student) UpdateGPA(newGPA float64) {
	oldGPA := s.gpa
	s.gpa = newGPA
	fmt.Printf("%s's GPA updated from %.2f to %.2f\n", s.firstName, oldGPA, newGPA)
}

func (s Student) GetCourseList() string {
	if len(s.courses) == 0 {
		return "No courses enrolled"
	}
	return "Courses: " + strings.Join(s.courses, ", ")
}

func (s Student) StudentInfo() string {
	return fmt.Sprintf("Student ID: %s, Major: %s, GPA: %.2f\n%s", 
		s.studentID, s.major, s.gpa, s.GetCourseList())
}

// Override the String method for Student
func (s Student) String() string {
	return fmt.Sprintf("Student{Name: %s, ID: %s, Major: %s, GPA: %.2f}", 
		s.FullName(), s.studentID, s.major, s.gpa)
}

// 8. Interface for polymorphism
type Introducer interface {
	Introduce() string
}

func MakeIntroduction(i Introducer) {
	fmt.Println("=== Introduction ===")
	fmt.Println(i.Introduce())
	fmt.Println("====================")
}

// 9. Another embedded type to demonstrate multiple composition
type Athlete struct {
	sport string
	level string
}

type StudentAthlete struct {
	Student
	Athlete
	scholarship bool
}

func NewStudentAthlete(firstName, lastName string, age int, email, studentID, major, sport, level string, scholarship bool) *StudentAthlete {
	return &StudentAthlete{
		Student: Student{
			Person: Person{
				firstName: firstName,
				lastName:  lastName,
				age:       age,
				email:     email,
			},
			studentID: studentID,
			major:     major,
		},
		Athlete: Athlete{
			sport: sport,
			level: level,
		},
		scholarship: scholarship,
	}
}

func (sa StudentAthlete) Introduce() string {
	baseIntro := sa.Student.Introduce()
	return fmt.Sprintf("%s. I also play %s at %s level.", baseIntro, sa.sport, sa.level)
}

func main() {
	fmt.Println("=== OBJECTS AND INHERITANCE IN GO DEMONSTRATION ===\n")

	// 10. Test Person object
	fmt.Println("1. TESTING PERSON OBJECT")
	fmt.Println("------------------------")
	
	// Create a Person object
	person := NewPerson("John", "Doe", 30, "john.doe@email.com")
	fmt.Println("Created person:", person)
	
	// Call Person methods
	fmt.Println("\nCalling Person methods:")
	fmt.Println("FullName:", person.FullName())
	fmt.Println("Introduce:", person.Introduce())
	fmt.Println("Age:", person.GetAge())
	fmt.Println("Email:", person.GetEmail())
	
	// Modify Person variables
	fmt.Println("\nModifying Person variables:")
	person.HaveBirthday()                    // Method that modifies age
	person.UpdateEmail("john.new@email.com") // Method that modifies email
	person.SetFirstName("Jonathan")          // Method that modifies first name
	
	fmt.Println("After modifications:")
	fmt.Println(person.Introduce())
	fmt.Println("String representation:", person)
	
	// 11. Test Student object
	fmt.Println("\n\n2. TESTING STUDENT OBJECT (INHERITANCE)")
	fmt.Println("----------------------------------------")
	
	// Create a Student object
	student := NewStudent("Jane", "Smith", 20, "jane.smith@university.edu", "S12345", "Computer Science")
	fmt.Println("Created student:", student)
	
	// Call inherited methods from Person
	fmt.Println("\nCalling inherited methods from Person:")
	fmt.Println("FullName (inherited):", student.FullName()) // Directly inherited
	fmt.Println("Age (inherited):", student.GetAge())        // Directly inherited
	
	// Call overridden methods
	fmt.Println("\nCalling overridden methods:")
	fmt.Println("Introduce (overridden):", student.Introduce())
	
	// Call Student-specific methods
	fmt.Println("\nCalling Student-specific methods:")
	student.AddCourse("Data Structures")
	student.AddCourse("Algorithms")
	student.AddCourse("Database Systems")
	student.UpdateGPA(3.75)
	
	fmt.Println(student.StudentInfo())
	
	// Modify inherited variables through Student
	fmt.Println("\nModifying inherited variables through Student:")
	student.HaveBirthday() // Can call Person methods on Student
	student.UpdateEmail("jane.new@university.edu")
	
	fmt.Println("After modifications:")
	fmt.Println(student.Introduce())
	fmt.Println("String representation:", student)
	
	// 12. Demonstrate polymorphism with interface
	fmt.Println("\n\n3. TESTING POLYMORPHISM WITH INTERFACE")
	fmt.Println("--------------------------------------")
	
	// Both Person and Student implement Introducer interface
	var introducer1 Introducer = person
	var introducer2 Introducer = student
	
	MakeIntroduction(introducer1)
	MakeIntroduction(introducer2)
	
	// 13. Test multiple composition with StudentAthlete
	fmt.Println("\n\n4. TESTING MULTIPLE COMPOSITION")
	fmt.Println("-------------------------------")
	
	studentAthlete := NewStudentAthlete(
		"Mike", "Johnson", 21, "mike.johnson@university.edu", 
		"S67890", "Biology", "Basketball", "Varsity", true,
	)
	
	// Add some courses
	studentAthlete.AddCourse("Genetics")
	studentAthlete.AddCourse("Biochemistry")
	studentAthlete.UpdateGPA(3.9)
	
	fmt.Println("Student Athlete:", studentAthlete)
	fmt.Println("\nIntroduction:")
	fmt.Println(studentAthlete.Introduce())
	fmt.Println("\nAthlete info:", studentAthlete.sport, "-", studentAthlete.level)
	fmt.Println("Scholarship:", studentAthlete.scholarship)
	
	// 14. Demonstrate accessing embedded struct methods directly
	fmt.Println("\n\n5. ACCESSING EMBEDDED STRUCT METHODS")
	fmt.Println("------------------------------------")
	
	// We can access the embedded Person's methods directly
	fmt.Println("Direct Person method:", studentAthlete.FullName())
	
	// Or through the embedded field
	fmt.Println("Through embedded field:", studentAthlete.Student.FullName())
	fmt.Println("Through double embedding:", studentAthlete.Student.Person.FullName())
	
	// 15. Show that all types satisfy the Stringer interface
	fmt.Println("\n\n6. STRINGER INTERFACE DEMONSTRATION")
	fmt.Println("-----------------------------------")
	
	// All our types automatically work with fmt.Println because they implement String()
	objects := []fmt.Stringer{person, student, studentAthlete}
	
	for i, obj := range objects {
		fmt.Printf("Object %d: %s\n", i+1, obj)
	}
	
	fmt.Println("\n=== DEMONSTRATION COMPLETE ===")
}