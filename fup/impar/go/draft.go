package main

import (
	"fmt"
)
func main() {
    var n, da, db int
    fmt.Scan(&n, &da, &db)
    if (da + db) % 2 == 0 && n == 0 {
         fmt.Println(0)
    } else if (da + db) % 2 == 1 && n == 1 {
         fmt.Println(0)
    } else {
        fmt.Println(1)
    }
    
}
