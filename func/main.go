package main

import "fmt"

func AreaOfRectangle (length,width int) int {
	return length*width
}

func main(){
	var l = 12
	var w = 21

	fmt.Println("Area of Rectangle is: ", AreaOfRectangle(l,w))
}