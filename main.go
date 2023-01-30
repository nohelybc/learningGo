package main

import "fmt"

func main() {
	// Declarando constantes
	const pi float64 = 3.14159265359
	fmt.Println("Our value of pi is:", pi)

	// Declarando variables enteras
	// Go no compila si las variables no son usadas
	base := 12
	var altura int = 14
	var area int // aqui go agrega 0 como valor de default

	fmt.Println("values:", base, altura, area)

	// Zero values (valores por default)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Wow! look at this pretty zeros:", a, b, c, d)
}
