package main

import "fmt"

//ejemplo de funciones anónimas
var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	//Funciones Anónimas
	fmt.Printf("La suma de 5 + 5 = %d", calculo(5, 5))

	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("\n El resultado de 9 - 2 es %d", calculo(9, 2))

	operaciones()
	//CLOUSER
	tablaDel := 3
	MiTabla := tabla(tablaDel)

	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func operaciones() {
	resultado := func() int {
		var a int = 12
		b := 34
		c := 45
		return a + b + c
	}
	fmt.Println("\nAquí", resultado())

}

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
