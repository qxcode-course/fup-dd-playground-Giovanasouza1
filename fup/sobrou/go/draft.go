package main
import "fmt"
func main() {
    var produto1, produto2, produto3 float64
    fmt.Scan(&produto1, &produto2, &produto3)
    var valor1, valor2, valor3 float64
    fmt.Scan(&valor1, &valor2, &valor3)
    var dinheiro float64
    fmt.Scan(&dinheiro)

    gasto := produto1 * valor1 + produto2 * valor2 + produto3 * valor3
    sobrou := dinheiro - gasto
    fmt.Printf("%.2f\n", sobrou)
}
