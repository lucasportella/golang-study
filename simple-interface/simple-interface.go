package main

import (
	"fmt"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (c circulo) calcArea() float64 {
	return 3.14 * (c.raio * c.raio)
}

func (q quadrado) calcArea() float64 {
	return q.lado * q.lado
}

type figura interface {
	calcArea() float64
}

func info(f figura) float64 {
	return f.calcArea()
}

func main() {
	quadrado1 := quadrado{
		lado: 5,
	}

	circulo1 := circulo{
		raio: 3,
	}
	fmt.Println(info(quadrado1))
	fmt.Println(info(circulo1))
}

