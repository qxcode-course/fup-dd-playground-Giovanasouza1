package main
import "fmt"
func main() {
    var menor, maior int
    fmt.Scan(&menor, &maior)
    fmt.Print("[ ")
    a := menor
    b := maior
    for a <= b {
        fmt.Print(a, " ")
        fmt.Print(b, " ")
        a++
        b--
    }
    for a

    }