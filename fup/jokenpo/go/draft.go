package main
import "fmt"
func main() {
    var a, b string   
    var jogador1, jogador2 int64
    fmt.Scan(&a, &b, &jog1, &jog2)
    if a == tesoura && b == papel {
        fmt.Println(jog1)
    } else if a == tesoura && b == pedra {
        fmt.Println(jog2)
    } else if (a == tesoura && b = tesoura) {
        fmt.Println(empate)
    } else if a == pedra && b = papel {
        fmt.Println(jog2)
    } else if a == pedra && b = pedra {
        fmt.Println(empate)
    } else if a == pedra && b = tesoura {
        fmt.Println(jog1)
    } else if a == papel && b = papel {
        fmt.Println(empate)
    } else if a == papel && b = pedra {
        fmt.Println(jogador2)
    } else if a == papel && b = tesoura {
        fmt.Println(jog2)
    }
}
