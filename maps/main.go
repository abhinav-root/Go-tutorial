package main

import "fmt"

func main() {
	myMap := map[string]int{}
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
