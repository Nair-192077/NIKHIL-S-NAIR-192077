package main

import "fmt"

func add(x int, y int) (int, int) { // functions can accept parameters (parametric function)
	return x + y, x - y
	//fmt.Println(x + y)
}

func sub(x, y int) (z1, z2 int) {
	defer fmt.Println("Nick")
	z1 = x - y
	z2 = y - x
	return
}

func main() { //function means a block of reusable code

	ans1, ans2 := add(5, 10) // function call
	fmt.Println(ans1, ans2)
	fmt.Println(sub(14, 7))

}
