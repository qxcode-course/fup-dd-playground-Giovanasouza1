package main
import "fmt"
func main() {
    var n1, n2 int
    var simb string
    fmt.Scan(&n1, &n2, &simb)
    if simb == "+" {
        fmt.Println(n1 + n2)
    } else if simb == "-" {
        fmt.Println(n1 - n2)
    } else if simb == "*" {
        fmt.Println(n1 * n2)
    } else if simb == "/" {
        fmt.Println(n1 / n2)
    }
   
}


