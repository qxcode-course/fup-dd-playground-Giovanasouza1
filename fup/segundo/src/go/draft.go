package main
import "fmt"
func main() {
    var h, m, s int
    fmt.Scan(&h, &m, &s)

    hora := h * 60 * 60
    min := m * 60
    
    total := hora + min + s
    total += 1
    tempo := total % 86400

   y := tempo / 3600
   z := (tempo % 3600) / 60
   segundo := tempo % 60

    fmt.Printf("%02d %02d %02d\n", y, z, segundo)
}