package main

// import brings external packages into file that we are working 
// where actually we need it
import (
	"fmt"
	"math"
)

func main(){
	// packagename.functionname	-> call a function from package
	fmt.Println("sqrt of 25", math.Sqrt(25))
	
}