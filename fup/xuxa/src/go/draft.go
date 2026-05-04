package main
import "fmt"
func main() {
    var nome string = "eu gosto de batata frita 5.6 frita/banana"
    palavras := strings.fields(frase)
    for_, elem := range palavras {
        fmt.Println(elem)
    }
    unidos := strings.Join(palavras, "-")
    fmt.Println(unidos)
}
