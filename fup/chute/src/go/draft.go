package main
import "fmt"
func main() {
    var v, a, b int
    fmt.Scan(&v, &a, &b)
    primeiro := v - a
    segundo := v - b

if primeiro < 0 {
    primeiro = primeiro * -1
}

if segundo < 0 {
    segundo = segundo * -1
}

if segundo > primeiro {
    fmt.Println("primeiro")
} else if segundo < primeiro {
    fmt.Println("segundo")
} else if segundo == primeiro {
    fmt.Println("empate")
}
}