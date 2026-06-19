package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
   
    vetornormal := make([]int, N)
    vetorinvertido := make([]int, N)
    
    for i := 0; i < N; i++ {
        fmt.Scan(&vetornormal[i])
    }

    for i := 0; i < N; i++ {
        vetorinvertido[i] = vetornormal[N - 1 - i]
}
fmt.Print("[ ")
for i := 0; i < N; i++ {
fmt.Print(vetorinvertido[i], " ")
}
fmt.Println("]")
}