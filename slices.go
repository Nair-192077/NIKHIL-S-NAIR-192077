package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3] //slicing array x and storing to new array
	fmt.Println(s)
	fmt.Println(cap(s))
	/*b = append(x , 10)                   // append add no to the array specified
	fmt.Println(x)*/
}
