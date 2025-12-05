package main
import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	for _, v := range numbers {
		fmt.Println(v)
	}
}
