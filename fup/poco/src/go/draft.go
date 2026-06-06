package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)
    altura := 0

   for  {
    topo := altura + s
        if s <= 0 {
            topo = altura
        }

        if topo >= p {
            fmt.Println(altura, "saiu")
            break
        }
   
       fmt.Println(altura, topo)
       altura = topo - e
       s -= 10

        if altura < 0 {
            fmt.Println(altura, "morreu")
            break
       }
    }
}