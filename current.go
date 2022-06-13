package main

import "fmt"

func main() {
	var i = 1
	var j = 1
	upto := 0
	var k int
	fmt.Println("Enter the digit upto Print:")
	fmt.Scanln(&upto)
	fmt.Println(i)
	fmt.Println(j)

	for a := 3; a <= upto; a++ {
		k = i + j

		fmt.Println(k)

		i = j
		j = k

	}
}
