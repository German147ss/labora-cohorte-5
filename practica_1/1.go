//Determinar si un número es mayor, igual o menor a cero:

package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Println("El número es mayor a cero.")
	} else if numero < 0 {
		fmt.Println("El número es menor a cero.")
	} else {
		fmt.Println("El número es igual a cero.")
	}
}
