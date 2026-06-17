package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)
    vetor := make([]int, qtd)

    for i := 0; i < qtd; i++{
        fmt.Scan(&vetor[i])
    }

    atual := 1
    maxfreq := 1
    for i := 0; i < qtd; i++{
        if vetor[i] == vetor[i + 1]{
            atual++
        }
        if atual > maxfre
        }
}
