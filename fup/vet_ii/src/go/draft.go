package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    numeros := make([]int, n)

    if n == 0 {
        fmt.Println("[ ]")
        return
    }
    
    fmt.Print("[ ")
    
    
    for i := 0; i < n; i++ {
        fmt.Scan(&numeros[i])
    }
    for i := 0; i < n; i++ {
        fmt.Print(numeros[i], " ")
    }
    fmt.Println("]")
}
