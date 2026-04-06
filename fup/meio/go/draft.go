package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    if a > b && a > c && c < b {
        fmt.Println(b)
    } else if a > b && a > c && c > b {
        fmt.Println(c)
    } else if b > a && b > c && c < a {
        fmt.Println(a)
    } else if b > a && b > c && a < c {
        fmt.Println(c)
    } else if c > a && c > b && b < a {
        fmt.Println(a)
    } else if c > a && c > b && b > a {
        fmt.Println(b)
    }

}
