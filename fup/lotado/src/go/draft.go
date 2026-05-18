package main
import "fmt"
func main() {
    var C, M int
    fmt.Scan(&C)
    for {
        fmt.Scan(&M)
        if M == 0 {
            fmt.Print("vazio")
            
        } else if M < C {
            fmt.Print("ainda cabe")
        
        } else if M > C {
            fmt.Print("lotado")
            
        } else if M == 2 * C {
            fmt.Print("hora de partir")
            break
        }
    }
}
