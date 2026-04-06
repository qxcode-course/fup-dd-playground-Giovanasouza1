package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)
    if a == 0 {
        fmt.Println("you must connect to wifi")
    } else if b == 0 {
        fmt.Println("you need to login first") 
    } else if c == 0 {
        fmt.Println("you must to login as admin")
    } else if a == 1 && b == 1 && c == 1 {
        fmt.Println("done")
    }
}