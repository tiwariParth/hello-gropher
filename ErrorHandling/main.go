package main

import "fmt"

func main(){
	var num1 int 
	_,err:=fmt.Scanf("%d",&num1)
	if err != nil{
		fmt.Println("Error", err)
	}

	var num2 int 
	_,err=fmt.Scanf("%d",&num2)
	if err != nil{
		fmt.Println("Error", err)
	}

	var sum int = num1+num2
	fmt.Println("Sum :",sum)


}