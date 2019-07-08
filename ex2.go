package main
import "fmt"
func main() {    
	fmt.Println(len("Hello,World"))    
	fmt.Println("Hello, World"[1]) /* returns a number.
	  This is because the character is represented by a byte (remember a byte is an integer). */    	
	fmt.Println("Hello, " + "World")
} 