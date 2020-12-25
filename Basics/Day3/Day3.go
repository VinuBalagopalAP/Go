package main

import "fmt"

func main() {

	var x int64
	var y int64
	fmt.Print("Enter first digit\n")
	fmt.Scan(&x)
	fmt.Print("Enter second digit\n")
	fmt.Scan(&y)

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y && x == y)
	fmt.Println(x == y || x < y)
	fmt.Println(!(x <= y))
}
