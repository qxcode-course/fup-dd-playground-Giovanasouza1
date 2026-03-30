package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2) 

    fmt.Printf("%v\n", n1/n2)
    fmt.Printf("%v\n", n1%n2)
    fmt.Printf("%.2f\n", float32(n1)/float32(n2))
    
}
