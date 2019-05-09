package main

import (
    "fmt"
    "math/rand"
    "time"
)

 func  matrizCuadradaRand()  {
    rand.Seed(time.Now().UnixNano())
    var a [5][5]int
    var i int
    var j int

    for  i = 0; i < 5; i++ {
        for j = 0; j < 5; j++{
            a[i][j] = rand.Intn(5)
            fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
    }

}
    return
}


func main() {
   matrizCuadradaRand()
}