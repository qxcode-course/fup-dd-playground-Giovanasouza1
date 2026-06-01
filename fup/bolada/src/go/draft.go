package main
import "fmt"
func main() {
    var saque string
    var F int
    fmt.Scan(& saque, &F)

    poder := 0
    if saque == "b" {
        poder = ((F * 20) - 80) / 10
    } else if saque == "c" {
        poder = ((F * 18) - 80) / 10
    }
   
    if poder < 150 {
        fmt.Println("Fraco, nem passou")
    } else if poder >= 150 && poder < 180 {
                fmt.Println("Perfeito")
    } else if poder >= 180 && poder < 210 {
                fmt.Println("Satisfeito")
    } else if poder > 210 {
                fmt.Println("Muito forte, bola fora")
    }
}