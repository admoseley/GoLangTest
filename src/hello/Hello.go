package main

import (
	"fmt"

	"github.com/admoseley/moseleyutils"
	"github.com/admoseley/stringutil"
)

/*
func add(x int, y int) int {
	return x + y
}
*/
func main() {
	var MyFirstName string = "Adrian"
	var MyLastName string = "Moseley"
	var MyAge int = 41
	var MyYear int = 1976
	fmt.Printf("First Name: %s\nLast Name: %s\nAge:%d\nBirth Year: %d\n", MyFirstName, MyLastName, MyAge, MyYear)
	fmt.Printf("Hello, world\n")
	fmt.Printf("Hello, Adrian!\n")
	fmt.Println(moseleyutils.Moseley_add(3, 4))
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	var MoseleyADD int = moseleyutils.Moseley_add(10, 11)

	fmt.Printf("Adding the Moseley Func of 10 + 20 is %d", MoseleyADD)
	/*
		go tool tour
	*/

	/** **/

}
