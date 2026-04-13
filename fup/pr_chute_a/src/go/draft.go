package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    rand.Seed(time.Now().UnixNano())
    numero_secreto := rand.Intn(99) + 1 
    
    for b := 1; b < 100; b++ {
        if b < a
            fmt.Println("Diga um numero entre 0, 100")
        } else if b < a {
            fmt.Println("o numero é maior do que esse")
        } else if b == a {
            fmt.Println("voce ganhou!")
        }
    } 
    fmt.Println(numero_secreto)
}
