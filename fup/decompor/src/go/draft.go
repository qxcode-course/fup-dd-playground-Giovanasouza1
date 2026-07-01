package main
import "fmt"
func main() {
    var numero int
    fmt.Scan(&numero)

    vetorvazio := make([]int, 8) 

    qtd := 0
    for numero > 0 {
        vetorvazio[qtd] = numero % 10
        qtd++
        numero = numero / 10
    }
     for i := qtd - 1; i >= 0; i-- {
        fmt.Print(vetorvazio[i])

        if i != 0 {
            fmt.Print(" ")
        }
    }
      fmt.Println()
}