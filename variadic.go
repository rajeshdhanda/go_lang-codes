package main
import "fmt"
func add(args...int) int {
	total := 0
	for _,value := range args{
		total += value
	}
	return total
}
func main(){
	result := add(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(result)
}