package main
import "fmt"
func main() {    
	panic("PANIC") // cause a run time error   
	str := recover() // this will never happen call to panic immediately stops execution of the function
	fmt.Println(str) 
} 