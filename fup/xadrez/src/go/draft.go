package main
import "fmt"
func main() {
    var L, C int
    fmt.Scan(&L, &C)

    if L % 2 != 0 && C % 2 != 0 {
        fmt.Println("1")
    } else if L % 2 == 0 && C % 2 == 0 {
        fmt.Println("1")
    } else if L % 2 == 0 && C % 2 != 0 {
        fmt.Println("0")
    } else if L % 2 != 0 && C % 2 == 0 {
        fmt.Println("0")
    }
}