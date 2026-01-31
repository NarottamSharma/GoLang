package main
import "fmt"

func main(){
	score:=12
	if score>50 && score <=70{
		fmt.Println("B")
	} else if score>70{
		fmt.Println("A")
	} else {
		fmt.Println("fail")
	}
}