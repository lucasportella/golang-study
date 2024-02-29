package main

import ("fmt"
	"runtime"
)

func main() {
	triggered := false
	for {
		if (!triggered) {
			fmt.Println("programa rodando no sistema: ",runtime.GOOS, runtime.GOARCH )
			fmt.Println("comando para windows: GOOS=windows GOARCH=amd64 go build arquivo.go")
			fmt.Println("comando para mac: GOOS=darwin GOARCH=amd64 go build arquivo.go")
			fmt.Println("comando para linux: GOOS=linux GOARCH=amd64 go build arquivo.go")
			triggered = true
		}
	}

}

