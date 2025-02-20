package main

import "fmt"

func main(){
	var name string
	var age int
	fmt.Println("Whats your name")
	fmt.Scanf("%s", &name)

	fmt.Println("Whats your age")
	fmt.Scanf("%d", &age)
	fmt.Printf("Hello %s,\nYou are %d years old.\n", name, age)
}