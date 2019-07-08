package main
import "fmt"
var x1 = "hi, i am global variable"
func main() {
    var x string = "hi, i am local variable"
    fmt.Println(x)
}
func f() {
	fmt.Println(x1)
}   