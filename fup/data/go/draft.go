package main
import "fmt"
func main() {
    var a, b, c, d, e int
    fmt.Scan(&a, &b, &c, &d, &e)
    fmt.Printf("%02d:%02d ", a, b)
    fmt.Printf("%.2d/%.2d/%.2d\n", c, d, e % 100)

}
