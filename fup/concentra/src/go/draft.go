package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i := a; i <= 10; i++{
        fmt.Print(i, " ")
    }
}
