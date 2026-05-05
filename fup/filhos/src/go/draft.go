package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    idade := (b * 2) + a
    for i := a; i < idade; i += 2 {
        fmt.Println(i)
    }
}
