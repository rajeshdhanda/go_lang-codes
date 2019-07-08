package main
import "fmt"
func zero(xPtr *int) {      
	*xPtr = 0  /* When we write *xPtr = 0, we are saying --->>  
	store the int 0 in the memory location xPtr refers to */

}  
func main() {      
	x := 5      
	zero(&x)      
	fmt.Println(x) // x is 0  
}   