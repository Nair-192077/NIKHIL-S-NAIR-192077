package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":5,
		"pear":6,
	}
	//mp := make(map[string]int)   // another way to write map
	mp["apple"] = 100              // change value of any item
	delete(mp,"pear")             // deletion of any key
	fmt.Println(mp)
	fmt.Println(len(mp))            // prints length of mp
}