package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    
    vetor := make([]int, N)
    

    for i := 0; i < N; i++{
        fmt.Scan(&vetor[i])
    }

    fmt.Print("[")
    for i := 0; i < N; i++ {
    if vetor[i] == 1 {
        fmt.Print("A")
    } else if vetor[i] == 11 {
        fmt.Print("J")
    } else if vetor[i] == 12 {
        fmt.Print("Q")
    } else if vetor[i] == 13 {
        fmt.Print("K")
    } else {
        fmt.Print(vetor[i])
    }

    if i < N-1 {
        fmt.Print(", ")
    }
}
fmt.Println("]")
}