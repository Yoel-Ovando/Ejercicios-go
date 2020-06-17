package main

import "fmt"

//var tabla [10]int

func main() {
	/* Matrices
	var matriz [3][3]int
	matriz[1][1] = 8
	fmt.Println(matriz)
	tabla := [10]int{1, 15, 8, 36, 23, 13, 15, 3, 25, 96}
	//tabla[0] = 1
	//tabla[5] = 15

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}*/

	//Slides

	//matriz := []int{1, 2, 3, 4, 5}
	variante2()
	variante3()

}

func variante2() {
	elemento := [5]int{1, 2, 3, 4, 5}
	porcion := elemento[3:]
	fmt.Println(porcion)
}

func variante3() {

	elementos := make([]int, 5, 20) //crea el slide
	fmt.Println("largo %d , Capacidad %d", len(elementos), cap(elementos))
}
