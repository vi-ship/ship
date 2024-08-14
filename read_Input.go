package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Scan(&name)
	fmt.Scan(&age)

	fmt.Println(name, "is", age, "years old.")
}
