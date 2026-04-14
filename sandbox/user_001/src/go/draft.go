package main
import "fmt"
func main() {
    // var qtd int
    // fmt.Scan(&qtd)
    // var idades []int = make([]int, qtd)
    // fmt.Println(idades)

    // arr := []int{9, 8, 4, 5, 6, 2, 3}
    // fmt.Print("[ ")
    // for i := 0; i < len(arr); i++ {
    // fmt.Printf(" %d ", arr[i])
    // }
    // fmt.Print("]\n")

    //arr := []int{20, 19, 18, 17, 16}
    //fmt.Print("[ ")
    //for _, valor := range arr {
        //fmt.Printf("%d ", valor)
    //} 
    //fmt.Print("]\n")

    //arr := []int{20, 19, 18, 17, 16}  // essa usa para colocar a virgula antes do numero menos no primeiro
    //fmt.Print("[ ")
    //for i, valor := range arr {
        //if i != 0 {
        //fmt.Print(", ")
    //} 
    //fmt.Printf("%d", valor)
    //} 
    //fmt.Print("]\n")

    var qtd int
    fmt.Scan(&qtd)
    var arr []int = []int{1, 2, 3, 4, 5}
    mostrar_vetor(arr, ", ")

}
