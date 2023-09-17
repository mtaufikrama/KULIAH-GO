package main

import "fmt"

func main() {
	fmt.Print("enter a nummber : ")
	var input float32
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
