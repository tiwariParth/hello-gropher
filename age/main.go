package main

import "fmt"

func main() {
    var age = 10
    fmt.Println("User's age:", age)

    if age <= 12 {
        fmt.Println("User is a child.")
    } else if age >= 13 && age <= 19 {
        fmt.Println("User is a teenager.")
    } else if age >= 20 && age <= 64 { 
        fmt.Println("User is an adult.")
    } else {
        fmt.Println("User is a senior.")
    }
}


