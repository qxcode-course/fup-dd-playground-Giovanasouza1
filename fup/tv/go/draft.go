package main
import "fmt"
func main() {
	var valor float64
	var parcelas int
	fmt.Scan(&valor, &parcelas)
	valor *= (1 + 0.05 + (parcelas - 1))
	cada := valor / float64(parcelas)
	fmt.Printf()

}