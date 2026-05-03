package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    v := a * b
    for i := a; i <= v; i += 2 {
        fmt.Println(i)
    }
}
