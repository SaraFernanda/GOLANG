package main

import ("fmt"
         "math/rand"
         "time")


func hola(){
   fmt.Println("==============================")
   fmt.Println("Hola, Mundo!")
   fmt.Println("==============================")
}

func vector(len int) float64{
   a := make([]float64, len)
   var i int
for  i = 0; i < len-0; i++ {

   fmt.Printf("%d = %f\n",i ,a[i])

   }
   fmt.Println("==============================")
   return 0
}

func vectorRandom(len int) []float64{
    a := make([]float64, len)
    rand.Seed(time.Now().UnixNano()) //generar nuevos numeros aleatorios
    for i := 0; i <= len-1; i++ {
        a[i] = rand.Float64()
        
        fmt.Println(i,"=", a[i])
    }
    fmt.Println("==============================")
    return a
}

func matrizCuadrada(){
   var grid [3][3]int
    fmt.Println(grid)
    fmt.Println("==============================")
    return 
}

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
   fmt.Println("==============================")
    return
}


func main() {
   hola()
   vector(5)
   vectorRandom(10)
   matrizCuadrada()
   matrizCuadradaRand()

   
}

	

