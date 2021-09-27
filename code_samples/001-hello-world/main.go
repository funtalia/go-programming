package main

import "fmt"

var y = 43; // we can declare a variable out of function body by var. scope of y is package level 
var z  string = "hi"
func main() {
	n, _ := fmt.Println("Hello New World", 42, true) // ...interface{}  unlimited inputs available
	fmt.Println(n)

	// short declaration operator
	x := 42; //declare and assign
	fmt.Println(x)

	
	fmt.Println(y)
	
	fmt.Printf("%T\n", z)
	// z =  5 ;  cannot use 5 (type untyped int) as type string in assignment
	// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
	// not a DYNAMIC programming language

}
