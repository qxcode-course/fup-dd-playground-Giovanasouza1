package main
import "fmt"
func main() {
    var C, M, x int
    fmt.Scan(&C)
    for {
        fmt.Scan(&M)
        x = M + x
        if x == 0 {
            fmt.Println("vazio")
            
        } else if x < C {
            fmt.Println("ainda cabe")
        
        } else if x >= C && x < 2 * C {
            fmt.Println("lotado")
            
        } else if x >= 2 * C {
            fmt.Println("hora de partir")
            break
        }
    }
}
