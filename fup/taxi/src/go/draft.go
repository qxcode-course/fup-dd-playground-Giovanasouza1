package main
import "fmt"
func main() {
    var A, G, Ra, Rg float32
    fmt.Scan(&A, &G, &Ra, &Rg)

    custoA := A / Ra
    custoG := G / Rg

    if custoA > custoG {
        fmt.Println("G")
    } else if custoA < custoG {
        fmt.Println("A") 
    } else if custoA == custoG {
         fmt.Println("G")
    }
}