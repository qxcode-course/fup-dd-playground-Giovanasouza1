package main
import "fmt"
func main() {
    var soma int
    fmt.Scan(&soma)
    
    calculo := (soma - 1) % 26 + 'a'

    if soma == 0 {
        fmt.Println("joguem de novo")
    } else {
        fmt.Printf("%c\n", calculo)
    }

}