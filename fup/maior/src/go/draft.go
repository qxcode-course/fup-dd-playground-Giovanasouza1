package main

import "fmt"

func main() {
    var chute, valor float64
    var escolha string

    fmt.Scan(&chute, &escolha, &valor)

    if chute == valor {
        fmt.Println("primeiro")
    } else if chute > valor && escolha == "m" {
        fmt.Println("segundo")
    } else if chute < valor && escolha == "M" {
        fmt.Println("segundo")
    } else if chute > valor && escolha == "M" {
        fmt.Println("primeiro")
    } else if chute < valor && escolha == "m" {
        fmt.Println("primeiro")
}}