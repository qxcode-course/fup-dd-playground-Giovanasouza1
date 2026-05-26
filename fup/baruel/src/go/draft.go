package main
import "fmt"
func main() {
    var talbum, tbaruel, tsequencia int
    fmt.Scan(&talbum, &tbaruel, &tsequencia)

    album := make([]int, talbum)
    baruel:= make([]int, tbaruel)
    sequencia := make([]int, tsequencia)

    for a := 0; a < talbum; a++ {
        fmt.Scan(&album[a])
    }

    for b := 0; b < tbaruel; b++ {
        fmt.Scan(&baruel[b])
    }

    for s := 0; s < b ; s++ {
        fmt.Scan(&sequencia[s])
    }

}
