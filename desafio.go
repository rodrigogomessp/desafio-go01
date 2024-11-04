package main

import "fmt"

const PONTO_EBULICAO_KELVIN = 373

func main() {
	tempK := 273
	tempC := PONTO_EBULICAO_KELVIN - tempK
	fmt.Println("Kelvin para Celsius = ", tempC)
}
