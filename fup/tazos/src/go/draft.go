package main
import "fmt"
func main() {
    var v int 
    fmt.Scan(&v)


    vetores := make([]int, v)
    fmt.Print("[ ")

    for i := 0; i < v; i++ {
        fmt.Scan(&vetores[i])
    }
    
    
    contador := 0
    for i := 0; i < v; i++ {
        if vetores[i] == vetores[i + 1] {
            contador++
        } else {
            fmt.Println(vetores[v)
            contador = 1
    }
}
    fmt.Println(vetores[v + 1], contador, "]")
}
