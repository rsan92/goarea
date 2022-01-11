package goarea

import "math"

// Pi e uma proporcao numerica
const Pi = 3.1416

// Circ calcula a area de um circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}
