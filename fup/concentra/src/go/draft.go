package main
import "fmt"
func main() {
    var menor, maior, k  int
    fmt.Scan(&menor, &maior)
    fmt.Print("[ ")
    k = maior
    for i := menor; i <= maior; i++ {
        fmt.Print(i, " ", k, " ")
        k--
    }
    fmt.Println("]")
    }