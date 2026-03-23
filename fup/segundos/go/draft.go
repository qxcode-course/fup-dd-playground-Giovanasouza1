package main
import "fmt"
func main() {
    var tempo int
    fmt.Scan(&tempo)
    hora := tempo/3600
    tempo = tempo % 3600
    min := tempo/60
    seg := (tempo%60)
    fmt.Printf(hora, min, seg)
}