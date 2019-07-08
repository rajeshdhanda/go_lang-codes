package main
import "fmt"

func main() {
	t := 1
	fmt.Println("initial value of t : ",t)
	for t < 7 {
		t+=1
		fmt.Println("value of t after increase of 1 is : ",t)
	}
	for m := 0; m<=5; m++{
			fmt.Println("m :",m)	
	}
}
