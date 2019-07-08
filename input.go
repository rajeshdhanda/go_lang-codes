package main
import "fmt"
var (    a = 5 
		 b = 10    
		 c = 15 
)
func main() {
	fmt.Println("Enter a number: ")    
	var input float64    
	fmt.Scanf("%f", &input)
	fmt.Print("double of your number is : ",2*input) 
}