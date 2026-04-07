package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    delta := math.Pow(b, 2) - 4 * a * c

    if delta > 0 {
        x1 := (- b + math.Sqrt(delta)) / (2 * a) 
        x2 := (- b - math.Sqrt(delta)) / (2 * a) 
        fmt.Println(x1)
        fmt.Println(x2)
    } if delta < 0 {
        fmt.Println("nao ha raiz real")
    } if delta = 0 {
        
    }

    
    
    
    x := (- b + math.Sqrt(delta)) / (2 * a) //Bhaskara
    fmt.Printf(x)
    if 
}
