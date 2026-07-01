package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }

    resultado := 0
     for i := 0; i < N; i++ {
        resultado = resultado*10 + vetor [i]
}
fmt.Println(resultado)
}