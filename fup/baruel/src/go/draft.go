package main
import "fmt"
func main() {
    var qtd, possui int
    fmt.Scan(&qtd,&possui)

    // Popula o Slice
    slice := make([]int, possui)
    for i := 0; i < possui; i++{
        fmt.Scan(&slice[i])
    }
    fmt.Print("[ ")

    // Verifica as repetidas
    for i:= 1; i < possui; i++{
        if (slice[i] == slice[i - 1]){
            fmt.Print(slice[i]," ");
        }
    }
    fmt.Print("]\n")

    // Encontra as que faltam
    fmt.Print("[ ")
    for i := 1; i <= qtd; i++{
        encontrado := false
        for j := 0; j < possui; j++{
            if slice[j] == i{
                encontrado = true
                break
            }
        }
        if encontrado == false{
            fmt.Print(i," ")
        }
    }
    fmt.Print("]\n")
}