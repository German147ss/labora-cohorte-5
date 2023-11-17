//Convertir segundos a días, horas, minutos y segundos:

package main

import "fmt"

func main() {
	var segundos int
	fmt.Print("Ingrese la cantidad de segundos: ")
	fmt.Scan(&segundos)

	dias := segundos / 86400
	segundos %= 86400
	horas := segundos / 3600
	segundos %= 3600
	minutos := segundos / 60
	segundos %= 60

	fmt.Printf("%d días, %d horas, %d minutos y %d segundos\n", dias, horas, minutos, segundos)
}
