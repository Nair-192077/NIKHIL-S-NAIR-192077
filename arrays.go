package main

import "fmt"

func main() {
    arr := [5]int{0,1,2,3,4}            // array syntax
	sum :=0

	arr2d := [2][2]int{{1,2},{3,4}}        // 2d array
	fmt.Println(arr2d)

	for i := 0; i < len(arr); i++{
		sum += arr[i]
	}	
	fmt.Println(sum)	
	}