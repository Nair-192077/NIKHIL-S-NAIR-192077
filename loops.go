//if else , else if 
package main

import (
	"fmt"
)

func main() {
	age := 10
	
	if age >=18 {                            // if statement
		fmt.Println("You can drive alone")

	}else if age>=14 {                          // else if statement
		fmt.Println("Can drive with a parent")
	} else{                                          // else statement
		fmt.Println("You cannot drive")
	}
    
}