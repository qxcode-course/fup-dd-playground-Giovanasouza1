package main
import "fmt"
func main() {
    var p, s, e int 
    fmt.Scan(&p, &s, &e)

    altura := 0

   for  {
        if altura + s >= p {
            fmt.Println(altura, "saiu")
            break
        } else {
            fmt.Println(altura, altura + s)
            altura = altura + s - e
        }
   }
}