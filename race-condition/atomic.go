package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var contador int64
var wg sync.WaitGroup


func incrementa() {
 atomic.AddInt64(&contador, 1)
 runtime.Gosched()
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
