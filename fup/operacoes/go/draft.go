package main
import "fmt"
func main() {
    var a, b float32
    fmt.Scan(&a, &b)
    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Printf("%.2f\n", a / b)
    fmt.Println(int(a) % int(b))
    }
