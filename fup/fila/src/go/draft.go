package main
import "fmt"
func main() {
    var qtd int
    fmt.Scan(&qtd)

    var profstars[]int
    var alunos[]int
    

    numeros := make([]int, qtd)
    
    for i := 0; i < qtd; i++ {
        fmt.Scan(&numeros[i])
        if numeros[i] % 2 == 0 {
            profstars = append(profstars,  numeros[i])
        } else {
            alunos = append(alunos, numeros[i])
        }
    }

    fmt.Print("[ ")

    for _, num :=  range alunos {
    fmt.Print(num, " ")
    }
    fmt.Println("]")

    fmt.Print("[ ")
    for _, num := range profstars {
        fmt.Print(num, " ")
    }
    fmt.Println("]")
    }
    