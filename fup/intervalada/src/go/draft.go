package main
import "fmt"
func main() {
    var N, primeiro, ultimo int
    fmt.Scan(&N, &primeiro, &ultimo)

    vetor := make([]int, N)
    
    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }

    cont := 0
    for  i := 0; i < N; i++ {
        if vetor[i] > primeiro && vetor[i] < ultimo {
            cont++
        } else if vetor[i] == primeiro || vetor[i] == ultimo {
            cont++
}
    }
  fmt.Println(cont)
}