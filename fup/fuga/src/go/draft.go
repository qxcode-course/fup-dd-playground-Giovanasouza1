package main
import "fmt"
func main() {
    var H, P, F, D int
    fmt.Scan(&H, &P, &F, &D)

    if P > H && P > F && D  {
        fmt.Print("")
    }
}
