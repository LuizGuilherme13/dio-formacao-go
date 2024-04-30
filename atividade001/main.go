package main

import "fmt"

// 1. Conversão de Escalas Termométricas
// Converter o valor  do ponto de ebulição da água de Kelvin para Celsius.

const ebulicaoKelvin = 373

func main() {
	ebulicaoCelsius := ebulicaoKelvin - 273

	fmt.Println("O ponto de ebulição da água em °C é:", ebulicaoCelsius)
}
