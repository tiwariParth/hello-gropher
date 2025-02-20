package main

import "fmt"

func realNumbers(prompt string) int {
 var num int
 for {
	fmt.Print(prompt)
	_, err := fmt.Scanf("%d", &num)
	if err == nil {
	    break
	}
	fmt.Println("Invalid input. Please enter a valid number.")
 }
 return num
}

func main(){
	num1 := realNumbers("Enter Number1")
	num2 := realNumbers("Enter Number2")
	fmt.Println(num1,num2)

	var opt string 
	_,err:=fmt.Scanf("%s",&opt)
	if err != nil{
		fmt.Println("Lol",err)
	}
	switch opt{
	case "+" :
		fmt.Println( num1 + num2)
		
	case "-" :
		fmt.Println( num1 - num2)
		
	case "*" :
		fmt.Println( num1 * num2)
		
	case "/" :
		if (num2 == 0){
			fmt.Println("Error: Division by zero is not allowed.")
		}
		fmt.Println( num1 / num2)
		
	case "%" :
		if (num2 == 0){
			fmt.Println("Error: Division by zero is not allowed.")
		}
		fmt.Println( num1 % num2)
	default:
		fmt.Println("invalid opt")
	}
}
