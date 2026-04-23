package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)
    c := m - (a + b)
    if c > a && c > b && c < m {
        fmt.Println(c)
    } else if a > b && a > c && a < m {
        fmt.Println(a)
    } else if b > a && b > c && b < m {
        fmt.Println(b)
    }
    
}
