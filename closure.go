package main
import "fmt"
func main() {
	add := func(a,b int) int {
		return a+b
	}
	fmt.Println(add(3,5))
	fmt.Println(add(5,6))
}