package main
import "fmt"
func main() {
    var a float64
    fmt.Scan(&a)
    if a <= 1000 {
        fmt.Printf("%.2f\n", (a * 1.2))
    } else if a <= 1500 && a > 1000 {
        fmt.Printf("%.2f\n", (a * 1.15))
    } else if a <= 2000 && a > 1500 {
        fmt.Printf("%.2f\n", (a * 1.1))
    } else if a > 2000 {
        fmt.Printf("%.2f\n", (a * 1.05))
    }
    
}
