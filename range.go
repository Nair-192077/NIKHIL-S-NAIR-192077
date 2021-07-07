package main

import "fmt"

func main()  {
	var a []int = []int{1,2,4,5,7,8,10,15,20}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, element := range a{                  //use of range
		fmt.Printf("%d : %d\n",i, element)
	}
}