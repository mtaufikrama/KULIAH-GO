package main

import "fmt"

func main() {
	fmt.Print("Enter a Number : ")
	var input float32
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
