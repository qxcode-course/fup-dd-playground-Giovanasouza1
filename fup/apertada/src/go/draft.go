package main
import 
"fmt"
func main() {
    numeros := make([]int, 5)

    for i := 0; i < 5; i++ {
        fmt.Scan(&numeros[i])
        }

        menor := numeros[0]

    for i := 0; i < 5; i++ {
        if numeros[i] < menor {
            menor = numeros[i]
        }
        }
        fmt.Println(menor)
    }