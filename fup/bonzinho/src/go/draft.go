package main
import "fmt"
func main() {
    var a, b, c, x float64
    fmt.Scan(&a, &b, &c, &x)
    menor := a

    if b < a && b < c {
        menor = b
    } else if c < a && c < b {
        menor = c
    } else if b == c && b < a{
        menor = b
    }
    MEDIA := ((a + b + c) - menor + x) / 3

    if MEDIA >= 7 {
        fmt.Printf("Aprovado com %.1f\n", MEDIA)
    } else{
        fmt.Printf("Final com %.1f\n", MEDIA)
    }
}