package main
import "fmt"
func main() {
    var a, b ,c int
    var h, l int
    fmt.Scan(&a, &b, &c, &h, &l)
    if a <= l && b <= h || a <= h && b <= l {
        fmt.Println("S")
    } else if c <= l && b <= h || c <= h && b <= l {
        fmt.Println("S")
    } else if a <= l && c <= h || a <= h && c <= l {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}