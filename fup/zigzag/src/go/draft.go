package main
import "fmt"
func main() {
    var a, z int
    fmt.Scan(&a, &z)
    for i := a; i <= z; i ++ {
        if i % 3 == 0|| i % 5 == 0 {
            fmt.Println()
        }
    }
    fmt.Println("Hello, World!")
}
