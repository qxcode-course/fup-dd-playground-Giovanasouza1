package main
import "fmt"
func main() {
  var D int
  fmt.Scan(&D)
  Kcal := make([]int, D)
  
  for i := 0; i < D ; i++ {
    fmt.Scan(&Kcal[i])
  
  }
  soma := 0
  for i := 0; i < D ; i++ {
    soma += Kcal[i]
}
media := float64(soma / D)
fmt.Printf("%.1f\n", media)
}