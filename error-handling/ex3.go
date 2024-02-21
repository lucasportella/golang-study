package main

import "fmt"

type erroEspecial struct {
	qualquerCoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("deu erro: %v", e.qualquerCoisa)
}

func erro(e error) {
	// "e" is an interface type, in Go, when you have an interface type, you need to use a type assertion to access the underlying concrete type and its fields.
	// although erroEspecial is a struct, since it implements the Error interface, it can be passe as an arg to this fn.
	fmt.Print(e.(erroEspecial).qualquerCoisa)
}

func main() {
	errEsp := erroEspecial{
		qualquerCoisa: "um erro de qualquer coisa",
	}
	erro(errEsp)
}
