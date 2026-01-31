package main

import "fmt"

func main(){
	isLogged:= true
	isAdmin:= true
	hasSubscription:= false
	//  AND 	&&
	canOpenDashboard := isLogged && hasSubscription
	fmt.Println(canOpenDashboard)

	canDeletePost := isAdmin || (isLogged && hasSubscription)
	fmt.Println(canDeletePost)

	age:= 25
	isAdult := age>18
	fmt.Println(isAdult)


}