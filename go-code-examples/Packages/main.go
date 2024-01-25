package main

import (
	"fmt"
	// Importing the package we created
	yourpackage "pckgs/yourpackage" 
	/* Name "yourpackage" can be here but not need to be.
	Overall can improve readability */
)

func main() {
	fmt.Println(yourpackage.Hello())
}
