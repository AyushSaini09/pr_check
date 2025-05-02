package main

import "fmt"

// Define a struct (similar to a class)
type Person struct {
    name string
    age  int
}

// A function that takes a Person and returns a greeting message
func greet(p Person) string {
    return fmt.Sprintf("Hello, %s! You are %d years old.", p.name, p.age)
}

func main() {
    // Declare and initialize variables
    var x int = 10
    y := 20  // Short-hand syntax

    fmt.Println("Sum:", x+y)

    // If-else conditional
    if x > y {
        fmt.Println("x is greater than y")
    } else {
        fmt.Println("x is less than or equal to y")
    }

    // For loop
    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }

    // Create and use a struct
    person := Person{name: "Alice", age: 30}
    message := greet(person)
    fmt.Println(message)
}
