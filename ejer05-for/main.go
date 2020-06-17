package main

import "fmt"

func main() {

	/*	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}*/
	/*numero := 0
	for {
		fmt.Println("continuo")
		fmt.Println("Ingresa un número")
		fmt.Scanln(&numero)
		if numero ==100{
			break
		}*/
	/*	i := 0
		for i < 10 {

			fmt.Println("El valor de i %d", i)
			if i == 5 {
				fmt.Println("Se multiplica por 2")
				i = i * 2
				fmt.Println("Resultado: %d", i)
				continue
			}
			fmt.Println("paso por aquí")
			i++

		}*/

	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("voy a RUTINA sumando 2 ")
			goto RUTINA
		}
		fmt.Println("valor de rutina i: %d \n", i)
		i++
	}
}
