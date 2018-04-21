package main

// Packages can be imported from other areas
// multiple packages are wrapped in curlies
import (
	"fmt"
	"os"
)

// To compile use cli command 'go build .' from the directory the prog is saved to
// GOOS=linux go build . --- would build a binary for linux, easy to build platform specific precompiled progs
// Go is JIT but can be AOT via build
// when compiling the static build gives you the option to define the compiled binary name using the '-o' switch

// Simplicity is Complicated video on YouTube for info on Go language creation and core
//Building docker containers is extermely simple with go

// Implicit type conversion
// Standard library contains almost every component necessary for programming
// tour.golang.org is an excellent resource ***** go playground is decent for such things as well
// golang.org/doc/effective_go.html is a good coverage of best-prax and coverage

/*
$GOPATH - location where your go source lives. Conforms strictly to directory structure:
`bin
`package
`src <-- program source saves here
*/

// Main section
func main() {

	//variabls are set declaritively
	var inputArg string = ""

	if len(os.Args) == 0 {
		inputArg = os.Args[1]
	} else {

		inputArg = "Nothing was entered!"
	}
	fmt.Printf("Cheetos %s\n", inputArg)
}
