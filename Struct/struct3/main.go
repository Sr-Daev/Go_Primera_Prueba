package main

import "fmt"

////Investigar GoRutinas

type Rectangulo struct {
	ancho, alto float64
}

func (r Rectangulo) area() float64 {
	return r.alto * r.ancho
}

func (r *Rectangulo) inc(i float64) {
	r.ancho *= i
	r.alto *= i
}

func main() {
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{alto: 11, ancho: 35}
	fmt.Println("Area", r1.area())
	println("Area", r2.area())
	fmt.Println(r1)
	r1.inc(10)
	fmt.Println(r1)
}
