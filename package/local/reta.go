package main

import "math"

// Iniciando com letra maiúscula significa publico  ex: GetAll
// Iniciando com letra minuscula significa que ele é privado e só pode ser acessado pelo pacote

// Exempo

// Ponto - função publica
// ponto ou _Ponto - função privada

// Ponto repecenta a coordenada no plano catersiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é responsável por calcular a distancia
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
