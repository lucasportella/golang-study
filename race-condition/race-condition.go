package main


import (
 "fmt"
 "runtime"
 "sync"
)


var contador int
var wg sync.WaitGroup


func incrementa() {
 v := contador
 v++
 runtime.Gosched()
 contador = v
 wg.Done()
}


func main() {
 wg.Add(999)
 for i := 0; i < 999; i++ {
   go incrementa()
 }
 wg.Wait()
 fmt.Println(contador)
}
