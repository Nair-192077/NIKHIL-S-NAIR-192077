package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv" 
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)    //user input
	fmt.Printf("Type the year you were born : ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(),10,64)   // base is 10 and size is 64 of the number we want to store
	fmt.Printf("You will be %d years old", 2020-input)

}