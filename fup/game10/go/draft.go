package main
import "fmt"
func main() {
    var n, d, a int
    fmt.Scan(&n, &d, &a)
    if d == a {
        fmt.Println(0)
    } else if d > a {
        fmt.Println(d - a)
    } else if d < a {
        fmt.Println((d - a)+ n )
    }
}
