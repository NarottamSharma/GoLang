package main

import (
	"fmt"
)

func main(){
	views1 := 10000
	views2 := 20000
	totalViews := views1+views2
	fmt.Println(totalViews)

	likes:=10
	fmt.Println(likes)
	likes++
	fmt.Println("10, After Increment",likes)

	avgViews := totalViews/2
	fmt.Println(avgViews)
	fmt.Println(totalViews, likes)


	rating1 := 4.5
	rating2 := 5.1
	avgRating := (rating1+rating2)/2
	fmt.Println(avgRating)

}