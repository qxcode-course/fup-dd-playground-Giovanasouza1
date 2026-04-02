package main
import "fmt"
func main() {
    var v float64
    var np int
    fmt.Scan(&v, &np)
    np = 1
    for {
        np += 1
        fmt.Println(np)
        if np <= 10 {
            break
        }
    } 
    valor := v *np





}