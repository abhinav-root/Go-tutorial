package main

import "fmt"

func main() {
	fmt.Println("Hello go")
	var age int = 18
	// we can omit the type
	var age2 = 23
	// we can also omit the var syntax
	age3 := 32
	var weight float32 = 74.23
	var myname string = "Abhinav Singh"
	var isMarried bool = false
	isMarried = true

	fmt.Println(age, age2, age3, weight, myname, isMarried)

	// constants never change
	// we cannot use short syntax when declaring constants.
	const moneyInRupees int = 73000
	fmt.Println(moneyInRupees)
}
