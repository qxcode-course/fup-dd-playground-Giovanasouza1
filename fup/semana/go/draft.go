package main
import "fmt"
func main() {
    var a int
    var b int
    fmt.Scan(&a, &b)
    if a == 1 || a == 7 && b > 11 || b < 14 && b > 11 || b > 17 || b < 8 { 
        fmt.Println("NAO")
    } else {
        fmt.Println("SIM") }
}
    


