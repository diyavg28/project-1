package main
import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p Product) Display() {
	fmt.Println(p.Name, p.Price)
}

func main() {
	p := Product{"Laptop", 60000}
	p.Display()
}
