package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Print("[ ")
    for i := a; i <= b; i++ {
        if i % 2 == 0 {
            continue
        } else if i % 2 != 0 {
            fmt.Printf("%d ", i)
        } else if i == b { 
            break
        }
    }
    fmt.Println("]")
}
