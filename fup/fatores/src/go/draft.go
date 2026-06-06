package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)

    fator := 2
    contagem := 0

    for a != 1{
        if a % fator == 0 {
            a = a / fator
            contagem ++
        } else {
            if contagem > 0 {
                fmt.Println(fator, contagem)
            }
            fator ++
            contagem = 0
    }
}
if contagem > 0 {
    fmt.Println(fator, contagem)
}
}