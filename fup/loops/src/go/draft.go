package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)

    if a < 0 && a != -360 {
        fmt.Println((a % 360) + 360 )
    } else if a > 0 {
        fmt.Println(a % 360)
    } else if a == 0 {
        fmt.Println(0)
    } else if a == - 360 {
        fmt.Println(0)
    }
    
}