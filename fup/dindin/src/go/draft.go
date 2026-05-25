package main

import (
	"fmt"
)
func main() {
    var D int
    fmt.Scan(&D)
    Sabor := make([]string, D)
    Turno := make([]string, D)

    for i := 0; i < D; i++{
    fmt.Scan(&Sabor[i], &Turno[i])
    }

    l := 0
    c := 0 


    for i := 0; i < D; i++{
        if Sabor[i] == "c" {
            c ++
        } else if Sabor[i] == "l" { 
            l ++
        } }
    if l == c {
        fmt.Println("empate")
    } else if l > c {
        fmt.Println("l") 
    } else if c > l {
        fmt.Println("c") 
    }

    m := 0
    t := 0 

    for i := 0; i < D; i ++ {
        if Turno[i] == "m" {
            m ++
        } else if Turno[i] == "t"{
            t ++
        }
    }

    if m == t {
        fmt.Println("empate")
    } else if m < t {
        fmt.Println("m")
    } else if t < m {
        fmt.Println("t")
    }
}