package main
import "fmt"
func main() {
    var h, M, D int
    var S string
    fmt.Scan(&h, &M, &S, &D)

    MIN := D * 10
    MINu := h * 60 + M
    if S == "H" {
        soma := MINu + MIN 
        if soma > 60 {
            if soma % 720
        }
        
    }
    fmt.Println("Hello, World!")
}