package main
import "fmt"
func main() {
    var L, N int
    fmt.Scan(&L, &N)

    jogadas := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scan(&jogadas[i])
    }   


    jog0 := 0
     for i := 0; i < N; i++ {
        if jogadas[i] < jogadas[i + 1]{
           fmt.Println(jogador[i + 1])
        } else if jogadas[i] == jogadas[i + 1] {
            fmt.
        } else {
            fmt.Print("nenhum")
        }
