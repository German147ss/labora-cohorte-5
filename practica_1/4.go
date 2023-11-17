//Convertir días, horas, minutos y segundos a segundos:

package main

import "fmt"

func main() {
	var dias, horas, minutos, segundos int
	fmt.Print("Ingrese días, horas, minutos y segundos: ")
	fmt.Scan(&dias, &horas, &minutos, &segundos)

	totalSegundos := dias*86400 + horas*3600 + minutos*60 + segundos
	fmt.Println("Total en segundos:", totalSegundos)
}
