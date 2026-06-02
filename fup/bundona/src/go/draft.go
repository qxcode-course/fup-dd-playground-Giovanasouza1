package main
import "fmt"
func main() {
    var x, M, D int
    var S string
    fmt.Scan(&x)
    fmt.Scan(&M)
    fmt.Scan(&S)
    fmt.Scan(&D)

    comeco := x * 60 + M
    distancia := D * 10

    var total int
    if S == "H" {
      total =  comeco + distancia
    } else {
       total = comeco - distancia
    }

   posicao := total % 720 
   if posicao < 0 {
    posicao += 720
   } 

   hora := posicao / 60
   min := posicao % 60
   fmt.Printf("%02d %02d\n", hora, min)

   

}