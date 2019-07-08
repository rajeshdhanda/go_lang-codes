package main
import "fmt"
func first() {    
	fmt.Println("1st") 
}
func second() {    
	fmt.Println("2nd") 
}
// defer schedules a function call to be run after the function completes. 
func main() {
	defer second()    
	first() 
} 

// program prints 1st followed by 2nd. Basically, defer moves the call to second to the end of the function:
