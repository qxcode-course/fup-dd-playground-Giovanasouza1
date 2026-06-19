package main
import "fmt"

    func media(vetor []int) float64 {
        soma := 0.0
        for i := 0; i < len(vetor); i++ {
            soma += float64(vetor[i])
        }
     media := soma / float64(len(vetor))
     return media
    } 
    

    func soldado(vetor []int) {
    m := media(vetor)
    for i := 0; i < len(vetor); i++ {
        if float64(vetor[i]) == m {
            fmt.Println("M")
        } else if float64(vetor[i]) < m {
            fmt.Println("P")
        } else if float64(vetor[i]) > m {
            fmt.Println("G")
        }
    }
    }
    func main (){
        var qtd int
    
        
    }