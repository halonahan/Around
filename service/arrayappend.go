package main
import "fmt"
func main() {
	var s []int
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1,2,3,4)
	printSlice(s)
}

func printSlice(s []int) {

	fmt.Printf("Array is %v\n", s)
}
