package main

import "fmt"


func changeValue(str *string){
	*str = "changed!"
}

func main(){                       // & refers to memmory location where the value is stores 
	x := 7
	y := &x
    fmt.Println(x,y)
	*y = 8                         // changes the value where the pointer is pointing
	fmt.Println(x,y)   
	
	toChange := "hello"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)
}