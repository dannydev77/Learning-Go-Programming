package main

import (
	"fmt"
)

// method 4
// Declaring at the package level
// you have to use the full declaration syntax
// var a int = 42

// method 5
// using a block of variables that are declared together
// var (
// 	studentName    string = "john doe"
// 	subject        string = "Informational sciences"
// 	years          int    = 4
// 	totalSemesters int    = 8
// )

var a int = 44

func main() {

	// method 3
	// a := 55

	var a int = 98
	fmt.Println("===================================")
	fmt.Println(a)
	fmt.Println("===================================")
}
