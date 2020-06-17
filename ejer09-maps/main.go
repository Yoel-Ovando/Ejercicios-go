package main

import (
	"fmt"
)

func main() {
	miMapa := make(map[string]string)
	fmt.Println(miMapa)
	miMapa["México"] = "D.F"
	miMapa["Argentina"] = "B.A"
	campeonato := map[string]int{
		"Equipo 1": 1,
		"Equipo 2": 2,
		"Equipo 3": 3,
		"Equipo 4": 4}
	/*
	   <<< MOSTRAR PARAMETROS EN FMT>>>
	   %s = string
	   %d =  enteros
	   %t = booleanos

	*/
	campeonato["Equipo 5"] = 5 //Agrega
	campeonato["Equipo 1"] = 9 //Actualiza
	delete(campeonato, "Equipo 3")
	fmt.Println(campeonato)
	for equipo, puntaje := range campeonato { //Muestra los elementos del MAP
		fmt.Printf("El equipo %s, tiene puntaje de: %d\n", equipo, puntaje)

	}

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe %t", puntaje, existe)
}
