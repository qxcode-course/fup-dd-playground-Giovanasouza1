package main
import "fmt"
func main() {
    const PEDRA int = 0
    const PAPEL int = 1
    const TESOURA int = 2
    const LAGARTO int = 3
    const SPOCK int = 4


    var jog1, jog2 int
    var vit1, vit2 int
    fmt.Scan(&vit1, &vit2)
    for {
        if vit1 > 2 && vit2 > 2 {
             break }

    fmt.Print("Jog1: digite 0(pedra), 1(papel), 2(tesoura), 3(lagarto), 4(spock):")
    fmt.Scan(&jog1)
    fmt.Print("Jog2: digite 0(pedra), 1(papel), 2(tesoura), 3(lagarto), 4(spock):")
    fmt.Scan(&jog2)

    if jog1 == jog2 {
        fmt.Println("empate")
    } else if (jog1 == 0 && jog2 == 3) || (jog1 == 0 && jog2 == 2) || (jog1 == 1 && jog2 == 4) || (jog1 == 1 && jog2 == 0) || (jog1 == 2 && jog2 == 3) || (jog1 == 2 && jog2 == 1) || (jog1 == 3 && jog2 == 1) || (jog1 == 3 && jog2 == 4) || (jog1 == 4 && jog2 == 0) || (jog1 == 4 && jog2 == 2){
        vit1 += 1
        fmt.Println("Placar vit1:", vit1, " Placar vit2:", vit2 )
        fmt.Println("jog1 ganhou!")
    } else {
        vit2 += 1
        fmt.Println("Placar vit1: ", vit1, "  Placar vit2: ", vit2)
        fmt.Println(" jog2 ganhou!")
        
    }
} ; 
if vit1 > vit2 {
    fmt.Println("Jog1 ganhou o jogo!!")
} else {
    fmt.Println("jog2 ganhou o jogo!!")
}
}
