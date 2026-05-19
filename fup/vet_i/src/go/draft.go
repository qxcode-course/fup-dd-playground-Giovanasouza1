package main
import "fmt"
    func main() {
        var n int
        fmt.Scan(&n)
    numeros := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }
    for i := 0; i < n; i++ {
        fmt.Println(numeros[i])
    }
    if n == 0 {
        fmt.Println("")
    }
}