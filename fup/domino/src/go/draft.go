package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    pecas := ((N + 1) * (N + 2))/ 2

    fmt.Println(pecas)
}