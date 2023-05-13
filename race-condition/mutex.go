package main


import (
 "fmt"
 "runtime"
 "sync"
)


var contador int
var wg sync.WaitGroup
var mu sync.Mutex


func incrementa() {
	mu.Lock()
 v := contador
 v++
 runtime.Gosched()
 contador = v
 mu.Unlock()
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
