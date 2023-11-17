//Determinar si un número es par:

package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("El número es par.")
	} else {
		fmt.Println("El número es impar.")
	}
}
