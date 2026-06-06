package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    contador := 0
    if b == 0 && a == 0 {
        fmt.Println(1)
    } else {
    for i := b; i < 0; i /= 10 {
         aux := i % 10
         if a == aux {
            contador++
         }
        }
    fmt.Println(contador)
}
}