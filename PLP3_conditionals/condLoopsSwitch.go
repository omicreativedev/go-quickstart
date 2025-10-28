package main

import "fmt"

func main() {
    // If/else statement
    x := true
    if x == true {
        fmt.Println("x is true")
    } else {
        fmt.Println("x is false")
    }

    // If with multiple conditions
    if x == true && 5 > 3 {
        fmt.Println("Both conditions met")
    }

    // Else if (emulated)
    score := 85
    if score >= 90 {
        fmt.Println("A")
    } else if score >= 80 {
        fmt.Println("B")
    } else {
        fmt.Println("C")
    }

    // Basic for loop
    for i := 0; i < 3; i++ {
        fmt.Println("Basic loop:", i)
    }

    // While loop (emulated with for)
    y := 2
    for y > 0 {
        fmt.Println("While loop:", y)
        y--
    }

    // For each with range
    numbers := []int{10, 20, 30}
    for index, value := range numbers {
        fmt.Printf("Index %d: %d\n", index, value)
    }

    // Infinite loop with break
    count := 0
    for {
        if count >= 2 {
            break
        }
        fmt.Println("Infinite loop")
        count++
    }

    // Continue statement
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println("Odd number:", i)
    }

    // Break with label
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break OuterLoop
            }
            fmt.Printf("Labeled break: i=%d, j=%d\n", i, j)
        }
    }

    // Switch statement
    day := "Friday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("Almost weekend")
    case "Saturday", "Sunday":
        fmt.Println("Weekend!")
    default:
        fmt.Println("Midweek")
    }

    // Switch with fallthrough
    number := 2
    switch number {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Other number")
    }

    // Short circuit evaluation
    if false && someFunction() {
        fmt.Println("This won't execute")
    }
    if true || someFunction() {
        fmt.Println("Short circuit prevents function call")
    }
}

func someFunction() bool {
    fmt.Println("This function was called")
    return true
}
