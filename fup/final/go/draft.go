package main
import "fmt"
func main() {
    var n1, n2, n3 int
    fmt.Scan(&n1, &n2, &n3)
    if ((n1 + n2)/2 >= 7) {
        fmt.Println("aprovado")
    } else if (n1 + n2)/2 < 4 {
        fmt.Println("reprovado")
    } else if (n1 + n2)/2 > 4 || (n1 + n2)/2 < 7 {
        if (n3 >= 5) {
            fmt. Println("aprovado na final")
        } else if (n3 < 5) {
            fmt.Println("reprovado na final")}
    } 
}