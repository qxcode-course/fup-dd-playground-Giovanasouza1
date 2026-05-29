package main
import "fmt"
func main() {
    var resp1, resp2, resp3, resp4 string
    fmt.Scan(&resp1, &resp2, &resp3, &resp4)

    acertou := 0
    if resp1 == "d" {
        acertou++
    } 

    if resp2 == "a" {
        acertou++
    }

    if resp3 == "c" {
        acertou++
    }

    if resp4 == "d" {
        acertou++
    }
    
    
    if acertou == 0 {
        fmt.Println("Nunca assistiu")
    } else if acertou == 1 {
        fmt.Println("Ja ouviu falar")
    } else if acertou == 2 {
        fmt.Println("Interessado no assunto")
    } else if acertou == 3 {
        fmt.Println("Fa")
    } else if acertou == 4 {
        fmt.Println("Super Fa")
    }
}