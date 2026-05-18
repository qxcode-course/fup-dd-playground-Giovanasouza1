package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)
    for {
        if F == 16{
            F = 0
        } else if F < 0 {
            F = 15
        }
        if F == H {
            fmt.Println("S")
            break
        } else if F == P {
            fmt.Println("N")
            break
        } 
        if D == -1 {
            F--
        } else {
            F++
        }

        }
    }
    