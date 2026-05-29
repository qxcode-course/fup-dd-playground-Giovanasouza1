package main

import "fmt"

func main() {
    var total int
    fmt.Scan(&total)
    var tam int
    fmt.Scan(&tam)

    vetor_normal := make([]int, tam)

    for i := 0; i < tam; i++ {
        fmt.Scan(&vetor_normal[i])
    }

    anterior := vetor_normal[0]
    count := 0
    vetor_repetidos := make([]int, tam)

    for i := 1; i < tam; i++ {
        if vetor_normal[i] == anterior {
            vetor_repetidos[count] = vetor_normal[i]
            count++
        } else {
            anterior = vetor_normal[i]
        }
    }

    vetor_faltantes := make([]int, tam)
    contar_indice := 0

    for i := 1; i <= total; i++ {
        encontrado := 0
        for j := 0; j < tam; j++ {
            if vetor_normal[j] == i {
                encontrado = 1
                break
            }
        }

        if encontrado == 0 {
            vetor_faltantes[contar_indice] = i
            contar_indice++
        }
    }

    for i := 0; i < count; i++ {
        if(i == count-1){
            fmt.Print(vetor_repetidos[i])
        }else{
            fmt.Print(vetor_repetidos[i], " ")
        }
    }

    if count == 0 {
    fmt.Print("N")
    }

    fmt.Println()

    for i := 0; i < contar_indice; i++ {
        if (i == contar_indice-1){
            fmt.Print(vetor_faltantes[i])
        }else{
            fmt.Print(vetor_faltantes[i], " ")
        }
        
    }
    if contar_indice == 0 {
    fmt.Print("N")
}

    fmt.Println()


}