package main

import "fmt"

func main() {
	ans := 2

	switch ans {                           // case syntax
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not a case")		
	}
}