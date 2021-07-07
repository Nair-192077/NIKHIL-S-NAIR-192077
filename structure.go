package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int32
}

type Point struct {
	x int32
	y int32
}

func main() {
	var p1 Point = Point{1, 2}
	fmt.Println(p1)

}
