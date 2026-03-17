package main
import "fmt"
func main() {
    var a = 2
    var b = 3
    var c = 4
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    resultado := a + b + c
    fmt.Println(resultado)
}
