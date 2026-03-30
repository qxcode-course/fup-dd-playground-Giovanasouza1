package main
import "fmt"
func main() {
    var a, b, c, d, e int
    fmt.Scan(&a, &b, &c, &d, &e)
    fmt.Printf("%.2d:%.2d\n", a, b)
    fmt.Print(c / d / e)
}
