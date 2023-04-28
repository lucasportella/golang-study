package main

import ("fmt"
	"runtime"
)

func main() {
fmt.Println("programa rodando no sistema: ",runtime.GOOS, runtime.GOARCH )
fmt.Println("comando windows: GOOS=windows GOARCH=amd64 go build arquivo.go")
fmt.Println("comando darwin: GOOS=darwin GOARCH=amd64 go build arquivo.go")
}

