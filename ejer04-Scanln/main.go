package main

import (
	"bufio"
	"fmt"
	"os"
)

var result, numero1, numero2 int
var leyenda string

func main() {

	fmt.Println("Escribe 1")
	fmt.Scanln(&numero1)

	fmt.Println("Escribe 2")
	fmt.Scanln(&numero2)

	fmt.Println("Por favor escribe una breve descripción")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	result = numero1 + numero2
	fmt.Println(leyenda, result)

}
