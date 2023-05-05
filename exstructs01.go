package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	circulo := Circulo{raio: 2}
	area := calcularArea(circulo)
	fmt.Printf("A área do círculo é %.2f ", area)
}
