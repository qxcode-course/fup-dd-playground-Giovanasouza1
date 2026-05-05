package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Print("[ ")
    for i := 0 ; i < 10; i++ {
        if n != i {
            fmt.Print(i," " )
     }
     /*
      if n != 10 && i == 9{
        fmt.Print("ceu ")
     } 
        */
    }
    
     if n != 10 {
        fmt.Print("ceu ")
     } 
    fmt.Println("]")
}