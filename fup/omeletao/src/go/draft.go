package main
import "fmt"
func main() {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)


    if a > b && a > c && a > d {
    fmt.Println(a)
    } else if b >  a && b > c && b > d {
    fmt.Println(b)
    } else if c >  a && c > b && c > d {
    fmt.Println(c)
    } else if d > a  && d > b && d > c {
    fmt.Println(d)
    } else if a == b && b == c && c== d {
        fmt.Println(a)
    } else if a == b && c == d && a > c{
        fmt.Println(a)
    }
}