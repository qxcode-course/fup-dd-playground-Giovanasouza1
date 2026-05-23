package main
import "fmt"
func main() {
    var B, C, Q int
    fmt.Scan(&B, &C, &Q)
    tipo := make([]string, Q)

    for i := 0; i < Q; i++ {
        fmt.Scan(&tipo[i])
    }

    Patas := 0

    for i := 0; i < Q; i++ {
    if tipo[i] == "v" {
        Patas+= 4
    } else if tipo[i] == "g" {
        Patas+= 2
    } else if tipo[i] == "c"{
        Patas+= 4
    
    }
    
    }
    fmt.Println(Patas)

    ChuteB := Patas - B
    ChuteC := Patas - C

    if ChuteB < 0 && ChuteC > 0{
       ChuteB = ChuteB * -1
    } else if ChuteC < 0 && ChuteB > 0{
        ChuteC = ChuteC * -1
    } else if ChuteC < 0 && ChuteB < 0 {
        if ChuteC < ChuteB {
            fmt.Println("Chico Bento")
        } else {
            fmt.Println("Cebolinha")
        }
    }

    if ChuteB < ChuteC {
        fmt.Println("Chico Bento")
    } else if ChuteB > ChuteC && ChuteB > 0 && ChuteC > 0 {
        fmt.Println("Cebolinha")
    } else if ChuteB == ChuteC  {
        fmt.Println("empate")
    }


}
