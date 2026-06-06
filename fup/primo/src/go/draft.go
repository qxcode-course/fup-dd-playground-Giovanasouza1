package main
import "fmt"
import "math"
func main() {
    var n int
    fmt.Scan(&n)
    primo := true
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            primo = false
            break
        }
}
if primo {
    fmt.Println(1)
} else {
    fmt.Println(0)
}
}