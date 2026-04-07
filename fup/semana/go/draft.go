package main
import "fmt"
func main() {
    var a string
    var b int
    fmt.Scan(&a, &b)
    if a == "1" {
        fmt.Println("NAO")
    } else {
            fmt.Println("SIM")
    }; if b > 7 && b < 11 {
        fmt.Println("SIM")
    } else if b > 14 && b < 17 {
        fmt.Println("SIM")
    } else if a == 7 && b > 11 {
        fmt.Println("NAO")
    }

}
