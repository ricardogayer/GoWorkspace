package main

import (
	"fmt"
	"modulo"
)

func main() {
   fmt.Println("Somando dois valores...")
   var a int64 = 200
   var b int64 = 150
   c := modulo.Funcao(a, b)
   fmt.Println(c)
}