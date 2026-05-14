package main
import "fmt"
func main() {
    var N int
    var lado string
    fmt.Scan(&N, &lado)
    fmt.Print("[ ")
    for i := 0; i <= 9; i++{ 
        if i == N {
        continue
    }
    if lado == "d" {
       if i % 2 == 0 {
        fmt.Print(i, "d", " ")
       } else {
        fmt.Print(i, "e", " ")
       }}
    if lado == "e" {
        if i % 2 == 0 {
         fmt.Print(i, "e", " ")
        } else {
         fmt.Print(i, "d", " ")
        }
    }
for i := N; i <= 9; i++{ 
    if N % 2 == 0 {
        fmt.Print("e")
    } else {
        fmt.Print("d")
    }
}
    }
fmt.Print("ceu ]")
    }


