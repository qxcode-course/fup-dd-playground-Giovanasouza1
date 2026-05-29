package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    dif := (b - a + 15) % 15
    if dif == 0 {
        fmt.Println("Empate")
    }
    if dif <= 7 && dif > 0 {
        fmt.Println("Jogador 1")
    } else if dif > 7 && dif > 0{
        fmt.Println("Jogador 2")
    }
}
