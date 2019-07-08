package main
import "fmt"
func main(){
	var x[4]int
	for i :=0; i<4; i++ {
		x[i] = 3.0*i
	}
	fmt.Println(x)
}