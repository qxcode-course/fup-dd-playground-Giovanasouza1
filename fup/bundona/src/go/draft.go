package main
import "fmt"
func main() {
    var x, M, D int
    var S rune
    fmt.Scanf("%d%d%c%d",&x,&M,&S,&D)
   fmt.Scanf("%d",&M)
    fmt.Scanf("%c\n",&S)
    fmt.Scanf("%d",&D)

    comeco := (x * 60) + M
    //fmt.Println("comeco: ", comeco)
    distancia := D * 10
    //fmt.Println("distancia: ",distancia)

    var total int
    if S == 'H' {
      total =  comeco + distancia
    } else {
       total = comeco - distancia
    }
    //fmt.Println("total: ",  total)

   posicao := total % 720
   //mt.Println("posicao: ",posicao) 
   if posicao < 0 {
      //fmt.Println("posicao neg: ",posicao)
    posicao += 720
   } 

   hora := posicao / 60
   min := posicao % 60
  fmt.Printf("%02d %02d\n", hora, min)

   

}