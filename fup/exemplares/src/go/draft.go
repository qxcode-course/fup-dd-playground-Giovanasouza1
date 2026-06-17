package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)

    for i := 0; i < N; i ++ {
        fmt.Scan(&vetor[i])
    }

     for i := 0; i < N; i ++ {
        repetido := false
        for j := 0; j < i; j++ {
        if i != j && vetor[j] == vetor[i]  {
            repetido = true
            break
        }
        }
        if !repetido {
            fmt.Print(vetor[i], " ")
}
}
}
