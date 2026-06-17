package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    inteiros := make([]int, qtd)


    for i := 0; i < qtd; i++ {
        fmt.Scan(&inteiros[i])
    }

    pares := 0
    impares := 0
    for i := 0; i < qtd; i++ {
        if inteiros[i] % 2 == 0 {
            pares+= inteiros[i] 
        } else {
            impares+= inteiros[i]
        }
}

    if pares > impares {
        fmt.Println("rebeldes")
    } else if pares < impares {
        fmt.Println("soldados")
    } else if pares == impares {
        fmt.Println("empate")
    }
}