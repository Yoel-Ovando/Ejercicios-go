package main

import "fmt"

func main() {

	fmt.Println(Calculo(2, 46))
	fmt.Println(Calculo(5, 23, 6, 48, 16))
	fmt.Println(Calculo(9))
	fmt.Println(Calculo(6, 12, 89, 36, 16, 25, 21))
	/*fmt.Println(uno(5))
	numero, estado := dos(1)
	fmt.Println(numero)
	fmt.Println(estado)*/
}

func uno(numero int) (z int) {
	z = numero * 2
	return numero * 2

}

func dos(numero int) (int, bool) {

	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d", item)
	}
	return total

}
