package main
import "fmt"
func main() {
	slice1:= []int{1,2,3,4}
	slice2 := append(slice1,5,6)
	fmt.Println("slice1 : ", slice1,"\n","slice2 : ",slice2) 
}
