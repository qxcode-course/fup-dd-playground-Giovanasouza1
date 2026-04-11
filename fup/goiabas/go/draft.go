package main
import "fmt"
func main() {
    var c, b, g, m int
    fmt.Scan(&c, &b, &g, &m)   
    x := b + g + m
    if x > c && (x % c) > 0 {
        fmt.Println((x / c) + 1)
    } else if x == c {
        fmt.Println(1)
    } else if x < c {
        fmt.Println(1)
    } else if x > c && (x % c ) == 0 {
        fmt.Println(x / c)
    }
}
