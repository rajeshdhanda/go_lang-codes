package main
import "fmt"
func main () {
	fmt.Println(true && true)   // && --->> and operator 
	fmt.Println(true && false)    
	fmt.Println(true || true)   // || --->> or operator 
	fmt.Println(true || false)
	fmt.Println(!true)   // ! --->> not operator
	fmt.Println(!false) 
}