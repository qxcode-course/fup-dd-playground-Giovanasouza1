package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)
    if a % 7 == 0 {
    fmt.Printf("SIM\n")
} else {
        fmt.Println("NAO") 
    } 
}