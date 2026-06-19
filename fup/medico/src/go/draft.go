package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }

    cont := 0
    for i := 0; i < N; i++ {
        if i == 0 && vetor[i] == 0 && vetor[i + 1] == 0 {
            cont++
        } else if i == N-1 && vetor[i - 1] == 0 && vetor[i] == 0 {
            cont ++
        } else if i > 0 && i < N-1 && vetor[i] == 0 && vetor[i-1] == 0 && vetor[i+1] == 0 {
            cont++
        }
}
fmt.Println(cont)
}