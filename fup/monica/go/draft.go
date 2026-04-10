package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)
    c := m - (a + b)
    if c > a && c > b {
        fmt.Println(c)
    } else if a > b && a > m {
        fmt.Println(a)
    } else if b > a && b > m {
        fmt.Println(b)
    
    }
    
}
