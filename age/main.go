package main

import "fmt"

func main(){
	var age = 10
	fmt.Println("User's age", age)

	if (age < 18 && age <= 10) {
		fmt.Println("User is child ")
	} else if (age > 12 && age < 18){
		fmt.Println("User is teenager")
	} else if (age > 18 && age < 50 ) {
		fmt.Println("User is adult")
	} else {
		fmt.Println("User is senior")
	}
}