package main
import "fmt"
func main() {
    var v, t, c float64
    var horas, distancia float64
    fmt.Scan(&v, &t, &c, &horas, &distancia)
    horas = t / 60
    distancia = v * horas
    desempenho := distancia / c
    fmt.Printf("%.2f\n", desempenho)
    


}
