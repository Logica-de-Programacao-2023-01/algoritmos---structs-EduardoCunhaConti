package main

import "fmt"

type Triangulo struct {
	altura float64
	base   float64
}

func AreaTriangulo(t Triangulo) float64 {
	return (t.base * t.altura) / 2
}

func main() {
	triangulo := Triangulo{base: 2.5, altura: 4.0}
	area := AreaTriangulo(triangulo)
	fmt.Printf("A área do triângulo é %.2f ", area)
}
