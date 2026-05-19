package main
import "fmt"
import "math"
func main() {
    var x, a int
    fmt.Scan(&x)
    a = 0
    b := x
    for b > 0 {
        b = b / 10 
        a++
    }
    var y int
    y = int(math.Pow(10, float64(a-1)))

    numeroinvertido := 0
    c := x
    for x > 0 {
        digito := x % 10
        numeroinvertido += digito * y
        x = x / 10
        y = y / 10
    }
    if c == numeroinvertido {
        fmt.Println(1)
    } else {
        fmt.Println(0)
    }
}
