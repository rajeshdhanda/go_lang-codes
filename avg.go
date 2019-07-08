package main
import "fmt"

func main() {
	ys := []float64{23,45,67,92,53,95,10,4,78,57,28,42}
	fmt.Println(average(ys))
}
func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {   // _ for index in array, key in map
		total += value
	}
	return total / float64(len(xs))
}